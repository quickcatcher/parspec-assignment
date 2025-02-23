package driver

import (
	persistenceDriver "parspec-assignment/core/persistence/driver"
	svc "parspec-assignment/core/service"
)

type OrderSVC interface {
	CreateOrder(req *svc.CreateOrderRequest) (resp *svc.Response, err error)
}

func NewOrderService() OrderSVC {
	return &svc.OrderService{
		OrderPersistence: persistenceDriver.NewOrderPersistence(),
	}
}
