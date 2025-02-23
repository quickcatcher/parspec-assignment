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

	// started go routine for asynchronous processing of queue
	go svc.ProcessQueueOrders(orderQueue, metrics)
	router.Use(middleware.DBConnection)

	// POST order route for adding orders
	router.POST("/order", func(c *gin.Context) {
		CreateOrder(c, orderQueue, metrics)
	})

	// GEt order route to get status and order details
	router.GET("/order/:id", func(c *gin.Context) {
		GetOrderStatus(c)
	})

	//  GET metrics route for getting metrics data
	router.GET("/metrics", func(c *gin.Context) {
		GetMetrics(c, metrics)
	})
}

func CreateOrder(c *gin.Context, orderQueue chan *domain.Orders, metrics *domain.Metrics) {
	request := &svc.CreateOrderRequest{}

	// validating request for adding order
	err := c.ShouldBindJSON(request)
	if err != nil {
		response := &svc.Response{Message: "Bad Request"}
		c.JSON(400, response)
		return
	}

	// defining order service
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
	//extracting order id from url parameter
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
