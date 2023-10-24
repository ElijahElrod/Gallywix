package helpers

type Signal int

const (
	Buy Signal = iota
	Sell
	No_op
)
