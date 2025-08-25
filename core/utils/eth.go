package utils

import (
	"crypto/ecdsa"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ConnectToNode(rpcURL string) (*ethclient.Client, error) {
	return ethclient.Dial(rpcURL)
}

func GetTransactOptsFromPrivateKeyString(privateKeyStr string, chainId *big.Int) (*bind.TransactOpts, error) {
	privateKeyStr = strings.TrimPrefix(privateKeyStr, "0x")
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, err
	}
	return bind.NewKeyedTransactorWithChainID(privateKey, chainId)
}

// ParseKeyFromHex parses a private key from a hex string
func ParseKeyFromHex(privateKeyHex string) (*ecdsa.PrivateKey, error) {
	// Decode hex to bytes
	privateKeyBytes, err := hexutil.Decode(privateKeyHex)
	if err != nil {
		return nil, err
	}

	// Load prviate key from bytes
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}
