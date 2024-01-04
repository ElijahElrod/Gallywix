package signal

import (
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
	Trend      Trend // Overall direction
	WindowSize int   // number of Ticks to track for trend and averages
	Average    float64
	Ticks      []float64 // Slice of windowSize model.Tick objects
}

// Guarantee Interface Match at Compile Time
var _ Signal = (*MovingAverage)(nil)

func newMovingAverage(windowSize int) *MovingAverage {
	return &MovingAverage{
		Trend:      Flat,
		WindowSize: windowSize,
		Ticks:      make([]float64, 0, windowSize),
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
	return len(ma.Ticks) == int(ma.WindowSize)
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
	ma.Average = runSum / float64(ma.WindowSize)
}

func (ma *MovingAverage) Name() string {
	//TODO implement me
	panic("Not implemented")
}

func (ma *MovingAverage) Details() string {
	//TODO implement me
	panic("Not implemented")
}

func (ma *MovingAverage) Evaluate(tick model.Tick) Side {
	//TODO implement me
	panic("Not implemented")
}
