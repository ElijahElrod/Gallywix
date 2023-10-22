package indicators

type ChannelType int8

const (
	Donchian ChannelType = iota
	Bollinger
	Keltner
)

type Channel interface {
	Update(tick float64)
	ChannelActive() bool
}

type DonchianChannel struct {
	upper   float64
	mid     float64
	lower   float64
	period  uint8
	channel ChannelType
}

func NewDonchianChannel(period uint8) *DonchianChannel {
	return &DonchianChannel{
		channel: Donchian,
		period:  period,
	}
}

func (dc *DonchianChannel) Update(tick float64) {
	if tick < dc.lower {
		dc.lower = tick
		dc._internalUpdate()
	}

	if tick > dc.upper {
		dc.upper = tick
		dc._internalUpdate()
	}
}

func (dc *DonchianChannel) _internalUpdate() {
	dc.mid = (dc.upper + dc.lower) / 2
}

type BollingerChannel struct {
	upper   float64
	lower   float64
	sma     MovingAverage
	channel ChannelType
}

func NewBollingerChannel() *BollingerChannel {
	return &BollingerChannel{
		channel: Bollinger,
	}
}

type KeltnerChannel struct {
}

func (kc *KeltnerChannel) PreviousTrueRange() bool {
	return false
}

func (kc *KeltnerChannel) CalculateATR() {
	if kc.PreviousTrueRange() {

		return
	}
}
