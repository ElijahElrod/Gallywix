package trend

import (
	"github.com/elijahelrod/vespene/internal/algo/signal"
	"github.com/elijahelrod/vespene/pkg/model"
)

// Short represents a 7-day window
// Medium represents a 14-day window
// Long represents a 21-day window
const (
	Short float64 = (iota + 1) * 7
	Medium
	Long
)

type MovingAverage struct {
	trend      signal.Trend // Overall direction
	windowSize float64      // number of Ticks to track for trend and averages
	average    float64
	Ticks      []float64 // Slice of windowSize model.Tick objects
}

func newMovingAverage(windowSize float64) *MovingAverage {
	return &MovingAverage{
		trend:      signal.Flat,
		windowSize: windowSize,
		Ticks:      make([]float64, 0, int(windowSize)),
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
	return len(ma.Ticks) == int(ma.windowSize)
}

func (ma *MovingAverage) Update(tick model.Tick) {

	// Take the most recent [ma.windowSize] data points if at capacity, normal add otherwise
	if ma.SignalActive() {
		ma.Ticks = append(ma.Ticks[1:], tick.Price)
	} else {
		ma.Ticks = append(ma.Ticks, tick.Price)
	}

	var runSum = 0.0
	for _, val := range ma.Ticks {
		runSum += val
	}
	ma.average = runSum / ma.windowSize
}
