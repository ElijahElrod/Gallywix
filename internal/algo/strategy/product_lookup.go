package strategy

// ProductLookup is used to map product-ids to respective strategies for order placing + tracking
type ProductLookup struct {
	mapping map[string][]Strategy
}

func NewStrategyProductLookup() *ProductLookup {
	return &ProductLookup{mapping: make(map[string][]Strategy)}
}

/*func (pl *ProductLookup) isEmpty() bool {
	return len(pl.mapping) == 0
}

func (pl *ProductLookup) getStrategies(productId string) []Strategy {
	strategies, ok := pl.mapping[productId]
	if !ok {
		return make([]Strategy, 0)
	}
	return strategies
}
*/
