package models

type Tick struct {
	Timestamp int64
	Bid       float64
	Ask       float64
	Symbol    string
}
