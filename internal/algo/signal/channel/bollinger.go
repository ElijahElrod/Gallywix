package channel

import (
	"github.com/elijahelrod/vespene/internal/algo/signal"
	"github.com/elijahelrod/vespene/internal/algo/signal/trend"
	"github.com/elijahelrod/vespene/pkg/model"
)

type Bollinger struct {
	upper   float64
	lower   float64
	sma     trend.MovingAverage
	channel signal.Type
}

func NewBollinger() *Bollinger {
	return &Bollinger{
		channel: signal.BollingerChannel,
	}
}

func (b *Bollinger) SignalActive() bool {
	return b.sma.SignalActive()
}

func (b *Bollinger) Update(tick model.Tick) {
	b.sma.Update(tick)
	b.upper = 0
	b.lower = 0
}
