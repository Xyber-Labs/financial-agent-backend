package server

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"regexp"

	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"financial-agent-backend/config"
	"financial-agent-backend/core/onchain"
	"financial-agent-backend/core/transactor"
)

var addressRegex = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
var signatureRegex = regexp.MustCompile("^0x[0-9a-fA-F]{130}$")

type HttpAgentServer struct {
	Config                  *config.HttpServerConfig
	Transactor              *transactor.Transactor
	Gin                     *gin.Engine
	trustManagementProvider *onchain.TrustManagementProvider
	logger                  zerolog.Logger
}

func NewHttpAgentServer(
	config *config.HttpServerConfig,
	transactor *transactor.Transactor,
	trustManagementProvider *onchain.TrustManagementProvider,
) *HttpAgentServer {

	s := &HttpAgentServer{
		Config:                  config,
		Transactor:              transactor,
		Gin:                     gin.New(),
		trustManagementProvider: trustManagementProvider,
		logger:                  log.With().Str("component", "HttpAgentServer").Logger(),
	}
	s.registerHandlers()
	return s
}

func (s *HttpAgentServer) Start(ctx context.Context) error {
	addr := fmt.Sprintf(":%d", s.Config.Port)
	s.logger.Info().Str("addr", addr).Msg("HttpAgentServer: starting server")
	errCh := make(chan error, 1)
	go func() {
		errCh <- s.Gin.Run(addr)
	}()

	select {
	case err := <-errCh:
		s.logger.Error().Err(err).Msg("HttpAgentServer: server error")
		return err
	case <-ctx.Done():
		s.logger.Info().Msg("HttpAgentServer: context done")
		return ctx.Err()
	}
}

func (s *HttpAgentServer) registerHandlers() {
	// s.Gin.POST("/claim", s.claimHandler())
	s.Gin.POST("/withdraw", s.withdrawHandler())
	s.Gin.POST("/withdraw-native", s.withdrawNativeHandler())
	s.Gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

// func (s *HttpAgentServer) claimHandler() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var req ClaimRequest
// 		if err := c.ShouldBindJSON(&req); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if req.TokenAddress == "" || !addressRegex.MatchString(req.TokenAddress) {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "tokenAddress is required"})
// 			return
// 		}

// 		if req.Amount == "" || req.Amount == "0" {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "amount is required"})
// 			return
// 		}

// 		claimAmount, ok := big.NewInt(0).SetString(req.Amount, 10)
// 		if !ok {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid amount"})
// 			return
// 		}

// 		tx, err := s.trustManagementProvider.Claim(
// 			ethcommon.HexToAddress(req.TokenAddress),
// 			ethcommon.HexToAddress(req.UserAddress),
// 			claimAmount,
// 		)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(http.StatusOK, ClaimResponse{Tx: tx.Hash().String()})
// 	}
// }

// @Summary withdraw assets on behalf of user, earned rewards included
// @Schemes
// @Description withdraw the provided amount of tokens + earned rewards from the protocol and pool where the token is staked
// @Accept json
// @Produce json
// @Param request body WithdrawRequest true "Withdraw request payload"
// @Success 200 {object} WithdrawResponse
// @Router /withdraw [post]
func (s *HttpAgentServer) withdrawHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req WithdrawRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			s.logger.Error().Err(err).Msg("withdrawHandler: failed to parse request body")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		s.logger.Info().
			Str("userAddress", req.UserAddress).
			Str("tokenAddress", req.TokenAddress).
			Str("amount", req.Amount).
			Msg("withdrawHandler: received request")

		if req.UserAddress == "" || !addressRegex.MatchString(req.UserAddress) {
			s.logger.Error().Str("userAddress", req.UserAddress).Msg("withdrawHandler: userAddress is required")
			c.JSON(http.StatusBadRequest, gin.H{"error": "userAddress is required"})
			return
		}

		if req.TokenAddress == "" || !addressRegex.MatchString(req.TokenAddress) {
			s.logger.Error().Str("tokenAddress", req.TokenAddress).Msg("withdrawHandler: tokenAddress is required")
			c.JSON(http.StatusBadRequest, gin.H{"error": "tokenAddress is required"})
			return
		}

		if req.Amount == "" || req.Amount == "0" {
			s.logger.Error().Str("amount", req.Amount).Msg("withdrawHandler: amount is required")
			c.JSON(http.StatusBadRequest, gin.H{"error": "amount is required"})
			return
		}

		withdrawAmount, ok := big.NewInt(0).SetString(req.Amount, 10)
		if !ok {
			s.logger.Error().Str("amount", req.Amount).Msg("withdrawHandler: unable to convert amount to *big.Int")
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid amount"})
			return
		}

		tx, err := s.trustManagementProvider.Withdraw(
			ethcommon.HexToAddress(req.TokenAddress),
			withdrawAmount,
			ethcommon.HexToAddress(req.UserAddress),
		)
		if err != nil {
			s.logger.Error().Err(err).Msg("withdrawHandler: withdraw failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, WithdrawResponse{Tx: tx.Hash().String()})
	}
}

// @Summary withdraw native tokens on behalf of user
// @Schemes
// @Description withdraw the provided amount of native tokens from the protocol, unwrap them to native and send to user address
// @Accept json
// @Produce json
// @Param request body WithdrawNativeRequest true "Withdraw native request payload"
// @Success 200 {object} WithdrawResponse
// @Router /withdraw-native [post]
func (s *HttpAgentServer) withdrawNativeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req WithdrawNativeRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		s.logger.Info().
			Str("userAddress", req.UserAddress).
			Str("amount", req.Amount).
			Msg("withdrawNativeHandler: received request")

		if req.UserAddress == "" || !addressRegex.MatchString(req.UserAddress) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "userAddress is required"})
			return
		}

		if req.Amount == "" || req.Amount == "0" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "amount is required"})
			return
		}

		withdrawAmount, ok := big.NewInt(0).SetString(req.Amount, 10)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid amount"})
			return
		}

		tx, err := s.trustManagementProvider.WithdrawNative(withdrawAmount, ethcommon.HexToAddress(req.UserAddress))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"tx": tx.Hash().String()})
	}
}

type ClaimRequest struct {
	UserAddress  string `json:"userAddress"`
	ChainId      uint64 `json:"chainId"`
	Amount       string `json:"amount"`
	TokenAddress string `json:"tokenAddress"`
}

type WithdrawNativeRequest struct {
	UserAddress string `json:"userAddress"`
	Amount      string `json:"amount"`
}

type WithdrawRequest struct {
	UserAddress  string `json:"userAddress"`
	ChainId      uint64 `json:"chainId"`
	Amount       string `json:"amount"`
	TokenAddress string `json:"tokenAddress"`
}

type WithdrawResponse struct {
	Tx string `json:"tx"`
}

type ClaimResponse struct {
	Tx string `json:"tx"`
}

type WithdrawNativeResponse struct {
	Tx string `json:"tx"`
}
