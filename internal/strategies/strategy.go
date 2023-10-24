// Package strategies contains the strategy struct used to create groupings of indicators for back-testing and live trading
package strategies

import (
	"fmt"
	"vespene.elijahelrod.com/internal/indicators"
)

// Strategy is ued to represent a trading strategy and holds a collection of indicators, as well as a flag for setting it to active or disabled
type Strategy struct {
	name       string                  // Name of the strategy
	Indicators []indicators.Indicators // Indicators used in the strategy for generating buy and sell signals
	Active     bool                    // Flag used to toggle if a strategy is currently active or not
}

func NewStrategy() *Strategy {
	return &Strategy{Active: false, Indicators: make([]indicators.Indicators, 10), name: "Test Strategy 1"}
}

func (s *Strategy) ToString() {
	fmt.Printf("Strategy %s has %d indicators [%v], Currently Active [%t]", s.name, len(s.Indicators), s.Indicators, s.Active)
}
