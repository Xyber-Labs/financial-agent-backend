package tests

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"testing"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"financial-agent-backend/config"
	"financial-agent-backend/core/abi/bindings/AavePool"
	"financial-agent-backend/core/abi/bindings/TEEWallet"
	"financial-agent-backend/core/abi/bindings/TrustManagementRouter"
	"financial-agent-backend/core/onchain"
	"financial-agent-backend/core/server"
	"financial-agent-backend/core/transactor"
	"financial-agent-backend/core/utils"
	"financial-agent-backend/tests/mocks"
)

func TestDeposit(t *testing.T) {
	r := require.New(t)
	ctx := context.Background()

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
	ethBackend.Commit()
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

	// Create onchain provider instance
	trustManagementProvider := onchain.NewTrustManagementProvider(
		ethBackend.Client(),
		testTransactor,
		trustManagementRouter,
		aavePool,
	)

	// Create HTTP server
	serverConfig := config.HttpServerConfig{
		Port: 8080,
	}
	agentServer := server.NewHttpAgentServer(
		&serverConfig,
		testTransactor,
		trustManagementProvider,
	)

	ginCtx, cancel := context.WithTimeout(ctx, 3000*time.Second)
	defer cancel()
	go agentServer.Start(ginCtx)

	// Arbitrary wait for server to startup
	time.Sleep(time.Second)

	// Get nonceBefore before sending the request
	nonceBefore, err := ethBackend.Client().NonceAt(ctx, adminPubkey, nil)
	r.NoError(err)
	fmt.Println("nonceBefore", nonceBefore)

	// Make /deposit HTTP request to server
	depositUrl := fmt.Sprintf("http://127.0.0.1:%d/deposit", serverConfig.Port)
	depositReqBody := server.DepositRequest{
		UserAddress:  "0xAc0974bec39a17E36Ba4a6B4d238FF944bAcB478",
		ChainId:      1,
		TokenAddress: "0xAc0974bec39a17E36Ba4a6B4d238FF944bAcB478",
		Amount:       "10000000000000000",
		Deadline:     1700000000,
		SigV:         27,
		SigR:         "0x0fe8e32c2e530da67ead433ce694e845bf72b711aa385add3dc79b423a812f3d",
		SigS:         "0x0fe8e32c2e530da67ead433ce694e845bf72b711aa385add3dc79b423a812f3d",
	}
	reqBody, err := json.Marshal(depositReqBody)
	r.NoError(err)
	resp, err := http.Post(depositUrl, "application/json", bytes.NewBuffer(reqBody))
	r.NoError(err)
	defer resp.Body.Close()
	ethBackend.Commit()

	var depositResponse server.DepositResponse
	err = json.NewDecoder(resp.Body).Decode(&depositResponse)
	r.NoError(err)
	fmt.Println("responseText", depositResponse)
	sentTx, _, err := ethBackend.Client().TransactionByHash(ctx, ethcommon.HexToHash(depositResponse.Tx))
	r.NoError(err)
	fmt.Println("sentTx", sentTx)

	// Get nonceAfter after sending the request
	nonceAfter, err := ethBackend.Client().NonceAt(ctx, adminPubkey, nil)
	r.NoError(err)
	fmt.Println("nonceAfter", nonceAfter)

	r.Equal(nonceBefore+1, nonceAfter)

	_ = agentServer
}
