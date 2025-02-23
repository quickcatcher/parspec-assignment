package service

type CreateOrderRequest struct {
	UserId      int     `json:"user_id" binding:"required"`
	ItemIDs     string  `json:"item_ids" binding:"required"`
	TotalAmount float64 `json:"total_amount" binding:"required"`
}

type Response struct {
	Code    int         `json:"-"` // This will be ignored during JSON marshaling
	Message string      `json:"message"`
	Model   interface{} `json:"model"`
}

type CreateOrderResponse struct {
	OrderId int `json:"order_id"`
}

type GetOrderStatusResponse struct {
	Status      string  `json:"status"`
	ItemIds     string  `json:"item_ids"`
	TotalAmount float64 `json:"total_amount"`
}

type GetMetricsResponse struct {
	TotalOrdersProcessed int
}
