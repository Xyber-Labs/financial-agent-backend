package tests

import (
	"context"
	"testing"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	"financial-agent-backend/core/transactor"
	"financial-agent-backend/core/utils"
	"financial-agent-backend/tests/mocks"
)

func TestDeposit(t *testing.T) {
	r := require.New(t)

	ethBackend := CreateSimulatedNode(ethtypes.GenesisAlloc{})
	senderPrivKeyStr := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	senderPrivKey, err := utils.ParseKeyFromHex(senderPrivKeyStr)
	r.NoError(err)

	chainId, err := ethBackend.Client().ChainID(context.Background())
	r.NoError(err)
	transactOpts, err := utils.GetTransactOptsFromPrivateKeyString(senderPrivKeyStr, chainId)
	r.NoError(err)

	mockedContracts := DeployMockedContracts(ethBackend.Client(), transactOpts)
	mteeService := mocks.NewMockTeeService(t)

	testTransactor, err := transactor.NewTransactor(
		ethBackend.Client(),
		transactOpts,
		mockedContracts.TrustManagementRouter,
		mockedContracts.TeeWallet,
		mteeService,
	)
	r.NoError(err)
}
