package transactor

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"

	"financial-agent-backend/core/abi/bindings/TEEWallet"
	"financial-agent-backend/core/abi/bindings/TrustManagementRouter"

	sgx_quote "github.com/Xyber-Labs/go-tee/sgx-quote"
)

type Transactor struct {
	client                *ethclient.Client
	teeService            TeeService
	teeSessionKey         *ecdsa.PrivateKey
	teeSessionAddress     ethcommon.Address
	transactOpts          *bind.TransactOpts
	TrustManagementRouter *TrustManagementRouter.TrustManagementRouterCaller
	TEEWallet             *TEEWallet.TEEWallet
}

func NewTransactor(
	client *ethclient.Client,
	transactOpts *bind.TransactOpts,
	trustManagementRouterAddress ethcommon.Address,
	teeWalletAddress ethcommon.Address,
	teeService TeeService,
) (*Transactor, error) {
	trustManagementRouter, err := TrustManagementRouter.NewTrustManagementRouterCaller(
		trustManagementRouterAddress,
		client,
	)
	if err != nil {
		return nil, err
	}
	teeWallet, err := TEEWallet.NewTEEWallet(
		teeWalletAddress,
		client,
	)
	if err != nil {
		return nil, err
	}
	return &Transactor{
		client:                client,
		teeService:            teeService,
		transactOpts:          transactOpts,
		TrustManagementRouter: trustManagementRouter,
		TEEWallet:             teeWallet,
	}, nil
}

// In transaction initialization we perform quote extraction
func (t *Transactor) InitializeOnChainSession() error {

	// First generate tee session key
	sessionKey, err := GenerateTeeSessionKey()
	if err != nil {
		return err
	}
	t.teeSessionKey = sessionKey
	t.teeSessionAddress = ethcrypto.PubkeyToAddress(sessionKey.PublicKey)
	log.Info().
		Str("address", t.teeSessionAddress.String()).
		Msg("NewTeeSession: generated session key for TeeSession")

	// Extract SGX quote from the environment
	quote, err := t.extractQuote(t.teeSessionAddress)
	if err != nil {
		return err
	}
	log.Info().
		Str("quote", string(quote)).
		Str("quoteHex", hex.EncodeToString(quote)).
		Msg("NewTeeSession: extracted quote for TeeSession")

	// Prepare SGX quote for use in TEEWallet
	sgxParser := sgx_quote.NewSgxParser()
	parsedQuote, err := sgxParser.ParseQuote(quote)
	if err != nil {
		return fmt.Errorf("TeeSession: failed to parse quote: %w", err)
	}

	teeWaletProof, err := FromSgxQuoteToTeeWalletProof(parsedQuote)
	if err != nil {
		return fmt.Errorf("TeeSession: failed to convert quote to tee wallet arguments: %w", err)
	}

	// Initialize session in TEEWallet contract
	tx, err := t.TEEWallet.InitSessionKey(
		t.transactOpts,
		teeWaletProof.Leaf,
		teeWaletProof.Intermediate,
		teeWaletProof.Quote,
		t.teeSessionAddress,
	)
	if err != nil {
		return fmt.Errorf("TeeSession: failed to send transaction: %w", err)
	}

	log.Info().Str("tx", tx.Hash().String()).Msg("InitializeOnChainSession: registered in TEEWallet")

	return nil
}

// SignAndSendTransaction function is meant to be used by main backend logic to send transactions
// to the network.
func (t *Transactor) SignAndSendTransaction(tx *ethtypes.Transaction) error {
	signedTx, err := t.transactOpts.Signer(t.transactOpts.From, tx)
	if err != nil {
		return err
	}
	// FIXME(ak): add proper session key signature
	return t.client.SendTransaction(context.Background(), signedTx)
}

func (t *Transactor) extractQuote(sessionKey ethcommon.Address) ([]byte, error) {
	return t.teeService.GetQuote(sessionKey[:])
	// return []byte("quote"), nil
}

// Creates a new ECDSA private key using the secp256k1 curve
// which is compatible with Ethereum's ecrecover function
func GenerateTeeSessionKey() (*ecdsa.PrivateKey, error) {
	// Generate a new private key using the secp256k1 curve
	privateKey, err := ecdsa.GenerateKey(ethcrypto.S256(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate ECDSA key: %w", err)
	}

	return privateKey, nil
}
