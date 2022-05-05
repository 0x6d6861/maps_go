package main

import (
	"direction_service/Controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	router := gin.Default()

	v2 := router.Group("/v2")
	{
		//Controllers.HomeController(api)
		direction := v2.Group("/direction")
		{
			Controllers.HomeController(direction)

		}

	}

	router.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
