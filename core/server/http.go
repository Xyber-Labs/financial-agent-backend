package server

import (
	"fmt"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"financial-agent-backend/config"
	"financial-agent-backend/core/abi/bindings/AavePool"
	"financial-agent-backend/core/abi/bindings/TrustManagementRouter"
	"financial-agent-backend/core/transactor"
)

type HttpAgentServer struct {
	Config                *config.HttpServerConfig
	Transactor            *transactor.Transactor
	TrustManagementRouter *TrustManagementRouter.TrustManagementRouter
	AavePool              *AavePool.AavePool
	createTxOpts          *bind.TransactOpts
	gin                   *gin.Engine
}

func NewHttpAgentServer(
	config *config.HttpServerConfig,
	transactor *transactor.Transactor,
	trustManagementRouter *TrustManagementRouter.TrustManagementRouter,
	aavePool *AavePool.AavePool,
) *HttpAgentServer {
	// createTxOpts only used to consturct the calldata for the transactions
	// to be later reconstructed in a batch transaction. Therefore we don't actually
	// want to sign, estimate or execute them.
	createTxOpts := bind.TransactOpts{
		NoSend: true,
	}
	s := &HttpAgentServer{
		Config:                config,
		Transactor:            transactor,
		TrustManagementRouter: trustManagementRouter,
		AavePool:              aavePool,
		createTxOpts:          &createTxOpts,
		gin:                   gin.New(),
	}
	s.registerHandlers()
	return s
}

func (s *HttpAgentServer) Start() error {
	addr := fmt.Sprintf(":%d", s.Config.Port)
	log.Info().Str("addr", addr).Msg("HttpAgentServer: starting server")
	return s.gin.Run(addr)
}

func (s *HttpAgentServer) registerHandlers() {
	s.gin.POST("/deposit", s.depositHandler())
	s.gin.POST("/claim", s.claimHandler())
	s.gin.POST("/withdraw", s.withdrawHandler())
}

// Deposit the tokens for the best-performing pool on behalf of the user.
// This is handled the following way:
// 1. Ask agent where to put it (AAVE current supported)
// 2. Create and send batched transaction that executes following:
// 2.1. Calls TrustManagementRouter.depositWithPermit
// 2.2. Calls TrustManagemtnRouter.execute with encoded AAVE deposit transaction
func (s *HttpAgentServer) depositHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req DepositRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Prepare TrustManagementRouter.depositWithPermit transaction
		userAddress := ethcommon.HexToAddress(req.UserAddress)
		tokenAddress := ethcommon.HexToAddress(req.TokenAddress)
		tokenAmount, ok := big.NewInt(0).SetString(req.Amount, 10)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid amount"})
			return
		}
		sigR := ethcommon.HexToHash(req.SigR)
		sigS := ethcommon.HexToHash(req.SigS)
		depositWithPermitTx, err := s.TrustManagementRouter.DepositWithPermit(
			s.createTxOpts,
			userAddress,
			tokenAddress,
			big.NewInt(int64(req.Deadline)),
			TrustManagementRouter.ITrustManagementStructsPermit{
				Amount:   tokenAmount,
				Deadline: big.NewInt(int64(req.Deadline)),
				V:        req.SigV,
				R:        sigR,
				S:        sigS,
			},
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Create AavePool.supply transaction
		aaveSupplyTx, err := s.AavePool.Supply(
			s.createTxOpts,
			tokenAddress,
			tokenAmount,
			userAddress,
			0,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Batch and execute transactions
		tx, err := s.Transactor.BatchAndExecute([]*ethtypes.Transaction{depositWithPermitTx, aaveSupplyTx})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"tx": tx.Hash().String()})
	}
}

// Claims the incrued rewards on behalf of the user
// Claim works the following way:
// 1.
func (s *HttpAgentServer) claimHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ClaimRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
}

func (s *HttpAgentServer) withdrawHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req WithdrawRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
}

type DepositRequest struct {
	UserAddress  string `json:"userAddress"`
	ChainId      uint64 `json:"chainId"`
	TokenAddress string `json:"tokenAddress"`
	Amount       string `json:"amount"`
	Deadline     uint64 `json:"deadline"`
	SigV         uint8  `json:"sigV"`
	SigR         string `json:"sigR"`
	SigS         string `json:"sigS"`
}

type ClaimRequest struct {
	ChainId uint64 `json:"chainId"`
}

type WithdrawRequest struct {
	ChainId uint64 `json:"chainId"`
}
