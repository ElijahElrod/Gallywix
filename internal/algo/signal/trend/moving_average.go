package trend

import (
	"github.com/elijahelrod/vespene/internal/algo/signal"
	"github.com/elijahelrod/vespene/pkg/model"
)

// Short represents a 7-day window
// Medium represents a 14-day window
// Long represents a 21-day window
const (
	Short int = (iota + 1) * 7
	Medium
	Long
)

type MovingAverage struct {
	trend      signal.Trend // Overall direction
	windowSize int          // number of Ticks to track for trend and averages
	average    float64
	Ticks      []model.Tick // Slice of windowSize model.Tick objects
}

func newMovingAverage(windowSize int) *MovingAverage {
	return &MovingAverage{
		trend:      signal.Flat,
		windowSize: windowSize,
		Ticks:      make([]model.Tick, 0, windowSize),
	}
}

func NewLongMovingAverage() *MovingAverage {
	return newMovingAverage(Long)
}

func NewMediumMovingAverage() *MovingAverage {
	return newMovingAverage(Medium)
}

func NewShortMovingAverage() *MovingAverage {
	return newMovingAverage(Short)
}

func (ma *MovingAverage) SignalActive() bool {
	return len(ma.Ticks) == ma.windowSize
}

func (ma *MovingAverage) Update(tick model.Tick) {

	// Take the most recent [ma.windowSize] data points if at capacity, normal add otherwise
	if len(ma.Ticks) == ma.windowSize {
		ma.Ticks = append(ma.Ticks[1:], tick)
	} else {
		ma.Ticks = append(ma.Ticks, tick)
	}
}
