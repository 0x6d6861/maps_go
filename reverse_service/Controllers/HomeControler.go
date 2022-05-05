package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reverse_service/Controllers/Base"
	"reverse_service/Src/Services"
)

func HomeController(router gin.IRouter) *Base.BaseController {
	controller := Base.NewBaseController(router)

	controller.GET("/", func(c *gin.Context) {
		var u Services.ReverseQuery
		err := c.BindQuery(&u)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		response, err := Services.OSMInstance.GetReverse(u)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		name := response.Features[0].Properties.Name

		if name == "" {
			name = response.Features[0].Properties.DisplayName
		}

		c.JSON(200, gin.H{
			"status": "000",
			"name":   name,
		})

	})

	controller.GET("/full", func(c *gin.Context) {
		var u Services.ReverseQuery
		err := c.BindQuery(&u)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		response, err := Services.OSMInstance.GetReverse(u)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"status":   "000",
			"response": response,
		})
	})

	//controller.PrintRoutes()

	return controller
}
