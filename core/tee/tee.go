package tee

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"

	sgx_quote "github.com/Xyber-Labs/go-tee/sgx-quote"
	"github.com/ethereum/go-ethereum/crypto"

)

func ExtractQuote(teeService *sgx_quote.TeeService, userData []byte) ([]byte, error) {
	return teeService.GetQuote(userData)
}

