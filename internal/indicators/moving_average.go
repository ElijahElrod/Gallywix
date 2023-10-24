package indicators

import (
	"vespene.elijahelrod.com/internal/helpers"
)

type SimpleMovingAverage struct {
	trend       helpers.Trend
	windowSize  float64
	average     float64
	closingData []float64
}

func (sma *SimpleMovingAverage) Update(tick float64) {
	if sma.closingData = append(sma.closingData, tick); len(sma.closingData) > int(sma.windowSize) {
		sma.closingData = sma.closingData[1:]
	}

	sum := 0.0
	for _, v := range sma.closingData {
		sum += v
	}
	sma.average = sum / sma.windowSize
}
