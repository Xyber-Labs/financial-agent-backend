package tests

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"financial-agent-backend/config"
	"financial-agent-backend/core/abi/bindings/AavePool"
	"financial-agent-backend/core/abi/bindings/TEEWallet"
	"financial-agent-backend/core/abi/bindings/TrustManagementRouter"
	"financial-agent-backend/core/server"
	"financial-agent-backend/core/transactor"
	"financial-agent-backend/core/utils"
	"financial-agent-backend/tests/mocks"
)

func TestDeposit(t *testing.T) {
	r := require.New(t)

	senderPrivKeyStr := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	adminKey, err := utils.ParseKeyFromHex(senderPrivKeyStr)
	r.NoError(err)
	pub, ok := adminKey.Public().(*ecdsa.PublicKey)
	r.True(ok)
	adminPubkey := ethcrypto.PubkeyToAddress(*pub)
	ethBackend := CreateSimulatedNode(ethtypes.GenesisAlloc{
		adminPubkey: {Balance: big.NewInt(10000000000000000)},
	})

	chainId, err := ethBackend.Client().ChainID(context.Background())
	r.NoError(err)
	transactOpts, err := utils.GetTransactOptsFromPrivateKeyString(senderPrivKeyStr, chainId)
	r.NoError(err)

	mockedContracts := DeployMockedContracts(ethBackend.Client(), transactOpts)
	mteeService := mocks.NewMockTeeService(t)

	trustManagementRouter, err := TrustManagementRouter.NewTrustManagementRouter(
		mockedContracts.TrustManagementRouter,
		ethBackend.Client(),
	)
	r.NoError(err)
	teeWallet, err := TEEWallet.NewTEEWallet(
		mockedContracts.TeeWallet,
		ethBackend.Client(),
	)
	r.NoError(err)
	aavePool, err := AavePool.NewAavePool(
		mockedContracts.AavePool,
		ethBackend.Client(),
	)
	r.NoError(err)

	testTransactor, err := transactor.NewTransactor(
		ethBackend.Client(),
		transactOpts,
		trustManagementRouter,
		teeWallet,
		mteeService,
	)
	r.NoError(err)

	// Create HTTP server
	serverConfig := config.HttpServerConfig{
		Port: 8080,
	}
	agentServer := server.NewHttpAgentServer(
		&serverConfig,
		testTransactor,
		trustManagementRouter,
		aavePool,
	)
	_ = agentServer
}
