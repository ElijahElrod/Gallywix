package signal

import (
	"github.com/elijahelrod/vespene/pkg/model"
)

type Bollinger struct {
	upper   float64
	lower   float64
	sma     MovingAverage
	channel Type
}

func NewBollinger() *Bollinger {
	return &Bollinger{
		channel: BollingerChannel,
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
