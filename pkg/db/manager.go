package db

type Manager interface {
	Ping() error
	Close()
}
