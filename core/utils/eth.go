package utils

import "github.com/ethereum/go-ethereum/ethclient"


func ConnectToNode(rpcURL string) (*ethclient.Client, error) {
	return ethclient.Dial(rpcURL)
}