package main

import (
	"direction_service/Controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		//log.Fatal("Error loading .env file")
	}
}

func main() {

	router := gin.Default()

	direction := router.Group("/direction")
	{
		//Controllers.HomeController(api)
		v2 := direction.Group("/v2")
		{
			Controllers.HomeController(v2)

		}

	}

	router.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
