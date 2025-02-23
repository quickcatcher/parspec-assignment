package driver

import (
	"parspec-assignment/core/domain"
	"parspec-assignment/core/persistence"
)

type OrderPersistence interface {
	AddOrder(order *domain.Orders) (orderid int, err error)
	GetOrderbyOrderId(orderId int) (order *domain.Orders, err error)
	UpdateOrderStatus(status string, processing_time float64, orderId int) (err error)
}

func NewOrderPersistence() OrderPersistence {
	return &persistence.OrderModel{}
}
