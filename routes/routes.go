package routes

import (
	svc "parspec-assignment/core/service"
	svcDriver "parspec-assignment/core/service/driver"
	"parspec-assignment/middleware"

	"github.com/gin-gonic/gin"
)

func EngineRoutes(router *gin.RouterGroup) {
	router.Use(middleware.DBConnection)
	router.POST("/order", CreateOrder)
}

func CreateOrder(c *gin.Context) {
	request := &svc.CreateOrderRequest{}

	err := c.ShouldBindJSON(request)
	if err != nil {
		response := &svc.Response{Message: "Bad Request"}
		c.JSON(400, response)
		return
	}

	orderSvc := svcDriver.NewOrderService()

	resp, err := orderSvc.CreateOrder(request)
	if err != nil {
		response := &svc.Response{Message: "Something went wrong"}
		c.JSON(500, response)
		return
	}
	c.JSON(200, resp)
}
