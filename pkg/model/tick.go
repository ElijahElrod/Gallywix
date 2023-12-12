package model

// Tick model for real-time market updates from the websocket exchange connection
type Tick struct {
	Timestamp int64
	Bid       float64
	Ask       float64
	Symbol    string
}
