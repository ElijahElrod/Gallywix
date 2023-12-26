// Package signal holds objects, and types for processing tick data and generating Buy/Sell/None signals off signal aggregation
package signal

import "github.com/elijahelrod/vespene/pkg/model"

type Type int8

// Trend represents whether a signal shows an up or down trend i.e. price is steadily increasing or decreasing
type Trend int8

// Side is used for evaluating trade signals i.e whether to Buy, Sell, or Do Nothing
type Side string

// Signal is an interface that represents trading signals which have a name, can update with new tick data, and evaluate new ticks using old tick data to generate trades and Sell/Buy signals
type Signal interface {
	Name() string
	Update(tick model.Tick)
	SignalActive() bool
	Details() string
	Evaluate(tick model.Tick) Side
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
