package service

type CreateOrderRequest struct {
	UserId      int     `json:"user_id" binding:"required"`
	ItemIDs     string  `json:"item_ids" binding:"required"`
	TotalAmount float64 `json:"total_amount" binding:"required"`
}

type Response struct {
	Message string      `json:"message"`
	Model   interface{} `json:"model"`
}

type CreateOrderResponse struct {
	OrderId string `json:"order_id"`
}
