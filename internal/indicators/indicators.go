package indicators

import "vespene.elijahelrod.com/internal/helpers"

// Indicators is an interface used for to gather if an idicator is currently active and the signal it returns (Buy/Sell/No-op)
type Indicators interface {
	IndicatorActive() bool
	Signal(float64) helpers.Signal
}
