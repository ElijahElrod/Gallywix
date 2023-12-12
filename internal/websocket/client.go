package websocket

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/elijahelrod/vespene/config"
	"github.com/elijahelrod/vespene/pkg/exchange"
	"github.com/elijahelrod/vespene/pkg/exchange/coinbase"
	"github.com/elijahelrod/vespene/pkg/logger"
	"github.com/elijahelrod/vespene/pkg/model"
)

type client struct {
	logger   logger.Logger
	conn     exchange.Manager
	products []string
	channels []string
}

func NewClient(conn exchange.Manager, logger logger.Logger, exchangeCfg config.ExchangeConfig) (*client, error) {
	if len(exchangeCfg.Symbols) == 0 {
		return nil, errors.New("no symbols available for tracking")
	}

	return &client{
		logger:   logger,
		conn:     conn,
		products: exchangeCfg.Symbols,
		channels: exchangeCfg.Channels,
	}, nil
}

func (c *client) Run(ctx context.Context) error {
	var errGroup = errgroup.Group{}
	var tickMap = make(map[string]chan model.Tick, len(c.products))

	for _, symbol := range c.products {
		tickMap[symbol] = make(chan model.Tick)
		errGroup.Go(func() error {
			for {
				select {
				case <-ctx.Done():
					return nil
				case v, ok := <-tickMap[symbol]:
					if ok {
						c.logger.Info(fmt.Sprintf("Ticker Update: %s > time:%d, bid:%f, ask:%f", v.Symbol, v.Timestamp, v.Bid, v.Ask))
					} else {
						return nil
					}
				}
			}
		})
	}

	subscribeMsg, _ := json.Marshal(map[string]interface{}{
		"type":        "subscribe",
		"product_ids": c.products,
		"channels":    c.channels,
	})
	err := c.conn.WriteMsg(subscribeMsg)
	if err != nil {
		return err
	}

	message, err := c.conn.ReadMsg()
	if err != nil {
		c.logger.Error(err)
		return err
	}
	result, err := coinbase.ParseResponse(message)
	if err != nil {
		c.logger.Error(err)
		return err
	}
	switch result.Type {
	case coinbase.Error:
		c.logger.Fatal(fmt.Sprintf("%s:%s", result.Message, result.Reason))
	case coinbase.Subscriptions:
		c.logger.Info(fmt.Sprintf("started subscription on products [%s]", strings.Join(c.products, ",")))
	}

	// writers
	for _, symbol := range c.products {
		errGroup.Go(func() error {
			return c.responseReader(symbol, tickMap)
		})
	}

	if err = errGroup.Wait(); err != nil {
		return err
	}

	return nil
}

// responseReader write to symbol channel from response socket data
func (c *client) responseReader(symbol string, tickMap map[string]chan model.Tick) error {

	var mu sync.Mutex
	var tickData *coinbase.Response

	for {
		message, err := c.conn.ReadMsg()
		if err != nil {
			c.logger.Error(err)
			if errors.Is(err, net.ErrClosed) {
				break
			}
			continue
		}

		tickData, err = coinbase.ParseResponse(message)
		if err != nil {
			c.logger.Error(err)
			continue
		}

		switch tickData.Type {
		case coinbase.Error:
			//TODO [critical] need break exchange and show global error?
			c.logger.Error(err)
			continue
		case coinbase.Ticker:
			mu.Lock()
			tickMap[symbol] <- model.Tick{
				Timestamp: time.Now().UnixNano(), // for exclude collision and accuracy time of ticker
				Bid:       tickData.BestBid,
				Ask:       tickData.BestAsk,
				Symbol:    tickData.ProductID,
			}
			mu.Unlock()
		}
	}

	return nil
}
