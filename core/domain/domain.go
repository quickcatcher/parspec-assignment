package domain

// order struct containing user order details
type Orders struct {
	Id             int     `json:"id"` // primary key of the table
	OrderId        string  `json:"order_id"`
	UserId         int     `json:"user_id"`
	ItemIDs        string  `json:"item_ids"`
	TotalAmount    float64 `json:"total_amount"`
	Status         string  `json:"status"`
	ProcessingTime float64 `json:"processing_time"`
}
