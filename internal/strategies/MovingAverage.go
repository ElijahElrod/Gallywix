package strategies

type Trend int8

const (
	Up Trend = iota
	Down
	Flat
)

type WindowSize int8

const (
	Short WindowSize = iota
	Long
)

type MovingAverage struct {
	trend      Trend
	windowSize WindowSize
}
