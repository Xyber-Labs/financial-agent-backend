package onchain

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
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

func TestCalculateWithdrawAmounts(t *testing.T) {
	testCases := []struct {
		desc                    string
		userWithdrawAmount      *big.Int
		unlockedDepositAmount   *big.Int
		aTokenBalance           *big.Int
		expectedBaseAmount      *big.Int
		expectedAmountWithYield *big.Int
		expectedErr             error
	}{
		{
			desc:                    "withdraw higher than yield: deposit 100, yield 20, withdraw 50 => base 30, withYield 50",
			userWithdrawAmount:      big.NewInt(50),
			unlockedDepositAmount:   big.NewInt(100),
			aTokenBalance:           big.NewInt(120),
			expectedBaseAmount:      big.NewInt(30),
			expectedAmountWithYield: big.NewInt(50),
			expectedErr:             nil,
		},
		{
			desc:                    "withdraw lower than yield: deposit 100, yield 20, withdraw 15 => base 0, withYield 15",
			userWithdrawAmount:      big.NewInt(15),
			unlockedDepositAmount:   big.NewInt(100),
			aTokenBalance:           big.NewInt(120),
			expectedBaseAmount:      big.NewInt(0),
			expectedAmountWithYield: big.NewInt(15),
			expectedErr:             nil,
		},
		{
			desc:                    "withdraw equal to yield: deposit 100, yield 20, withdraw 20 => base 0, withYield 20",
			userWithdrawAmount:      big.NewInt(20),
			unlockedDepositAmount:   big.NewInt(100),
			aTokenBalance:           big.NewInt(120),
			expectedBaseAmount:      big.NewInt(0),
			expectedAmountWithYield: big.NewInt(20),
			expectedErr:             nil,
		},
		{
			desc:                  "error when aTokenBalance < unlockedDepositAmount",
			userWithdrawAmount:    big.NewInt(10),
			unlockedDepositAmount: big.NewInt(100),
			aTokenBalance:         big.NewInt(90),
			expectedErr:           fmt.Errorf("INTERNAL ERROR: aToken balance is less than unlocked deposit amount, aToken balance: %s, unlocked deposit amount: %s", big.NewInt(90), big.NewInt(100)),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			r := require.New(t)
			res, err := calculateWithdrawAmounts(tc.userWithdrawAmount, tc.unlockedDepositAmount, tc.aTokenBalance)
			if tc.expectedErr != nil {
				r.Error(err)
				r.EqualError(err, tc.expectedErr.Error())
				return
			}
			r.NoError(err)
			r.Equal(tc.expectedBaseAmount, res.BaseAmount, "base amount mismatch")
			r.Equal(tc.expectedAmountWithYield, res.AmountWithYield, "amount with yield mismatch")
		})
	}
}
