package db

type Manager interface {
	PingDB() error
	CloseConnection()
}
