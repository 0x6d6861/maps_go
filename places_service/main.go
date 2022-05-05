package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"mpasGo/Controllers"
	"mpasGo/Database"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
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

	v1 := router.Group("/v3")
	{
		places := v1.Group("/places")
		{
			//Controllers.HomeController(api)
			Controllers.PlacesController(places)
		}

	}

	router.Run(":3002") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
