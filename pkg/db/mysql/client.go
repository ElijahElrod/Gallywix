// Package mysql implements utilities for interacting with a MySQL database configured in [DatabaseConfig]
package mysql

import (
	"database/sql"
	"fmt"

	"github.com/elijahelrod/vespene/pkg/logger"
	"github.com/elijahelrod/vespene/pkg/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type client struct {
	*sqlx.DB
	logger logger.Logger
}

// NewClient creates a new MySQL Client for interacting with the databased configured by
// [config.DatabaseConfig]
func NewClient(host, username, password, base string, logger logger.Logger) (*client, error) {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@(%s)/%s", username, password, host, base))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &client{db, logger}, err
}

func (db *client) PingDB() error {
	return db.Ping()
}

func (db *client) CloseConnection() error {
	return db.Close()
}

func (db *client) QueryByOrderId(orderId string) ([]model.OrderTableRow, error) {
	dbRows, err := db.Query("SELECT * FROM tbl_orders WHERE orderId = ? LIMIT 1", orderId)
	if err != nil {
		return []model.OrderTableRow{}, err
	}
	defer func(dbRows *sql.Rows) {
		err := dbRows.Close()
		if err != nil {
			db.logger.Error(err)
		}
	}(dbRows)

	orders := make([]model.OrderTableRow, 1)
	for dbRows.Next() {
		var order model.OrderTableRow
		if err := dbRows.Scan(&order); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil

}
