package transactor

import (
	"financial-agent-backend/core/abi/bindings/TrustManagementRouter"

	ethcommon "github.com/ethereum/go-ethereum/common"
	sgx_quote "github.com/Xyber-Labs/go-tee/sgx-quote"
)

type TeeService interface {
	GetQuote(userData []byte) ([]byte, error)
}

type TeeXyberProof struct {
	Leaf         TrustManagementRouter.ChunkedX509Cert
	Intermediate TrustManagementRouter.ChunkedX509Cert
	Quote        TrustManagementRouter.ChunkedSGXQuote
}

func FromSgxQuoteToProof(quote *sgx_quote.SgxQuote) (TeeXyberProof, error) {
	// The certificates are named differently in go-tee and Governor API.
	// Here's the mapping (go-tee -> Governor API):
	// * Device -> Leaf
	// * Platform -> Intermediary

	governorLeafCertificate := TrustManagementRouter.ChunkedX509Cert{
		BodyPartOne: ethcommon.FromHex(quote.Certificates.Device.X509Data.BodyPartOne),
		PublicKey:   ethcommon.FromHex(quote.Certificates.Device.X509Data.PublicKey),
		BodyPartTwo: ethcommon.FromHex(quote.Certificates.Device.X509Data.BodyPartTwo),
		Signature:   ethcommon.FromHex(quote.Certificates.Device.X509Data.Signature),
	}
	governorIntermediateCertificate := TrustManagementRouter.ChunkedX509Cert{
		BodyPartOne: ethcommon.FromHex(quote.Certificates.Platform.X509Data.BodyPartOne),
		PublicKey:   ethcommon.FromHex(quote.Certificates.Platform.X509Data.PublicKey),
		BodyPartTwo: ethcommon.FromHex(quote.Certificates.Platform.X509Data.BodyPartTwo),
		Signature:   ethcommon.FromHex(quote.Certificates.Platform.X509Data.Signature),
	}
	governorQuote := TrustManagementRouter.ChunkedSGXQuote{
		Header:               quote.Header,
		IsvReport:            quote.Report,
		IsvReportSignature:   quote.IsvReportSignature,
		AttestationKey:       quote.EcdsaAttestationKey,
		QeReport:             quote.QeReport,
		QeReportSignature:    quote.QeReportSignature,
		QeAuthenticationData: quote.QeAuthenticationData,
	}
	return TeeXyberProof{
		Leaf:         governorLeafCertificate,
		Intermediate: governorIntermediateCertificate,
		Quote:        governorQuote,
	}, nil
}
