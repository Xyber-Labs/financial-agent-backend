package tests

import (
	"testing"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

func TestDeposit(t *testing.T) {
	r := require.New(t)

	ethBackend := CreateSimulatedNode(ethtypes.GenesisAlloc{})

	r.True(false)
}
