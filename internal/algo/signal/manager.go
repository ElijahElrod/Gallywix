package signal

import "github.com/elijahelrod/vespene/pkg/model"

type Type int8
type Trend int8
type Side string

type Signal interface {
	Name() string
	Update(tick model.Tick)
	SignalActive() bool
	Details() string
}

const (
	DonchianChannel Type = iota + 1
	BollingerChannel
)

const (
	Flat Trend = iota
	Up
	Down
)

const (
	NONE Side = "None"
	BUY  Side = "Buy"
	SELL Side = "Sell"
)
