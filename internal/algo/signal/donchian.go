package signal

import (
	"fmt"
	"slices"

	"github.com/elijahelrod/vespene/pkg/model"
)

type Donchian struct {
	upper      float64
	mid        float64
	lower      float64
	lowPeriod  int
	highPeriod int
	channel    Type
	highs      []float64
	lows       []float64
}

var _ Signal = (*Donchian)(nil)

func NewDonchian(highPeriod, lowPeriod int) *Donchian {
	return &Donchian{
		channel: DonchianChannel,

		highPeriod: highPeriod,
		highs:      make([]float64, 0, highPeriod),

		lowPeriod: lowPeriod,
		lows:      make([]float64, 0, lowPeriod),
	}
}

func (d *Donchian) Name() string {
	return "Donchian"
}

func (d *Donchian) Update(tick model.Tick) {

	if len(d.highs) == d.highPeriod {
		d.highs = append(d.highs[1:], tick.Ask)
	} else {
		d.highs = append(d.highs, tick.Ask)
	}

	if len(d.lows) == d.lowPeriod {
		d.lows = append(d.lows[1:], tick.Bid)
	} else {
		d.lows = append(d.lows, tick.Bid)
	}

	d.lower = slices.Min(d.lows)
	d.upper = slices.Max(d.highs)
	d.updateMid()

}

func (d *Donchian) Evaluate(tick model.Tick) Side {

	if d.SignalActive() {
		if tick.Price > d.upper*0.95 {
			return BUY
		}

		if tick.Price < d.lower {
			return SELL
		}
	}

	return NONE
}

func (d *Donchian) UpdateAndEvaluate(tick model.Tick) Side {
	d.Update(tick)
	return d.Evaluate(tick)
}

func (d *Donchian) SignalActive() bool {
	return len(d.highs) == d.highPeriod &&
		len(d.lows) == d.lowPeriod
}

func (d *Donchian) updateMid() {
	d.mid = (d.upper + d.lower) / 2
}

func (d *Donchian) Details() string {
	return fmt.Sprintf("Signal: %s :: Upper: %f, Mid: %f, Lower: %f", d.Name(), d.upper, d.mid, d.lower)
}
