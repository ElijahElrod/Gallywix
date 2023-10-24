package web

import (
	"flag"
	"vespene.elijahelrod.com/internal/strategies"
)

// application struct for holding global references
type application struct {
}

// main creates the application object, and starts web handlers & sockets used for management and trading
func main() {

	flag.Parse()
	// Get Active Strategies
	// Connect to exchanges
	//

	currentStrategy := strategies.NewStrategy()
	currentStrategy.ToString()

	for {
		tick := 0.0
		count := 0
		// Evaluate indicator signals for a given strategy
		for _, indicator := range currentStrategy.Indicators {

			if indicator.IndicatorActive() {
				if currSignal := indicator.Signal(tick); currSignal == 0 {
					count += 1
				} else {
					count -= 1
				}
			}
		}

		if count < 0 {
			// Place Sell Order
		} else if count > 0 {
			// Place Buy Order
		}
	}

}
