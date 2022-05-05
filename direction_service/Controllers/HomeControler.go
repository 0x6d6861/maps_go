package Controllers

import (
	"direction_service/Controllers/Base"
	"direction_service/Src/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeController(router gin.IRouter) *Base.BaseController {
	controller := Base.NewBaseController(router)

	controller.GET("/full", func(c *gin.Context) {
		var u Services.DirectionQueryRequest
		err := c.BindQuery(&u)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		directionQuery := Services.DirectionQuery{
			Points: u.Point,
		}

		response, err := Services.OSMInstance.GetDirection(directionQuery)

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
