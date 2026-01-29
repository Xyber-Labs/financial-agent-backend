package protocol

import (
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type ProtocolProvider interface {
	BuildDepositTransaction() (*ethtypes.Transaction, error)
}
