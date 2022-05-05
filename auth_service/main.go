package main

import (
	"auth_service/Controllers"
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

	// router.Use(Middleware.HasKey())

	auth := router.Group("/auth")
	{
		authV1 := auth.Group("/v1")
		{
			//Controllers.HomeController(api)
			Controllers.HomeController(authV1)
		}

	}

	router.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
