package server

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"regexp"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"financial-agent-backend/config"
	"financial-agent-backend/core/abi/bindings/TrustManagementRouter"
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
	}
	s.registerHandlers()
	return s
}

func (s *HttpAgentServer) Start(ctx context.Context) error {
	addr := fmt.Sprintf(":%d", s.Config.Port)
	log.Info().Str("addr", addr).Msg("HttpAgentServer: starting server")
	errCh := make(chan error, 1)
	go func() {
		errCh <- s.Gin.Run(addr)
	}()

	select {
	case err := <-errCh:
		log.Error().Err(err).Msg("HttpAgentServer: server error")
		return err
	case <-ctx.Done():
		log.Info().Msg("HttpAgentServer: context done")
		return ctx.Err()
	}
}

func (s *HttpAgentServer) registerHandlers() {
	s.Gin.POST("/deposit", s.depositHandler())
	s.Gin.POST("/claim", s.claimHandler())
	s.Gin.POST("/withdraw", s.withdrawHandler())
}

// Deposit the tokens for the best-performing pool on behalf of the user.
// This is handled the following way:
// 1. Ask agent where to put it (AAVE current supported)
// 2. Create and send batched transaction that executes following:
// 2.1. Calls TrustManagementRouter.depositWithPermit
// 2.2. Calls TrustManagemtnRouter.execute with encoded AAVE deposit transaction
func (s *HttpAgentServer) depositHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Avoid logging the entire *http.Request as it contains function fields (e.g., GetBody)
		// which are not JSON serializable and cause marshaling errors.
		log.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("content_type", c.ContentType()).
			Msg("depositHandler: received request")
		var req DepositRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Error().Err(err).Msg("depositHandler: failed to bind request")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.UserAddress == "" || !addressRegex.MatchString(req.UserAddress) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "userAddress is required"})
			return
		}

		if req.TokenAddress == "" || !addressRegex.MatchString(req.TokenAddress) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "tokenAddress is required"})
			return
		}

		if req.Amount == "" || req.Amount == "0" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "amount is required"})
			return
		}

		if req.Deadline == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "deadline is required"})
			return
		}

		if req.SigV == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "sigV is required"})
			return
		}

		if req.SigR == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "sigR is required"})
			return
		}

		if req.SigS == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "sigS is required"})
			return
		}

		log.Info().Msg("checks done, parsing request")

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

		log.Info().
			Interface("userAddress", userAddress).
			Interface("tokenAddress", tokenAddress).
			Interface("tokenAmount", tokenAmount).
			Interface("deadline", req.Deadline).
			Interface("sigV", req.SigV).
			Interface("sigR", sigR).
			Interface("sigS", sigS).
			Msg("Creating TrustManagementRouter.depositWithPermit transaction")

		tx, err := s.trustManagementProvider.Deposit(
			userAddress,
			tokenAddress,
			tokenAmount,
			big.NewInt(int64(req.Deadline)),
			TrustManagementRouter.Permit{
				Holder:   userAddress,
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

		c.JSON(http.StatusOK, DepositResponse{Tx: tx.Hash().String()})
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

		if req.TokenAddress == "" || !addressRegex.MatchString(req.TokenAddress) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "tokenAddress is required"})
			return
		}

		if req.Amount == "" || req.Amount == "0" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "amount is required"})
			return
		}

		if req.Deadline == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "deadline is required"})
			return
		}

		if req.Signature == "" || !signatureRegex.MatchString(req.Signature) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "deadline is required"})
			return
		}

		claimAmount, ok := big.NewInt(0).SetString(req.Amount, 10)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid amount"})
			return
		}
		deadline := big.NewInt(int64(req.Deadline))
		signature := ethcommon.FromHex(req.Signature)

		tx, err := s.trustManagementProvider.Claim(
			ethcommon.HexToAddress(req.TokenAddress),
			ethcommon.HexToAddress(req.UserAddress),
			claimAmount,
			signature,
			deadline,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, ClaimResponse{Tx: tx.Hash().String()})
	}
}

func (s *HttpAgentServer) withdrawHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info().Msg("withdrawHandler: received request")
		var req WithdrawRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Error().Err(err).Msg("withdrawHandler: failed to parse request body")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.UserAddress == "" || !addressRegex.MatchString(req.UserAddress) {
			log.Error().Str("userAddress", req.UserAddress).Msg("withdrawHandler: userAddress is required")
			c.JSON(http.StatusBadRequest, gin.H{"error": "userAddress is required"})
			return
		}

		if req.TokenAddress == "" || !addressRegex.MatchString(req.TokenAddress) {
			log.Error().Str("tokenAddress", req.TokenAddress).Msg("withdrawHandler: tokenAddress is required")
			c.JSON(http.StatusBadRequest, gin.H{"error": "tokenAddress is required"})
			return
		}

		if req.Amount == "" || req.Amount == "0" {
			log.Error().Str("amount", req.Amount).Msg("withdrawHandler: amount is required")
			c.JSON(http.StatusBadRequest, gin.H{"error": "amount is required"})
			return
		}

		if req.Signature == "" || !signatureRegex.MatchString(req.Signature) {
			log.Error().Str("signature", req.Signature).Msg("withdrawHandler: signature is required")
			c.JSON(http.StatusBadRequest, gin.H{"error": "signature is required"})
			return
		}

		withdrawAmount, ok := big.NewInt(0).SetString(req.Amount, 10)
		if !ok {
			log.Error().Str("amount", req.Amount).Msg("withdrawHandler: unable to convert amount to *big.Int")
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid amount"})
			return
		}

		// Parse signature, it should be in the format of 0x[a-ZA-Z0-9]{130}
		signature := ethcommon.FromHex(req.Signature)

		if len(signature) != 65 {
			log.Error().Str("signature", req.Signature).Msg("withdrawHandler: signature must be 65 bytes long")
			c.JSON(http.StatusBadRequest, gin.H{"error": "signature must be 65 bytes long"})
			return
		}

		deadline := big.NewInt(time.Now().Unix())
		tx, err := s.trustManagementProvider.Withdraw(
			ethcommon.HexToAddress(req.TokenAddress),
			withdrawAmount,
			ethcommon.HexToAddress(req.UserAddress),
			deadline,
			signature,
		)
		if err != nil {
			log.Error().Err(err).Msg("withdrawHandler: withdraw failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, WithdrawResponse{Tx: tx.Hash().String()})
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

type DepositResponse struct {
	Tx string `json:"tx"`
}

type ClaimRequest struct {
	UserAddress  string `json:"userAddress"`
	ChainId      uint64 `json:"chainId"`
	Amount       string `json:"amount"`
	TokenAddress string `json:"tokenAddress"`
	Signature    string `json:"signature"`
	Deadline     uint64 `json:"deadline"`
}

type WithdrawRequest struct {
	UserAddress  string `json:"userAddress"`
	ChainId      uint64 `json:"chainId"`
	Amount       string `json:"amount"`
	TokenAddress string `json:"tokenAddress"`
	Signature    string `json:"signature"`
}

type WithdrawResponse struct {
	Tx string `json:"tx"`
}

type ClaimResponse struct {
	Tx string `json:"tx"`
}
