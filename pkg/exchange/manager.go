package exchange

type Manager interface {
	Close() error
	Write(message []byte) (int, error)
	Read() ([]byte, error)
}
