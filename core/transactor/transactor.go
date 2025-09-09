package transactor

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/rs/zerolog/log"
	sgx_quote "github.com/Xyber-Labs/go-tee/sgx-quote"

	"financial-agent-backend/core/abi/bindings/TrustManagementRouter"
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
		Str("quote", hex.EncodeToString(quote)).
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

	log.Info().
		Str("proof.leaf.bodyPartOne", hex.EncodeToString(teeXyberProof.Leaf.BodyPartOne)).
		Str("proof.leaf.publicKey", hex.EncodeToString(teeXyberProof.Leaf.PublicKey)).
		Str("proof.leaf.bodyPartTwo", hex.EncodeToString(teeXyberProof.Leaf.BodyPartTwo)).
		Str("proof.leaf.signature", hex.EncodeToString(teeXyberProof.Leaf.Signature)).
		Str("proof.intermediate.bodyPartOne", hex.EncodeToString(teeXyberProof.Intermediate.BodyPartOne)).
		Str("proof.intermediate.publicKey", hex.EncodeToString(teeXyberProof.Intermediate.PublicKey)).
		Str("proof.intermediate.bodyPartTwo", hex.EncodeToString(teeXyberProof.Intermediate.BodyPartTwo)).
		Str("proof.intermediate.signature", hex.EncodeToString(teeXyberProof.Intermediate.Signature)).
		Str("proof.quote.Header", hex.EncodeToString(teeXyberProof.Quote.Header)).
		Str("proof.quote.IsvReport", hex.EncodeToString(teeXyberProof.Quote.IsvReport)).
		Str("proof.quote.IsvReportSignature", hex.EncodeToString(teeXyberProof.Quote.IsvReportSignature)).
		Str("proof.quote.AttestationKey", hex.EncodeToString(teeXyberProof.Quote.AttestationKey)).
		Str("proof.quote.QeReport", hex.EncodeToString(teeXyberProof.Quote.QeReport)).
		Str("proof.quote.QeReportSignature", hex.EncodeToString(teeXyberProof.Quote.QeReportSignature)).
		Str("proof.quote.QeAuthenticationData", hex.EncodeToString(teeXyberProof.Quote.QeAuthenticationData)).
		Msg("TeeSession: Prepared quote for initSessionKey transaction")

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

	deadline := big.NewInt(time.Now().Add(10 * time.Minute).Unix())
	signature, err := t.createTeeSessionSignature(deadline, transactionsArg)
	if err != nil {
		return nil, fmt.Errorf("failed to create tee session signature: %w", err)
	}

	tx, err := t.TrustManagementRouter.Execute(t.TransactOpts, transactionsArg, signature, deadline)
	if err != nil {
		log.Info().Str("error", err.Error()).Msg("Transactor.BatchAndExecute: failed to send transaction, trying with custom gas limit")
		// If tx failed, try again with custom gas limit
		opts := *t.TransactOpts
		opts.GasLimit = 5000000
		tx, err = t.TrustManagementRouter.Execute(&opts, transactionsArg, signature, deadline)
		if err != nil {
			return nil, fmt.Errorf("failed to send transaction: %w", err)
		}
	}

	log.Info().Str("tx", tx.Hash().String()).Int("tx_count", len(innerTxs)).Msg("Transactor.BatchAndExecute: sent batched transaction")
	return tx, nil
}

func (t *Transactor) extractQuote(sessionKey ethcommon.Address) ([]byte, error) {
	// On-chain contract expects userData to be `keccak256(abi.encodePacked(sessionKey))`.
	// Convert it in that format here.
	userData := ethcrypto.Keccak256(sessionKey[:])
	return t.TeeService.GetQuote(userData)
}

// Generates signature using t.sessionKey with the following message format:
// keccak256(abi.encode(block.chainid, address(router), deadline, txs))
func (t *Transactor) createTeeSessionSignature(
	deadline *big.Int,
	transactions []TrustManagementRouter.Transaction,
) ([]byte, error) {
	sig, _, err := CreateTeeSessionSignature(t.ChainId, t.TeeSessionKey, t.TrustManagementRouterAddress, deadline, transactions)
	if err != nil {
		return nil, err
	}
	return sig, nil
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

func CreateTeeSessionSignature(
	chainId *big.Int,
	teeSessionKey *ecdsa.PrivateKey,
	trustManagementRouterAddress ethcommon.Address,
	deadline *big.Int,
	transactions []TrustManagementRouter.Transaction,
) ([]byte, []byte, error) {
	if teeSessionKey == nil {
		return nil, nil, fmt.Errorf("createTeeSessionSignature: TeeSessionKey is not initialized")
	}
	if chainId == nil {
		return nil, nil, fmt.Errorf("createTeeSessionSignature: ChainId is not set")
	}
	if deadline == nil {
		return nil, nil, fmt.Errorf("createTeeSessionSignature: Deadline is not set")
	}
	if transactions == nil {
		return nil, nil, fmt.Errorf("createTeeSessionSignature: Transactions is not set")
	}

	// Prepare ABI types: (uint256 chainId, address router, uint256 deadline, tuple(address,uint256,bytes)[] txs)
	uint256Type, err := abi.NewType("uint256", "", nil)
	if err != nil {
		return nil, nil, fmt.Errorf("createTeeSessionSignature: failed to create uint256 ABI type: %w", err)
	}
	addressType, err := abi.NewType("address", "", nil)
	if err != nil {
		return nil, nil, fmt.Errorf("createTeeSessionSignature: failed to create address ABI type: %w", err)
	}
	txTupleArrayType, err := abi.NewType("tuple[]", "", []abi.ArgumentMarshaling{
		{Name: "target", Type: "address"},
		{Name: "value", Type: "uint256"},
		{Name: "data", Type: "bytes"},
	})
	if err != nil {
		return nil, nil, fmt.Errorf("createTeeSessionSignature: failed to create tx tuple[] ABI type: %w", err)
	}

	args := abi.Arguments{
		{Type: uint256Type},      // chainId
		{Type: addressType},      // router address
		{Type: uint256Type},      // deadline
		{Type: txTupleArrayType}, // transactions
	}

	packed, err := args.Pack(chainId, trustManagementRouterAddress, deadline, transactions)
	if err != nil {
		return nil, nil, fmt.Errorf("createTeeSessionSignature: ABI pack failed: %w", err)
	}
	fmt.Println("packedData:", hex.EncodeToString(packed))

	// Solidity: keccak256(abi.encode(chainid, address(router), deadline, txs))
	msgHash := ethcrypto.Keccak256(packed)
	fmt.Println("hashedData:", hex.EncodeToString(msgHash))

	signature, err := ethcrypto.Sign(msgHash, teeSessionKey)
	if err != nil {
		return nil, nil, fmt.Errorf("createTeeSessionSignature: sign failed: %w", err)
	}

	return signature, msgHash, nil
}
