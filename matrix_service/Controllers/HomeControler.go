package Controllers

import (
	"github.com/gin-gonic/gin"
	"matrix_service/Controllers/Base"
	"matrix_service/Src/Services"
	"net/http"
	"strings"
)

func HomeController(router gin.IRouter) *Base.BaseController {
	controller := Base.NewBaseController(router)

	controller.GET("/", func(c *gin.Context) {
		var u Services.MatrixQuery
		err := c.BindQuery(&u)

		u.Origins = strings.Split(c.Request.URL.Query()["origins"][0], "|")
		u.Destinations = strings.Split(c.Request.URL.Query()["destinations"][0], "|")

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		response, _ := Services.GraphHopperInstance.GetMatrix(u)
		c.JSON(200, response)
	})

	controller.GET("/single", func(c *gin.Context) {
		var u Services.MatrixQuery
		err := c.BindQuery(&u)

		u.Origins = []string{c.Request.URL.Query()["origin"][0]}
		u.Destinations = []string{c.Request.URL.Query()["destination"][0]}

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		response, _ := Services.GraphHopperInstance.GetMatrix(u)
		c.JSON(200, response)
	})

	//controller.PrintRoutes()

	return controller
}
