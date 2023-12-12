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
	*ws.Conn
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

func (c *client) Write(msg []byte) (int, error) {
	return c.Write(msg)
}

func (c *client) Read() ([]byte, error) {
	return c.Read()
}

func (c *client) Close() error {
	return c.Close()
}
