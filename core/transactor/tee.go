package transactor

type TeeService interface {
	GetQuote(userData []byte) ([]byte, error)
}
