package onchain

import (
	"context"
	"financial-agent-backend/core/abi/bindings/AaveAToken"
	"financial-agent-backend/core/abi/bindings/AavePool"
	"financial-agent-backend/core/abi/bindings/ERC20"
	"financial-agent-backend/core/abi/bindings/TrustManagementRouter"
	"financial-agent-backend/core/abi/bindings/TrustManagementWallet"
	"financial-agent-backend/core/abi/bindings/WETH"
	"financial-agent-backend/core/transactor"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
)

type TrustManagementProvider struct {
	client                bind.ContractBackend
	TrustManagementRouter *TrustManagementRouter.TrustManagementRouter
	AavePoolAddress       ethcommon.Address
	AavePool              *AavePool.AavePool
	Transactor            *transactor.Transactor
	createTxOpts          *bind.TransactOpts
	callOpts              *bind.CallOpts

	// Optinal wrapped ERC20 native token for the network. If provided, enables DepositNative function
	NativeErc20Address *ethcommon.Address
	NativeErc20        *WETH.WETH
}

func NewTrustManagementProvider(
	client bind.ContractBackend,
	transactor *transactor.Transactor,
	trustManagementRouter *TrustManagementRouter.TrustManagementRouter,
	aavePoolAddress ethcommon.Address,
	aavePool *AavePool.AavePool,
	callOpts *bind.CallOpts,
	nativeErc20Address *ethcommon.Address,
	nativeErc20 *WETH.WETH,
) *TrustManagementProvider {
	// createTxOpts only used to consturct the calldata for the transactions
	// to be later reconstructed in a batch transaction. Therefore we don't actually
	// want to sign, estimate or execute them.
	createTxOpts := bind.TransactOpts{
		From:     ethcommon.Address{},
		NoSend:   true,
		GasLimit: 1000000,
		Signer: func(address ethcommon.Address, tx *ethtypes.Transaction) (*ethtypes.Transaction, error) {
			return tx, nil
		},
		Context: context.Background(),
	}
	return &TrustManagementProvider{
		client:                client,
		Transactor:            transactor,
		createTxOpts:          &createTxOpts,
		TrustManagementRouter: trustManagementRouter,
		AavePoolAddress:       aavePoolAddress,
		AavePool:              aavePool,
		callOpts:              callOpts,
		NativeErc20Address:    nativeErc20Address,
		NativeErc20:           nativeErc20,
	}
}

func (p *TrustManagementProvider) Deposit(
	userAddress ethcommon.Address,
	tokenAddress ethcommon.Address,
	tokenAmount *big.Int,
) (*ethtypes.Transaction, error) {
	log.Info().
		Str("userAddress", userAddress.String()).
		Str("tokenAddress", tokenAddress.String()).
		Int64("tokenAmount", tokenAmount.Int64()).
		Msg("Executing TrustManagementProvider.Deposit")

	// Get user wallet address
	userWalletAddress, err := p.TrustManagementRouter.GetWalletAddress(p.callOpts, userAddress)
	if err != nil {
		return nil, err
	}

	if userWalletAddress.IsDeployed == false {
		return nil, fmt.Errorf("user wallet %s is not deployed", userWalletAddress.WalletAddress.String())
	}

	userWallet, err := TrustManagementWallet.NewTrustManagementWallet(userWalletAddress.WalletAddress, p.client)
	if err != nil {
		return nil, err
	}

	token, err := ERC20.NewERC20(tokenAddress, p.client)
	if err != nil {
		return nil, err
	}

	// Create Token.Approve transaction
	approveTx, err := token.Approve(p.createTxOpts, p.AavePoolAddress, tokenAmount)
	if err != nil {
		return nil, err
	}

	if approveTx.To() == nil {
		return nil, fmt.Errorf("INTERNAL ERROR: approve transaction has no recipient")
	}

	// Since we need to call approve from wallet's context, we need to wrap the
	// approve transaction with Wallet.execute and pass it to Router. In this case
	// Router.execute executes Wallet.execute which calls Token.Approve from wallet's context.
	wrappedApproveTx := TrustManagementWallet.Transaction{
		Target: *approveTx.To(),
		Data:   approveTx.Data(),
		Value:  approveTx.Value(),
	}

	// Create AavePool.supply transaction
	aaveSupplyTx, err := p.AavePool.Supply(
		p.createTxOpts,
		tokenAddress,
		tokenAmount,
		userAddress,
		0,
	)
	if err != nil {
		return nil, err
	}

	if aaveSupplyTx.To() == nil {
		return nil, fmt.Errorf("INTERNAL ERROR: supply transaction has no recipient")
	}

	wrappedSupplyTx := TrustManagementWallet.Transaction{
		Target: *aaveSupplyTx.To(),
		Data:   aaveSupplyTx.Data(),
		Value:  aaveSupplyTx.Value(),
	}

	walletExecuteTx, err := userWallet.Execute(p.createTxOpts, []TrustManagementWallet.Transaction{wrappedApproveTx, wrappedSupplyTx})
	if err != nil {
		return nil, err
	}

	log.Info().Msg("Batching and executing transactions")

	// Batch and execute transactions
	tx, err := p.Transactor.BatchAndExecute([]*ethtypes.Transaction{walletExecuteTx})
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (p *TrustManagementProvider) DepositNative(
	userAddress ethcommon.Address,
	nativeAmount *big.Int,
) (*ethtypes.Transaction, error) {

	if p.NativeErc20Address == nil || p.NativeErc20 == nil {
		return nil, fmt.Errorf("NativeErc20 is not set, DepositNative is not available")
	}

	nativeWrapOpts := bind.TransactOpts{
		From:     userAddress,
		NoSend:   true,
		Value:    nativeAmount,
		GasLimit: 1000000,
		Signer: func(address ethcommon.Address, tx *ethtypes.Transaction) (*ethtypes.Transaction, error) {
			return tx, nil
		},
		Context: context.Background(),
	}
	nativeWrapTx, err := p.NativeErc20.Deposit(
		&nativeWrapOpts,
	)
	if err != nil {
		return nil, err
	}

	// Create AavePool.supply transaction
	aaveSupplyTx, err := p.AavePool.Supply(
		p.createTxOpts,
		*p.NativeErc20Address,
		nativeAmount,
		userAddress,
		0,
	)
	if err != nil {
		return nil, err
	}

	log.Info().Msg("Batching and executing transactions")

	// Batch and execute transactions
	tx, err := p.Transactor.BatchAndExecute([]*ethtypes.Transaction{nativeWrapTx, aaveSupplyTx})
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (p *TrustManagementProvider) Claim(
	tokenAddress ethcommon.Address,
	userAddress ethcommon.Address,
	amount *big.Int,
) (*ethtypes.Transaction, error) {
	// Create aave transaction that withdraws yield to user
	aaveWithdrawTx, err := p.AavePool.Withdraw(
		p.createTxOpts,
		tokenAddress,
		amount,
		userAddress,
	)
	if err != nil {
		return nil, err
	}

	// Create tx that executes TrustManagementRouter.claim
	claimTx, err := p.TrustManagementRouter.Claim(
		p.createTxOpts,
		userAddress,
		userAddress,
		[]ethcommon.Address{tokenAddress},
		[]*big.Int{amount},
	)
	if err != nil {
		return nil, err
	}

	// Batch and execute transactions
	tx, err := p.Transactor.BatchAndExecute([]*ethtypes.Transaction{aaveWithdrawTx, claimTx})
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (p *TrustManagementProvider) Withdraw(
	tokenAddress ethcommon.Address,
	amount *big.Int,
	userAddress ethcommon.Address,
) (*ethtypes.Transaction, error) {

	// Call AavePool.withdraw
	aaveWithdrawTx, err := p.AavePool.Withdraw(
		p.createTxOpts,
		tokenAddress,
		amount,
		userAddress,
	)
	if err != nil {
		return nil, err
	}

	deposits, err := p.TrustManagementRouter.GetDeposits(
		p.callOpts,
		userAddress,
		tokenAddress,
	)
	if err != nil {
		return nil, err
	}

	var depositIds []*big.Int
	unlockedDepositAmount := big.NewInt(0)
	currentTime := big.NewInt(time.Now().Unix())

	// In this loop we need to identify deposits which we will have to liquidate.
	// Deposits are qualified for liquidation when they are unlocked.
	// We also need to verify here that amount of deposits are enough to cover
	// the withdraw amount.
	for i, deposit := range deposits {
		if deposit.LockedUntil.Cmp(currentTime) < 0 {
			depositIds = append(depositIds, big.NewInt(int64(i)))
			unlockedDepositAmount = unlockedDepositAmount.Add(unlockedDepositAmount, deposit.Amount)
		}
	}

	if unlockedDepositAmount.Cmp(amount) < 0 {
		return nil, fmt.Errorf("not enough deposits to cover withdraw amount, total deposit amount: %s, withdraw amount: %s", unlockedDepositAmount, amount)
	}

	// Calculate amount of token with yield considering compounded yield of atoken on Aave
	aTokenAddress, err := p.AavePool.GetReserveAToken(p.callOpts, tokenAddress)
	if err != nil {
		return nil, err
	}

	if (aTokenAddress == ethcommon.Address{}) {
		return nil, fmt.Errorf("aToken is nil or zero address")
	}

	aToken, err := AaveAToken.NewAaveAToken(aTokenAddress, p.client)
	if err != nil {
		return nil, err
	}

	userATokenBalance, err := aToken.BalanceOf(p.callOpts, userAddress)
	if err != nil {
		return nil, err
	}

	amountWithYield, err := calculateAmountWithYield(amount, unlockedDepositAmount, userATokenBalance)
	if err != nil {
		return nil, err
	}

	// Call TrustManagementRouter.withdraw
	routerWithdrawTx, err := p.TrustManagementRouter.Withdraw(
		p.createTxOpts,
		userAddress,
		tokenAddress,
		userAddress,
		depositIds,
		amountWithYield,
	)
	if err != nil {
		return nil, err
	}

	// Batch and execute transactions
	tx, err := p.Transactor.BatchAndExecute([]*ethtypes.Transaction{aaveWithdrawTx, routerWithdrawTx})
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func calculateAmountWithYield(
	amount *big.Int,
	totalDepositAmount *big.Int,
	aTokenBalance *big.Int,
) (*big.Int, error) {

	if aTokenBalance.Cmp(totalDepositAmount) < 0 {
		return nil, fmt.Errorf(
			"INTERNAL ERROR: aToken balance is less than total deposit amount, aToken balance: %s, total deposit amount: %s",
			aTokenBalance, totalDepositAmount,
		)
	}

	yieldAmount := new(big.Int).Sub(aTokenBalance, totalDepositAmount)
	amountWithYield := yieldAmount.Add(yieldAmount, amount)
	return amountWithYield, nil
}
