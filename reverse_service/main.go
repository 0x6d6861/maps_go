package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"reverse_service/Controllers"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		//log.Fatal("Error loading .env file")
	}
}

func main() {

	router := gin.Default()

	//router.Use(Middleware.HasCorrectParams())

	reverse := router.Group("/reverse")
	{
		v1 := reverse.Group("/v1")
		{
			Controllers.HomeController(v1)

		}
	}

	router.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
