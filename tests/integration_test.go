package tests

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"financial-agent-backend/config"
	"financial-agent-backend/core/abi/bindings/AavePool"
	"financial-agent-backend/core/abi/bindings/TrustManagementRouter"
	"financial-agent-backend/core/onchain"
	"financial-agent-backend/core/server"
	"financial-agent-backend/core/transactor"
	"financial-agent-backend/core/utils"
	"financial-agent-backend/tests/contracts"
	"financial-agent-backend/tests/mocks"
)

const senderPrivKeyStr = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

func TestDeposit(t *testing.T) {
	r := require.New(t)
	ctx := context.Background()

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

	aavePool, err := AavePool.NewAavePool(
		mockedContracts.AavePool,
		ethBackend.Client(),
	)
	r.NoError(err)

	testTransactor, err := transactor.NewTransactor(
		ethBackend.Client(),
		chainId,
		transactOpts,
		mockedContracts.TrustManagementRouter,
		trustManagementRouter,
		mteeService,
	)
	r.NoError(err)

	// Create onchain provider instance
	trustManagementProvider := onchain.NewTrustManagementProvider(
		ethBackend.Client(),
		testTransactor,
		trustManagementRouter,
		aavePool,
		&bind.CallOpts{},
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
	r.Equal(http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	ethBackend.Commit()

	var depositResponse server.DepositResponse
	err = json.NewDecoder(resp.Body).Decode(&depositResponse)
	r.NoError(err)
	fmt.Println("depositResponse", depositResponse)
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

func TestWithdraw(t *testing.T) {
	r := require.New(t)
	ctx := context.Background()

	adminKey, err := utils.ParseKeyFromHex(senderPrivKeyStr)
	r.NoError(err)
	pub, ok := adminKey.Public().(*ecdsa.PublicKey)
	r.True(ok)
	adminPubkey := ethcrypto.PubkeyToAddress(*pub)
	ethBackend := CreateSimulatedNode(ethtypes.GenesisAlloc{
		adminPubkey: {Balance: big.NewInt(10000000000000000)},
	})

	chainId, err := ethBackend.Client().ChainID(ctx)
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

	testTransactor, err := transactor.NewTransactor(
		ethBackend.Client(),
		chainId,
		transactOpts,
		mockedContracts.TrustManagementRouter,
		trustManagementRouter,
		mteeService,
	)
	r.NoError(err)

	aavePool, err := AavePool.NewAavePool(
		mockedContracts.AavePool,
		ethBackend.Client(),
	)
	r.NoError(err)

	trustManagementProvider := onchain.NewTrustManagementProvider(
		ethBackend.Client(),
		testTransactor,
		trustManagementRouter,
		aavePool,
		&bind.CallOpts{},
	)

	// Create HTTP server
	serverConfig := config.HttpServerConfig{
		Port: 8081,
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

	mockedTrustManagementRouter, err := contracts.NewMockTrustManagementRouter(mockedContracts.TrustManagementRouter, ethBackend.Client())
	r.NoError(err)

	mockedAavePool, err := contracts.NewMockAavePool(mockedContracts.AavePool, ethBackend.Client())
	r.NoError(err)

	_, err = mockedAavePool.MockSetReserveAToken(transactOpts, mockedContracts.ERC20)
	r.NoError(err)
	ethBackend.Commit()

	mockedErc20, err := contracts.NewMockERC20(mockedContracts.ERC20, ethBackend.Client())
	r.NoError(err)

	testCases := []struct {
		desc          string
		deposits      []contracts.MockTrustManagementRouterDeposit
		amount        string
		aTokenBalance *big.Int
		expectedErr   error
	}{
		{
			desc: "sucessfull withdraw of 100%",
			deposits: []contracts.MockTrustManagementRouterDeposit{
				{
					Amount:      big.NewInt(5000),
					LockedUntil: big.NewInt(time.Now().Add(-time.Hour).Unix()),
				},
				{
					Amount:      big.NewInt(5000),
					LockedUntil: big.NewInt((time.Now().Add(-time.Hour)).Unix()),
				},
			},
			amount:        "10000",
			aTokenBalance: big.NewInt(10000),
			expectedErr:   nil,
		},
		{
			desc: "error on withdraw of 110%",
			deposits: []contracts.MockTrustManagementRouterDeposit{
				{
					Amount:      big.NewInt(5000),
					LockedUntil: big.NewInt(time.Now().Add(-time.Hour).Unix()),
				},
				{
					Amount:      big.NewInt(5000),
					LockedUntil: big.NewInt((time.Now().Add(-time.Hour)).Unix()),
				},
			},
			amount:        "11000",
			aTokenBalance: big.NewInt(10000),
			expectedErr:   fmt.Errorf("{\"error\":\"not enough deposits to cover withdraw amount, total deposit amount: 10000, withdraw amount: 11000\"}"),
		},
	}

	// Make /withdraw HTTP request to server
	withdrawUrl := fmt.Sprintf("http://127.0.0.1:%d/withdraw", serverConfig.Port)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {

			// Set mock aToken balance
			_, err = mockedErc20.MockSetBalance(transactOpts, tc.aTokenBalance)
			r.NoError(err)
			ethBackend.Commit()

			// Set mock deposits in the contract
			_, err = mockedTrustManagementRouter.MockSetDeposits(transactOpts, tc.deposits)
			r.NoError(err)
			ethBackend.Commit()

			// Prepare HTTP request
			withdrawReqBody := server.WithdrawRequest{
				UserAddress:  "0xAc0974bec39a17E36Ba4a6B4d238FF944bAcB478",
				ChainId:      1,
				Amount:       tc.amount,
				TokenAddress: "0xAc0974bec39a17E36Ba4a6B4d238FF944bAcB478",
				Signature:    "0x107c123c2415c39cd7da08ceb1a53acfb6978067fd86b0ec8e2aa0c8e673255e107c123c2415c39cd7da08ceb1a53acfb6978067fd86b0ec8e2aa0c8e673255e01",
			}
			reqBody, err := json.Marshal(withdrawReqBody)
			r.NoError(err)
			resp, err := http.Post(withdrawUrl, "application/json", bytes.NewBuffer(reqBody))
			r.NoError(err)
			defer resp.Body.Close()
			if tc.expectedErr != nil {
				errBytes := make([]byte, 1000)
				nBytes, _ := resp.Body.Read(errBytes)
				r.Greater(nBytes, 0)
				r.Equal(errors.New(string(errBytes[:nBytes])), tc.expectedErr)
				r.Equal(http.StatusInternalServerError, resp.StatusCode)
				return
			}

			ethBackend.Commit()
			r.Equal(http.StatusOK, resp.StatusCode)
			var withdrawResponse server.WithdrawResponse
			err = json.NewDecoder(resp.Body).Decode(&withdrawResponse)
			r.NoError(err)
		})
	}
}

func TestClaim(t *testing.T) {
	r := require.New(t)
	ctx := context.Background()

	adminKey, err := utils.ParseKeyFromHex(senderPrivKeyStr)
	r.NoError(err)
	pub, ok := adminKey.Public().(*ecdsa.PublicKey)
	r.True(ok)
	adminPubkey := ethcrypto.PubkeyToAddress(*pub)
	ethBackend := CreateSimulatedNode(ethtypes.GenesisAlloc{
		adminPubkey: {Balance: big.NewInt(10000000000000000)},
	})

	chainId, err := ethBackend.Client().ChainID(ctx)
	r.NoError(err)
	transactOpts, err := utils.GetTransactOptsFromPrivateKeyString(senderPrivKeyStr, chainId)
	r.NoError(err)

	mockedContracts := DeployMockedContracts(ethBackend.Client(), transactOpts)
	ethBackend.Commit()
	mteeService := mocks.NewMockTeeService(t)
	r.NoError(err)

	trustManagementRouter, err := TrustManagementRouter.NewTrustManagementRouter(
		mockedContracts.TrustManagementRouter,
		ethBackend.Client(),
	)
	r.NoError(err)

	testTransactor, err := transactor.NewTransactor(
		ethBackend.Client(),
		chainId,
		transactOpts,
		mockedContracts.TrustManagementRouter,
		trustManagementRouter,
		mteeService,
	)
	r.NoError(err)

	aavePool, err := AavePool.NewAavePool(
		mockedContracts.AavePool,
		ethBackend.Client(),
	)
	r.NoError(err)

	trustManagementProvider := onchain.NewTrustManagementProvider(
		ethBackend.Client(),
		testTransactor,
		trustManagementRouter,
		aavePool,
		&bind.CallOpts{},
	)

	serverConfig := config.HttpServerConfig{
		Port: 8082,
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

	// Make /claim HTTP request to server
	claimUrl := fmt.Sprintf("http://127.0.0.1:%d/claim", serverConfig.Port)
	claimReqBody := server.ClaimRequest{
		UserAddress:  "0xAc0974bec39a17E36Ba4a6B4d238FF944bAcB478",
		ChainId:      1,
		TokenAddress: "0xAc0974bec39a17E36Ba4a6B4d238FF944bAcB478",
		Amount:       "10000000000000000",
		Deadline:     1700000000,
		Signature:    "0x107c123c2415c39cd7da08ceb1a53acfb6978067fd86b0ec8e2aa0c8e673255e107c123c2415c39cd7da08ceb1a53acfb6978067fd86b0ec8e2aa0c8e673255e01",
	}
	reqBody, err := json.Marshal(claimReqBody)
	r.NoError(err)
	resp, err := http.Post(claimUrl, "application/json", bytes.NewBuffer(reqBody))
	r.NoError(err)
	r.Equal(http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	ethBackend.Commit()

	var claimResponse server.ClaimResponse
	err = json.NewDecoder(resp.Body).Decode(&claimResponse)
	r.NoError(err)
	fmt.Println("claimResponse", claimResponse)
	sentTx, _, err := ethBackend.Client().TransactionByHash(ctx, ethcommon.HexToHash(claimResponse.Tx))
	r.NoError(err)
	fmt.Println("sentTx", sentTx)

}
