package transactor

import (
	"financial-agent-backend/core/abi/bindings/TEEWallet"

	sgx_quote "github.com/Xyber-Labs/go-tee/sgx-quote"
)

type TeeService interface {
	GetQuote(userData []byte) ([]byte, error)
}

type TeeWalletProof struct {
	Leaf         TEEWallet.ChunkedX509Cert
	Intermediate TEEWallet.ChunkedX509Cert
	Quote        TEEWallet.ChunkedSGXQuote
}

func FromSgxQuoteToTeeWalletProof(quote *sgx_quote.SgxQuote) (TeeWalletProof, error) {
	// The certificates are named differently in go-tee and Governor API.
	// Here's the mapping (go-tee -> Governor API):
	// * Device -> Leaf
	// * Platform -> Intermediary

	governorLeafCertificate := TEEWallet.ChunkedX509Cert{
		BodyPartOne: []byte(quote.Certificates.Device.X509Data.BodyPartOne),
		PublicKey:   []byte(quote.Certificates.Device.X509Data.PublicKey),
		BodyPartTwo: []byte(quote.Certificates.Device.X509Data.BodyPartTwo),
		Signature:   []byte(quote.Certificates.Device.X509Data.Signature),
	}
	governorIntermediateCertificate := TEEWallet.ChunkedX509Cert{
		BodyPartOne: []byte(quote.Certificates.Platform.X509Data.BodyPartOne),
		PublicKey:   []byte(quote.Certificates.Platform.X509Data.PublicKey),
		BodyPartTwo: []byte(quote.Certificates.Platform.X509Data.BodyPartTwo),
		Signature:   []byte(quote.Certificates.Platform.X509Data.Signature),
	}
	governorQuote := TEEWallet.ChunkedSGXQuote{
		Header:               quote.Header,
		IsvReport:            quote.Report,
		IsvReportSignature:   quote.IsvReportSignature,
		AttestationKey:       quote.EcdsaAttestationKey,
		QeReport:             quote.QeReport,
		QeReportSignature:    quote.QeReportSignature,
		QeAuthenticationData: quote.QeAuthenticationData,
	}
	return TeeWalletProof{
		Leaf:         governorLeafCertificate,
		Intermediate: governorIntermediateCertificate,
		Quote:        governorQuote,
	}, nil
}
