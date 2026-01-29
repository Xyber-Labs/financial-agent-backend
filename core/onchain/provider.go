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
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// In TrustManagement systems operations with native use this address to indicate that.
// We parse things such as events with this address to determine that it's a native token
// operation.
const TrustManagementNativeTokenLabel = "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"

type TrustManagementProvider struct {
	client                bind.ContractBackend
	TrustManagementRouter *TrustManagementRouter.TrustManagementRouter
	AavePoolAddress       ethcommon.Address
	AavePool              *AavePool.AavePool
	Transactor            *transactor.Transactor
	createTxOpts          *bind.TransactOpts
	callOpts              *bind.CallOpts

	logger zerolog.Logger

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
		logger:                log.With().Str("component", "TrustManagementProvider").Logger(),
	}
}

func (p *TrustManagementProvider) Deposit(
	userAddress ethcommon.Address,
	tokenAddress ethcommon.Address,
	tokenAmount *big.Int,
) (*ethtypes.Transaction, error) {
	p.logger.Info().
		Str("userAddress", userAddress.String()).
		Str("tokenAddress", tokenAddress.String()).
		Int64("tokenAmount", tokenAmount.Int64()).
		Msg("Executing TrustManagementProvider.Deposit")

	// Get user wallet address
	userWalletAddress, err := p.TrustManagementRouter.GetWalletAddress(p.callOpts, userAddress)
	if err != nil {
		return nil, err
	}

	if !userWalletAddress.IsDeployed {
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
		userWalletAddress.WalletAddress,
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

	p.logger.Info().
		Str("userAddress", userAddress.String()).
		Str("userWallet", userWalletAddress.WalletAddress.String()).
		Str("tokenAddress", tokenAddress.String()).
		Str("tokenAmount", tokenAmount.String()).
		Str("aavePool", p.AavePoolAddress.String()).
		Msg("Executing batched transaction for deposit")

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

	userWalletData, err := p.TrustManagementRouter.GetWalletAddress(p.callOpts, userAddress)
	if err != nil {
		return nil, err
	}

	if !userWalletData.IsDeployed {
		return nil, fmt.Errorf("user wallet %s is not deployed", userWalletData.WalletAddress.String())
	}

	userWallet, err := TrustManagementWallet.NewTrustManagementWallet(userWalletData.WalletAddress, p.client)
	if err != nil {
		return nil, err
	}

	nativeDepositTx, err := p.NativeErc20.Deposit(p.createTxOpts)
	if err != nil {
		return nil, err
	}

	nativeApproveTx, err := p.NativeErc20.Approve(p.createTxOpts, p.AavePoolAddress, nativeAmount)
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

	walletExecuteTxs := []TrustManagementWallet.Transaction{
		{
			Target: *nativeDepositTx.To(),
			Data:   nativeDepositTx.Data(),
			Value:  nativeDepositTx.Value(),
		},
		{
			Target: *nativeApproveTx.To(),
			Data:   nativeApproveTx.Data(),
			Value:  nativeApproveTx.Value(),
		},
		{
			Target: *aaveSupplyTx.To(),
			Data:   aaveSupplyTx.Data(),
			Value:  aaveSupplyTx.Value(),
		},
	}

	walletExecuteTx, err := userWallet.Execute(
		p.createTxOpts,
		walletExecuteTxs,
	)
	if err != nil {
		return nil, err
	}

	p.logger.Info().
		Str("userAddress", userAddress.String()).
		Str("userWallet", userWalletData.WalletAddress.String()).
		Str("nativeAmount", nativeAmount.String()).
		Str("aavePool", p.AavePoolAddress.String()).
		Msg("Executing batched transaction for native deposit")

	tx, err := p.Transactor.BatchAndExecute([]*ethtypes.Transaction{walletExecuteTx})
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

	userWalletAddress, err := p.TrustManagementRouter.GetWalletAddress(p.callOpts, userAddress)
	if err != nil {
		return nil, err
	}

	if !userWalletAddress.IsDeployed {
		return nil, fmt.Errorf("wallet is not deployed")
	}

	userATokenBalance, err := aToken.BalanceOf(p.callOpts, userWalletAddress.WalletAddress)
	if err != nil {
		return nil, err
	}

	amountWithYield, err := calculateAmountWithYield(amount, unlockedDepositAmount, userATokenBalance)
	if err != nil {
		return nil, err
	}

	userWallet, err := TrustManagementWallet.NewTrustManagementWallet(userWalletAddress.WalletAddress, p.client)
	if err != nil {
		return nil, err
	}

	// Call AavePool.withdraw
	aaveWithdrawTx, err := p.AavePool.Withdraw(
		p.createTxOpts,
		tokenAddress,
		amountWithYield,
		userWalletAddress.WalletAddress,
	)
	if err != nil {
		return nil, err
	}

	if aaveWithdrawTx.To() == nil {
		return nil, fmt.Errorf("aave withdraw tx to is nil")
	}

	wrappedAaveWithdrawTx := TrustManagementWallet.Transaction{
		Target: *aaveWithdrawTx.To(),
		Data:   aaveWithdrawTx.Data(),
		Value:  aaveWithdrawTx.Value(),
	}

	walletExecuteTx, err := userWallet.Execute(p.createTxOpts, []TrustManagementWallet.Transaction{wrappedAaveWithdrawTx})
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

	p.logger.Info().
		Str("userAddress", userAddress.String()).
		Str("tokenAddress", tokenAddress.String()).
		Str("userWallet", userWalletAddress.WalletAddress.String()).
		Str("requestedAmount", amount.String()).
		Str("amountWithYield", amountWithYield.String()).
		Int("depositIdsCount", len(depositIds)).
		Msg("Executing batched transaction for withdraw")

	// Batch and execute transactions
	tx, err := p.Transactor.BatchAndExecute([]*ethtypes.Transaction{walletExecuteTx, routerWithdrawTx})
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
