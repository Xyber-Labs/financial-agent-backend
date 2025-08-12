package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"financial-agent-backend/config"
)

type HttpAgentServer struct {
	Config *config.HttpServerConfig
	gin    *gin.Engine
}

func NewHttpAgentServer(config *config.HttpServerConfig) *HttpAgentServer {
	s := &HttpAgentServer{
		Config: config,
		gin:    gin.New(),
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

func (s *HttpAgentServer) depositHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req DepositRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
}

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
	ChainId           uint64 `json:"chainId"`
	InputTokenAddress string `json:"inputTokenAddress"`
	Amount            string `json:"amount"`
}

type ClaimRequest struct {
	ChainId uint64 `json:"chainId"`
}

type WithdrawRequest struct {
	ChainId uint64 `json:"chainId"`
}
