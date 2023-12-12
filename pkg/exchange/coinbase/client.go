package coinbase

import (
	"fmt"

	ws "github.com/gorilla/websocket"
)

const (
	ErrRequireConfigParameters = "Missing config parameters ( url )"
)

var wsDialer ws.Dialer

type Subscribe struct {
	Type       string   `json:"type"`
	ProductIds []string `json:"product_ids"`
	Channels   []string `json:"channels"`
}

type client struct {
	conn *ws.Conn
}

func NewClient(url string) (*client, error) {
	if url == "" {
		return nil, fmt.Errorf("%s", ErrRequireConfigParameters)
	}

	conn, _, err := wsDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}

	return &client{conn}, nil
}

func (c *client) WriteMsg(msg []byte) error {
	return c.conn.WriteMessage(2, msg)
}

func (c *client) ReadMsg() ([]byte, error) {
	_, msg, err := c.conn.ReadMessage()
	return msg, err
}

func (c *client) CloseConnection() error {
	return c.conn.Close()
}
