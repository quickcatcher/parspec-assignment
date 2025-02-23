package service

import (
	"fmt"
	"math/rand"
	"parspec-assignment/core/domain"
	persistenceDriver "parspec-assignment/core/persistence/driver"
	"time"
)

type OrderService struct {
	OrderPersistence persistenceDriver.OrderPersistence
}

func (a *OrderService) CreateOrder(req *CreateOrderRequest, orderQueue chan *domain.Orders, metrics *domain.Metrics) (resp *Response, err error) {
	order := &domain.Orders{
		UserId:      req.UserId,
		ItemIds:     req.ItemIDs,
		TotalAmount: req.TotalAmount,
		Status:      "Pending",
	}
	metrics.MutexLock()
	metrics.OrderStatusCounts["Pending"]++
	metrics.MutexUnLock()

	orderId, err := a.OrderPersistence.AddOrder(order)
	if err != nil {
		fmt.Println("Error while creating order: ", err)
		return
	}

	orderQueue <- order
	metrics.MutexLock()
	metrics.OrderStatusCounts["Pending"]--
	metrics.OrderStatusCounts["Processing"]++
	metrics.MutexUnLock()

	resp = &Response{
		Code:    200,
		Message: "Order created successfully",
		Model: &CreateOrderResponse{
			OrderId: orderId,
		},
	}
	return
}

func (a *OrderService) GetOrderStatus(orderId int) (resp *Response, err error) {
	order, err := a.OrderPersistence.GetOrderbyOrderId(orderId)
	if err != nil {
		fmt.Println("Error while fetching order: ", err)
	}
	resp = &Response{}
	if order == nil {
		resp.Code = 404
		resp.Message = "Order not found"
		return
	}

	resp.Code = 200
	resp.Message = "Order Found"
	resp.Model = &GetOrderStatusResponse{
		Status:      order.Status,
		ItemIds:     order.ItemIds,
		TotalAmount: order.TotalAmount,
	}
	return
}

func ProcessQueueOrders(orderQueue chan *domain.Orders, metrics *domain.Metrics) {
	var err error
	for {
		order := <-orderQueue
		fmt.Println(order)

		processingTime := rand.Intn(10) // assuming that the processing takes anything between 0-10 seconds
		time.Sleep(time.Duration(processingTime) * time.Second)

		orderPersistence := persistenceDriver.NewOrderPersistence()
		err = orderPersistence.UpdateOrderStatus("Completed", float64(processingTime), order.OrderId)
		if err != nil {
			orderQueue <- order
		}
		metrics.MutexLock()
		metrics.TotalOrdersProcessed++
		metrics.AverageProcessingTime = (metrics.AverageProcessingTime*float64(metrics.TotalOrdersProcessed-1) + float64(processingTime)) / float64(metrics.TotalOrdersProcessed)
		metrics.OrderStatusCounts["Processing"]--
		metrics.OrderStatusCounts["Completed"]++
		metrics.MutexUnLock()
	}
}
