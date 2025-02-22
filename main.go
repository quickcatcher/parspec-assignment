package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	// default gin engine
	r := gin.Default()
	// grouped the base path for endoints
	public := r.Group("/parspec")

	public.POST("/order", func(c *gin.Context) {
		c.JSON(200, "Success")
	})

	r.Run(":9000")

}
