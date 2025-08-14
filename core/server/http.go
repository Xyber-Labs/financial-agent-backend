package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"financial-agent-backend/config"
	"financial-agent-backend/core/transactor"
)

type HttpAgentServer struct {
	Config     *config.HttpServerConfig
	Transactor *transactor.Transactor
	gin        *gin.Engine
}

func NewHttpAgentServer(config *config.HttpServerConfig, transactor *transactor.Transactor) *HttpAgentServer {
	s := &HttpAgentServer{
		Config:     config,
		Transactor: transactor,
		gin:        gin.New(),
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
