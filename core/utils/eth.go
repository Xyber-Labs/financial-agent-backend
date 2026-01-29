package utils

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
