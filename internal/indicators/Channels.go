// Package indicators is used to hold trading indicators types and methods that will generate buy/sell signals for trading strategies
package indicators

import (
	"math"
	"vespene.elijahelrod.com/internal/helpers"
)

type ChannelType int8

const (
	Donchian ChannelType = iota
	Bollinger
	Keltner
)

type Channel interface {
	Update(tick float64)
}

// DonchianChannel is used to represent the Donchian type for creating trading indicators
type DonchianChannel struct {
	upper   float64
	mid     float64
	lower   float64
	period  int
	data    []float64
	channel ChannelType
}

func NewDonchianChannel(period int) *DonchianChannel {
	return &DonchianChannel{
		channel: Donchian,
		period:  period,
		data:    make([]float64, period),
	}
}

// Update adds the latest tick to the underlying data field, and recalculates the channel's highs / lows, and average
func (dc *DonchianChannel) update(tick float64) {

	// Add new data, check if we have more elements than specified by period, remove oldest
	if dc.data = append(dc.data, tick); len(dc.data) > dc.period {
		dc.data = dc.data[1:]
	}

	dc.updateChannelBounds()

}

// findMinMax checks a DonchianChannel's data to return the lowest low and highest high respectively in the current slice
func (dc *DonchianChannel) updateChannelBounds() {
	minTick, maxTick := math.MaxFloat64, 0.0
	for _, v := range dc.data {
		if v < minTick {
			minTick = v
		}

		if v > maxTick {
			maxTick = v
		}
	}

	dc.lower = minTick
	dc.upper = maxTick
	dc.mid = (dc.upper + dc.lower) / 2

}

// IndicatorActive checks if the underlying data is fully saturated with DonchianChannel's period number of value points
func (dc *DonchianChannel) IndicatorActive() bool {
	return len(dc.data) == dc.period
}

// Signal returns if the current indicator is favoring Buy or Sell Conditions
func (dc *DonchianChannel) Signal(tick float64) helpers.Signal {
	dc.update(tick)
	if dc.IndicatorActive() {
		// Buy when we're lower in the channel
		if tick < dc.mid {
			return helpers.Buy
		}

		if tick > dc.mid {
			return helpers.Sell
		}

	}

	return helpers.No_op
}

type BollingerChannel struct {
	upper   float64
	lower   float64
	sma     SimpleMovingAverage
	channel ChannelType
}

func NewBollingerChannel() *BollingerChannel {
	return &BollingerChannel{
		channel: Bollinger,
	}
}

func (bc *BollingerChannel) IndicatorActive() bool {
	return false
}

type KeltnerChannel struct {
}

func (kc *KeltnerChannel) PreviousTrueRange() bool {
	panic("Not yet implemented")
}

func (kc *KeltnerChannel) CalculateATR() {
	panic("Not yet implemented")
	// Calculate the first True Range
}
