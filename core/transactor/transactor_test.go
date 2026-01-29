package transactor

import (
	"errors"
	"slices"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestInitialization(t *testing.T) {
	t.Skip("TODO: add unit tests")
}

// Tests for InitializeOnChainSession using mockery-generated MockTeeService
func TestInitializeOnChainSession(t *testing.T) {
	testCases := []struct {
		name        string
		getQuoteRes []byte
		getQuoteErr error
		wantErr     bool
	}{
		{
			name:        "success: generates session, calls GetQuote with address bytes",
			getQuoteRes: []byte("quote"),
			getQuoteErr: nil,
			wantErr:     false,
		},
		{
			name:        "error: propagates GetQuote error",
			getQuoteRes: nil,
			getQuoteErr: errors.New("boom"),
			wantErr:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			// Arrange
			mteeService := NewMockTeeService(t)
			var captured []byte
			mteeService.EXPECT().GetQuote(mock.Anything).Run(func(b []byte) {
				// capture argument for later comparison
				if b != nil {
					captured = slices.Clone(b)
				}
			}).Return(tc.getQuoteRes, tc.getQuoteErr)

			// Build Transactor via constructor with minimal deps
			client := &ethclient.Client{}
			opts := &bind.TransactOpts{From: ethcommon.Address{}}
			routerAddr := ethcommon.Address{0x1}
			walletAddr := ethcommon.Address{0x2}
			transactor, nerr := NewTransactor(client, opts, routerAddr, walletAddr, mteeService)
			r.NoError(nerr)

			// Assert
			err := transactor.InitializeOnChainSession()
			if tc.wantErr {
				r.Error(err)
				return
			}
			r.NoError(err)
			r.NotNil(transactor.teeSessionKey)
			r.NotEqualValues([20]byte{}, transactor.teeSessionAddress)
			r.Equal(transactor.teeSessionAddress[:], captured)
		})
	}
}
