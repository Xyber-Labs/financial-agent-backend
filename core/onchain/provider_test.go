package onchain

import (
	"math/big"
	"testing"
)

func TestCalculateAmountWithYield(t *testing.T) {

	testCases := []struct {
		desc                    string
		amount                  *big.Int
		totalDepositAmount      *big.Int
		aTokenBalance           *big.Int
		expectedAmountWithYield *big.Int
		expectedErr             error
	}{
		{
			desc:                    "user withdraws 7k when total deposit is 10k and aTokenBalance is 11k",
			amount:                  big.NewInt(7000),
			totalDepositAmount:      big.NewInt(10000),
			aTokenBalance:           big.NewInt(11000),
			expectedAmountWithYield: big.NewInt(8000), // 7k + 1k yield
			expectedErr:             nil,
		},
		{
			desc:                    "user withdraws 10k when total deposit is 10k and aTokenBalance is 10.5k",
			amount:                  big.NewInt(10000),
			totalDepositAmount:      big.NewInt(10000),
			aTokenBalance:           big.NewInt(10500),
			expectedAmountWithYield: big.NewInt(10500), // 10k + 500 yield
			expectedErr:             nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			amountWithYield, err := calculateAmountWithYield(tc.amount, tc.totalDepositAmount, tc.aTokenBalance)
			if tc.expectedErr != nil {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				if err.Error() != tc.expectedErr.Error() {
					t.Errorf("expected error %s, got %s", tc.expectedErr, err)
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}
			if amountWithYield.Cmp(tc.expectedAmountWithYield) != 0 {
				t.Errorf("expected amount with yield %s, got %s", tc.expectedAmountWithYield, amountWithYield)
			}
		})
	}

}
