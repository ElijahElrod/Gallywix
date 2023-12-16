package channel

import (
	"fmt"

	"github.com/elijahelrod/vespene/internal/algo/signal"
	"github.com/elijahelrod/vespene/pkg/model"
)

type Donchian struct {
	upper   float64
	mid     float64
	lower   float64
	period  int
	channel signal.Type
	ticks   []model.Tick
}

func NewDonchian(period int) *Donchian {
	return &Donchian{
		channel: signal.DonchianChannel,
		period:  period,
		ticks:   make([]model.Tick, 0, period),
	}
}

func (d *Donchian) Name() string {
	return "Donchian"
}

func (d *Donchian) Update(tick model.Tick) {

	if len(d.ticks) == d.period {
		d.ticks = append(d.ticks[1:], tick)
	} else {
		d.ticks = append(d.ticks, tick)
	}

	if len(d.ticks) == 1 {
		d.lower = tick.Bid
		d.upper = tick.Ask
	}

	if tick.Bid < d.lower {
		d.lower = tick.Bid
		d.updateMid()
	}

	if tick.Bid > d.upper {
		d.upper = tick.Bid
		d.updateMid()
	}
}

func (d *Donchian) SignalActive() bool {
	return len(d.ticks) == d.period
}

func (d *Donchian) updateMid() {
	d.mid = (d.upper + d.lower) / 2
}

func (d *Donchian) Details() string {
	return fmt.Sprintf("Signal: %s :: Upper: %f, Mid: %f, Lower: %f", d.Name(), d.upper, d.mid, d.lower)
}
