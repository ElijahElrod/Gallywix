package model

// OrderBody model used for creating the POST request body when placing Orders
type OrderBody struct {
	ProductId string `json:"product_id"`
	Side      string `json:"side"`
	Size      string `json:"size"`
	Price     string `json:"price"`
}

type OrderPlaceResponse struct {
	Success       bool
	FailureReason string
	OrderId       string
}

type OrderStatusResponse struct {
	OrderId      string `json:"order_id"`
	ProductId    string `json:"product_id"`
	UserId       string `json:"user_id"`
	Status       string `json:"status"`
	FillSize     string `json:"filled_size"`
	AvgFillPrice string `json:"average_filled_price"`
}

type OrderCancelResponse struct {
	Success       bool   `json:"success"`
	FailureReason string `json:"failure_reason"`
	OrderId       string `json:"order_id"`
}

type OrderTableRow struct {
	Timestamp int64
	OrderId   string
	ProductId string
	Side      string
	Size      string
	Price     string
	Status    string
}
