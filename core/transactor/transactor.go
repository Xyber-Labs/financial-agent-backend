package transactor

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"

	sgx_quote "github.com/Xyber-Labs/go-tee/sgx-quote"
)

type Transactor struct {
	client *ethclient.Client
	teeService *sgx_quote.TeeService
	teeSessionKey *ecdsa.PrivateKey
	teeSessionAddress common.Address
}

func NewTransactor(client *ethclient.Client) *Transactor {
	teeService := sgx_quote.NewTeeService(false)
	return &Transactor{
		client: client,
		teeService: teeService,
	}
}

// In transaction initialization we perform quote extraction
func (t *Transactor) InitializeOnChainSession() error {

	sessionKey, err := GenerateTeeSessionKey()
	if err != nil {
		return err
	}
	t.teeSessionKey = sessionKey
	t.teeSessionAddress = crypto.PubkeyToAddress(sessionKey.PublicKey)
	log.Info().
		Str("address", t.teeSessionAddress.String()).
		Msg("NewTeeSession: generated session key for TeeSession")

	quote, err := t.extractQuote(t.teeSessionAddress)
	if err != nil {
		return err
	}
	log.Info().
		Str("quote", string(quote)).
		Msg("NewTeeSession: extracted quote for TeeSession")
	
	return nil
}

func (t *Transactor) extractQuote(sessionKey common.Address) ([]byte, error) {
	return t.teeService.GetQuote(sessionKey[:])
}

// Creates a new ECDSA private key using the secp256k1 curve
// which is compatible with Ethereum's ecrecover function
func GenerateTeeSessionKey() (*ecdsa.PrivateKey, error) {
	// Generate a new private key using the secp256k1 curve
	privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate ECDSA key: %w", err)
	}

	return privateKey, nil
}
