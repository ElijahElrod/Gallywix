package indicators

import (
	helpers "gallywix.elijahelrod.com/internal/helpers"
)

type WindowSize int8

const (
	Short WindowSize = iota
	Long
)

type MovingAverage struct {
	trend      helpers.Trend
	windowSize WindowSize
	average    float64
}
