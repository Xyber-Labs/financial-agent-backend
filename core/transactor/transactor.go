package transactor

import "github.com/ethereum/go-ethereum/ethclient"

type Transactor struct {
	client *ethclient.Client
}

func NewTransactor(client *ethclient.Client) *Transactor {
	return &Transactor{
		client: client,
	}
}

func (t *Transactor) Initialize() error {
	
}