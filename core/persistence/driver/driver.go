package driver

import (
	"parspec-assignment/core/domain"
	"parspec-assignment/core/persistence"
)

type OrderPersistence interface {
	AddOrder(order *domain.Orders) (err error)
}

func NewOrderPersistence() OrderPersistence {
	return &persistence.OrderModel{}
}
