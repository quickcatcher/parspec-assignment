package driver

import (
	"parspec-assignment/core/domain"
	persistenceDriver "parspec-assignment/core/persistence/driver"
	svc "parspec-assignment/core/service"
)

// Service interface for order management. It can support multiple implementation which can be useful while launching new features.
type OrderSVC interface {
	CreateOrder(req *svc.CreateOrderRequest, orderQueue chan *domain.Orders, metrics *domain.Metrics) (resp *svc.Response, err error)
	GetOrderStatus(orderId int) (resp *svc.Response, err error)
}

func NewOrderService() OrderSVC {
	return &svc.OrderService{
		OrderPersistence: persistenceDriver.NewOrderPersistence(),
	}
}
