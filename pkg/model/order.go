package model

type OrderBody struct {
	ProductId string `json:"product_id"`
	Side      string `json:"side"`
	Size      string `json:"size"`
	Price     string `json:"price"`
}
