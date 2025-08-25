package transactor

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/rs/zerolog/log"

	"financial-agent-backend/core/abi/bindings/TrustManagementRouter"

	sgx_quote "github.com/Xyber-Labs/go-tee/sgx-quote"
)

type Transactor struct {
	Client                       bind.ContractBackend
	ChainId                      *big.Int
	TeeService                   TeeService
	TeeSessionKey                *ecdsa.PrivateKey
	TeeSessionAddress            ethcommon.Address
	TransactOpts                 *bind.TransactOpts
	TrustManagementRouterAddress ethcommon.Address
	TrustManagementRouter        *TrustManagementRouter.TrustManagementRouter
}

func NewTransactor(
	client bind.ContractBackend,
	chainId *big.Int,
	transactOpts *bind.TransactOpts,
	trustManagementRouterAddress ethcommon.Address,
	trustManagementRouter *TrustManagementRouter.TrustManagementRouter,
	teeService TeeService,
) (*Transactor, error) {
	return &Transactor{
		Client:                       client,
		ChainId:                      chainId,
		TeeService:                   teeService,
		TransactOpts:                 transactOpts,
		TrustManagementRouterAddress: trustManagementRouterAddress,
		TrustManagementRouter:        trustManagementRouter,
	}, nil
}

// In transaction initialization we perform quote extraction
func (t *Transactor) InitializeOnChainSession() error {

	// First generate tee session key
	sessionKey, err := GenerateTeeSessionKey()
	if err != nil {
		return err
	}
	t.TeeSessionKey = sessionKey
	t.TeeSessionAddress = ethcrypto.PubkeyToAddress(sessionKey.PublicKey)
	log.Info().
		Str("address", t.TeeSessionAddress.String()).
		Msg("NewTeeSession: generated session key for TeeSession")

	// Extract SGX quote from the environment
	quote, err := t.extractQuote(t.TeeSessionAddress)
	if err != nil {
		return err
	}
	log.Info().
		Str("quote", string(quote)).
		Str("quoteHex", hex.EncodeToString(quote)).
		Msg("NewTeeSession: extracted quote for TeeSession")

	// Prepare SGX quote
	sgxParser := sgx_quote.NewSgxParser()
	parsedQuote, err := sgxParser.ParseQuote(quote)
	if err != nil {
		return fmt.Errorf("TeeSession: failed to parse quote: %w", err)
	}

	teeXyberProof, err := FromSgxQuoteToProof(parsedQuote)
	if err != nil {
		return fmt.Errorf("TeeSession: failed to convert quote to tee wallet arguments: %w", err)
	}

	// Initialize onchain tee session session
	tx, err := t.TrustManagementRouter.InitSessionKey(
		t.TransactOpts,
		teeXyberProof.Leaf,
		teeXyberProof.Intermediate,
		teeXyberProof.Quote,
		t.TeeSessionAddress,
	)
	if err != nil {
		return fmt.Errorf("TeeSession: failed to send transaction: %w", err)
	}

	log.Info().Str("tx", tx.Hash().String()).Msg("InitializeOnChainSession: registered in TrustManagementRouter")

	return nil
}

// BatchAndExecute function batches provided transactions into TrustManagementRouter.execute multicall
// And sends it to the network.
func (t *Transactor) BatchAndExecute(innerTxs []*ethtypes.Transaction) (*ethtypes.Transaction, error) {

	// Prepare arguments for TrustManagementRouter.execute multicall
	transactionsArg := make([]TrustManagementRouter.Transaction, len(innerTxs))
	for i, tx := range innerTxs {
		if tx.To() == nil {
			return nil, fmt.Errorf("BatchAndExecute: transaction %d has nil To field", i)
		}
		transactionsArg[i] = TrustManagementRouter.Transaction{
			Target: *tx.To(),
			Value:  tx.Value(),
			Data:   tx.Data(),
		}
	}

	deadline := time.Now().Add(10 * time.Minute)
	signature, err := t.createTeeSessionSignature(big.NewInt(deadline.Unix()), transactionsArg)
	if err != nil {
		return nil, fmt.Errorf("failed to create tee session signature: %w", err)
	}

	tx, err := t.TrustManagementRouter.Execute(t.TransactOpts, transactionsArg, signature, big.NewInt(deadline.Unix()))
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %w", err)
	}

	log.Info().Str("tx", tx.Hash().String()).Int("tx_count", len(innerTxs)).Msg("Transactor.BatchAndExecute: sent batched transaction")
	return tx, nil
}

func (t *Transactor) extractQuote(sessionKey ethcommon.Address) ([]byte, error) {
	return t.TeeService.GetQuote(sessionKey[:])
	// return []byte("quote"), nil
}

// Generates signature using t.sessionKey with the following message format:
// keccak256(abi.encode(block.chainid, address(router), deadline, txs))
func (t *Transactor) createTeeSessionSignature(
	deadline *big.Int,
	transactions []TrustManagementRouter.Transaction,
) ([]byte, error) {
	return []byte{}, nil
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
