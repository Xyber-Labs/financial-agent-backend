package tests

import (
	"financial-agent-backend/tests/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethsim "github.com/ethereum/go-ethereum/ethclient/simulated"
)

type DeployedMockContracts struct {
	TeeWallet             ethcommon.Address
	AavePool              ethcommon.Address
	TrustManagementRouter ethcommon.Address
	ERC20                 ethcommon.Address
}

func DeployMockedContracts(client bind.ContractBackend, opts *bind.TransactOpts) DeployedMockContracts {
	teeWalletAddress, _, _, err := contracts.DeployMockTEEWallet(opts, client)
	if err != nil {
		panic(err)
	}

	aavePoolAddress, _, _, err := contracts.DeployMockAavePool(opts, client)
	if err != nil {
		panic(err)
	}

	trustManagementRouterAddress, _, _, err := contracts.DeployMockTrustManagementRouter(opts, client)
	if err != nil {
		panic(err)
	}

	erc20Address, _, _, err := contracts.DeployMockERC20(opts, client)
	if err != nil {
		panic(err)
	}

	return DeployedMockContracts{
		TeeWallet:             teeWalletAddress,
		AavePool:              aavePoolAddress,
		TrustManagementRouter: trustManagementRouterAddress,
		ERC20:                 erc20Address,
	}
}

func CreateSimulatedNode(alloc ethtypes.GenesisAlloc) *ethsim.Backend {
	backend := ethsim.NewBackend(alloc)
	return backend
}
