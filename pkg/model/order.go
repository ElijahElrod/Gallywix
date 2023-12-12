package model

// OrderBody model used for creating the POST request body when placing Orders
type OrderBody struct {
	ProductId string `json:"product_id"`
	Side      string `json:"side"`
	Size      string `json:"size"`
	Price     string `json:"price"`
}
