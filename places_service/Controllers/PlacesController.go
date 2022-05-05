package Controllers

import (
	"github.com/gin-gonic/gin"
	"mpasGo/Controllers/Base"
	"mpasGo/Src/Services"
	"net/http"
)

func PlacesController(router gin.IRouter) *Base.BaseController {
	controller := Base.NewBaseController(router)
	controller.GET("/", func(c *gin.Context) {

		var u Services.PlaceQuery
		err := c.BindQuery(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		// TODO:
		// - decide city, especially Dar
		// - get city from latlng
		// - Capitalize City and Country
		// - SearchDB
		// - Search OSM
		// - Search Google
		// - Capitalize Description from every response

		//places := Services.GoogleServiceInstance.AutoComplete(u)
		places, err := Services.MongoServiceInstance.AutoComplete(u)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		// params := c.Request.URL.Query()
		c.JSON(200, places)
	})

	//controller.PrintRoutes()

	return controller
}
