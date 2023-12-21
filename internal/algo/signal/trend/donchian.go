package trend

import (
	"fmt"
	"slices"

	"github.com/elijahelrod/vespene/internal/algo/signal"
	"github.com/elijahelrod/vespene/pkg/model"
)

type Donchian struct {
	upper      float64
	mid        float64
	lower      float64
	lowPeriod  int
	highPeriod int
	channel    signal.Type
	highs      []float64
	lows       []float64
}

func NewDonchian(highPeriod, lowPeriod int) *Donchian {
	return &Donchian{
		channel: signal.DonchianChannel,

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
		d.highs = append(d.highs[1:], tick.DailyHigh)
	} else {
		d.highs = append(d.highs, tick.DailyHigh)
	}

	if len(d.lows) == d.lowPeriod {
		d.lows = append(d.lows[1:], tick.DailyLow)
	} else {
		d.lows = append(d.lows, tick.DailyLow)
	}

	d.lower = slices.Min(d.lows)
	d.upper = slices.Max(d.highs)
	d.updateMid()

}

func (d *Donchian) Evaluate(tick model.Tick) signal.Side {

	if tick.Price > d.upper {
		return signal.BUY
	}

	if tick.Price < d.lower {
		return signal.SELL
	}

	return signal.NONE
}

func (d *Donchian) UpdateAndEvaluate(tick model.Tick) signal.Side {
	d.Update(tick)
	if d.SignalActive() {
		return d.Evaluate(tick)
	}
	return signal.NONE
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
