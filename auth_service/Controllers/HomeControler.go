package Controllers

import (
	"auth_service/Controllers/Base"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeController(router gin.IRouter) *Base.BaseController {
	controller := Base.NewBaseController(router)

	controller.GET("/key", func(c *gin.Context) {
		key := c.Request.URL.Query()["key"]
		fmt.Println(key)
		if len(key) != 1 || key[0] != "a97d6edd-ff2d-4ace-9ddd-9e784ab5bf5c" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		// log login
		c.JSON(200, gin.H{
			"message": "login",
		})
	})

	controller.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//controller.PrintRoutes()

	return controller
}
