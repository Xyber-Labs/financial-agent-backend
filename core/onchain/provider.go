package onchain

import (
	"context"
	"financial-agent-backend/core/abi/bindings/AavePool"
	"financial-agent-backend/core/abi/bindings/TrustManagementRouter"
	"financial-agent-backend/core/transactor"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
)

type TrustManagementProvider struct {
	client                bind.ContractBackend
	TrustManagementRouter *TrustManagementRouter.TrustManagementRouter
	AavePool              *AavePool.AavePool
	Transactor            *transactor.Transactor
	createTxOpts          *bind.TransactOpts
}

func NewTrustManagementProvider(
	client bind.ContractBackend,
	transactor *transactor.Transactor,
	trustManagementRouter *TrustManagementRouter.TrustManagementRouter,
	aavePool *AavePool.AavePool,
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
		AavePool:              aavePool,
	}
}

func (p *TrustManagementProvider) Deposit(
	userAddress ethcommon.Address,
	tokenAddress ethcommon.Address,
	tokenAmount *big.Int,
	deadline *big.Int,
	permit TrustManagementRouter.Permit,
) (*ethtypes.Transaction, error) {
	depositWithPermitTx, err := p.TrustManagementRouter.DepositWithPermit(
		p.createTxOpts,
		userAddress,
		tokenAddress,
		deadline,
		permit,
	)
	if err != nil {
		return nil, err
	}

	log.Info().Msg("Creating AavePool.supply transaction")

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

	log.Info().Msg("Batching and executing transactions")

	// Batch and execute transactions
	tx, err := p.Transactor.BatchAndExecute([]*ethtypes.Transaction{depositWithPermitTx, aaveSupplyTx})
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
	withdrawTx, err := p.AavePool.Withdraw(
		p.createTxOpts,
		tokenAddress,
		amount,
		userAddress,
	)
	if err != nil {
		return nil, err
	}

	// Batch and execute transactions
	tx, err := p.Transactor.BatchAndExecute([]*ethtypes.Transaction{withdrawTx})
	if err != nil {
		return nil, err
	}

	return tx, nil
}
