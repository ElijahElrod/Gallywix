package strategies

type ChannelType int8

const (
	Donchian ChannelType = iota
	Bollinger
	Keltner
)

type Channel interface {
	Update(tick float64)
}

type DonchianChannel struct {
	upper  float64
	mid    float64
	lower  float64
	period uint8
}

func NewDonchianChannel() *DonchianChannel {
	return &DonchianChannel{}
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
	upper float64
	lower float64
	sma   MovingAverage
}
