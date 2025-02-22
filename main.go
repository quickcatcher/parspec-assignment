package main

import (
	"parspec-assignment/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// default gin engine
	r := gin.Default()
	// grouped the base path for endoints
	router := r.Group("/parspec")
	routes.EngineRoutes(router)

	r.Run(":9000")

}
