package routes

import (
	"parspec-assignment/middleware"

	"github.com/gin-gonic/gin"
)

func EngineRoutes(router *gin.RouterGroup) {
	router.Use(middleware.DBConnection)
	router.POST("/order", func(c *gin.Context) {
		c.JSON(200, "Success")
	})
}
