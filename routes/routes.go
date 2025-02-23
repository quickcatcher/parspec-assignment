package routes

import (
	"parspec-assignment/core/domain"
	svc "parspec-assignment/core/service"
	svcDriver "parspec-assignment/core/service/driver"
	"parspec-assignment/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func EngineRoutes(router *gin.RouterGroup, orderQueue chan *domain.Orders, metrics *domain.Metrics) {

	go svc.ProcessQueueOrders(orderQueue, metrics)
	router.Use(middleware.DBConnection)

	router.POST("/order", func(c *gin.Context) {
		CreateOrder(c, orderQueue, metrics)
	})

	router.GET("/order/:id", func(c *gin.Context) {
		GetOrderStatus(c)
	})

	router.GET("/metrics", func(c *gin.Context) {
		GetMetrics(c, metrics)
	})
}

func CreateOrder(c *gin.Context, orderQueue chan *domain.Orders, metrics *domain.Metrics) {
	request := &svc.CreateOrderRequest{}

	err := c.ShouldBindJSON(request)
	if err != nil {
		response := &svc.Response{Message: "Bad Request"}
		c.JSON(400, response)
		return
	}

	orderSvc := svcDriver.NewOrderService()

	resp, err := orderSvc.CreateOrder(request, orderQueue, metrics)
	if err != nil {
		response := &svc.Response{Message: "Something went wrong"}
		c.JSON(500, response)
		return
	}
	c.JSON(resp.Code, resp)
}

func GetOrderStatus(c *gin.Context) {
	orderId := c.Param("id")
	if orderId == "" {
		response := &svc.Response{Message: "OrderId is mandatory"}
		c.JSON(400, response)
		return
	}

	orderSvc := svcDriver.NewOrderService()

	resp, err := orderSvc.GetOrderStatus(cast.ToInt(orderId))
	if err != nil {
		response := &svc.Response{Message: "Something went wrong"}
		c.JSON(500, response)
		return
	}
	c.JSON(resp.Code, resp)
}

func GetMetrics(c *gin.Context, metrics *domain.Metrics) {
	metrics.MutexLock()
	defer metrics.MutexUnLock()
	c.JSON(200, metrics)
}
