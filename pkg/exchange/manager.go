package exchange

type Manager interface {
	CloseConnection() error
	WriteMsg(message []byte) error
	ReadMsg() ([]byte, error)
}
