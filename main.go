package main

import (
	"parspec-assignment/core/domain"
	"parspec-assignment/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// creating in memory queue as golang channel and metrics which will be maintained through every API calls
	orderQueue := make(chan *domain.Orders)
	metrics := &domain.Metrics{
		OrderStatusCounts: map[string]int{
			"Pending":    0,
			"Processing": 0,
			"Completed":  0,
		},
	}

	// default gin engine
	r := gin.Default()
	// grouped the base path for endoints
	router := r.Group("/parspec")
	routes.EngineRoutes(router, orderQueue, metrics)

	r.Run(":9000")

}
