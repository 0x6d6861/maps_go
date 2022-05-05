package Controllers

import (
	"github.com/gin-gonic/gin"
	"mpasGo/Controllers/Base"
)

func HomeController(router gin.IRouter) *Base.BaseController {
	controller := Base.NewBaseController(router)

	controller.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//controller.PrintRoutes()

	return controller
}
