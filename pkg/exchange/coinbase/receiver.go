package coinbase

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Type      ResponseType `json:"type,string"`
	Message   string       `json:"message,omitempty"`
	Reason    string       `json:"reason,omitempty"`
	ProductID string       `json:"product_id"`
	BestBid   float64      `json:"best_bid,string"`
	BestAsk   float64      `json:"best_ask,string"`
	Price     float64      `json:"price,string"`
	DailyLow  float64      `json:"low_24h,string"`
	DailyHigh float64      `json:"high_24h,string"`
	DailyVol  float64      `json:"volume_24h,string"`
	Sequence  int64        `json:"sequence"`
}

type ResponseType int

const (
	Error ResponseType = iota
	Subscriptions
	Unsubscribe
	Heartbeat
	Ticker
	Level2
)

var responseTypes = [...]string{"error", "subscriptions", "unsubscribe", "heartbeat", "ticker", "level2"}

func (r ResponseType) String() string {
	return responseTypes[r]
}

func (r *ResponseType) UnmarshalJSON(v []byte) error {
	str := string(v)

	for i, name := range responseTypes {
		if name == str {
			*r = ResponseType(i)
			return nil
		}
	}

	return fmt.Errorf("invalid locality type %q", str)
}

func ParseResponse(message []byte) (response *Response, err error) {

	err = json.Unmarshal(message, &response)
	if err != nil {
		return nil, err
	}

	return
}
