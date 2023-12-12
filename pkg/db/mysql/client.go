package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type client struct {
	*sqlx.DB
}

func NewClient(host, username, password, base string) (*client, error) {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@(%s)/%s", username, password, host, base))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &client{db}, err
}

func (db *client) Ping() error {
	return db.Ping()
}

func (db *client) Close() error {
	return db.Close()
}
