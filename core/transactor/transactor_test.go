package transactor

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestInitialization(t *testing.T) {
	t.Skip("TODO: add unit tests")
}

// Tests for InitializeOnChainSession using mockery-generated MockTeeService
func TestInitializeOnChainSession(t *testing.T) {
    type testCase struct {
        name     string
        retQuote []byte
        retErr   error
        wantErr  bool
    }

    testCases := []testCase{
        {
            name:     "success: generates session, calls GetQuote with address bytes",
            retQuote: []byte("quote"),
            retErr:   nil,
            wantErr:  false,
        },
        {
            name:     "error: propagates GetQuote error",
            retQuote: nil,
            retErr:   errors.New("boom"),
            wantErr:  true,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            r := require.New(t)

            // Arrange
            mtee := NewMockTeeService(t)
            var captured []byte
            mtee.EXPECT().GetQuote(mock.Anything).Run(func(b []byte) {
                // capture argument for later comparison
                if b != nil {
                    captured = append([]byte(nil), b...)
                }
            }).Return(tc.retQuote, tc.retErr)

            subject := &Transactor{teeService: mtee}

            // Act
            err := subject.InitializeOnChainSession()

            // Assert
            if tc.wantErr {
                r.Error(err)
                return
            }
            r.NoError(err)
            r.NotNil(subject.teeSessionKey)
            // Ensure the address is set and was used as input to GetQuote
            r.NotEqualValues([20]byte{}, subject.teeSessionAddress)
            r.Equal(subject.teeSessionAddress[:], captured)
        })
    }
}
