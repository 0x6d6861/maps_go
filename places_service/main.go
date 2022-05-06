package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"mpasGo/Controllers"
	"mpasGo/Database"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		//log.Fatal("Error loading .env file")
	}
}

func main() {

	defer func() {
		if err := Database.DatabaseInstance.Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	router := gin.Default()

	// router.Use(Middleware.HasKey())

	places := router.Group("/places")
	{
		v3 := places.Group("/v3")
		{
			//Controllers.HomeController(api)
			Controllers.PlacesController(v3)
		}

	}

	router.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
