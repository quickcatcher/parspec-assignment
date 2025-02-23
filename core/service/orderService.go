package service

import (
	"fmt"
	"parspec-assignment/core/domain"
	persistenceDriver "parspec-assignment/core/persistence/driver"

	"github.com/google/uuid"
)

type OrderService struct {
	OrderPersistence persistenceDriver.OrderPersistence
}

func (a *OrderService) CreateOrder(req *CreateOrderRequest) (resp *Response, err error) {
	orderId := uuid.New().String()
	order := domain.Orders{
		OrderId:     orderId, // Generate UUID,
		UserId:      req.UserId,
		ItemIDs:     req.ItemIDs,
		TotalAmount: req.TotalAmount,
		Status:      "Pending",
	}

	err = a.OrderPersistence.AddOrder(&order)
	if err != nil {
		fmt.Println("Error while creating order: ", err)
		return
	}

	resp = &Response{
		Message: "Order created successfully",
		Model: CreateOrderResponse{
			OrderId: orderId,
		},
	}
	return
}
