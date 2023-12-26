package main

import (
	"errors"
	"fmt"
	"log"

	"product-listing-api/controller"
	"product-listing-api/db"
	"product-listing-api/model"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
}

func loadEnv() {
	if err := godotenv.Load(".env.local"); err != nil {
		handleError(errors.New("error loading .env file"))
	}
}

func loadDatabase() {
	if err := db.Connect(); err != nil {
		handleError(err)
	}

	if err := db.Database.AutoMigrate(&model.Product{}); err != nil {
		handleError(err)
	}
	fmt.Println("Migrations executed successfully")
}

func handleError(err error) {
	log.Fatal(err)
}

func serveApplication() {
	router := gin.Default()

	userRoutes := router.Group("/product")
	userRoutes.POST("", controller.CreateProduct)
	userRoutes.GET("", controller.GetAllProducts)
	userRoutes.GET("/:id", controller.GetProduct)
	userRoutes.PATCH("/:id", controller.UpdateProduct)
	userRoutes.DELETE("/:id", controller.DeleteProduct)

	if err := router.Run(":1234"); err != nil {
		handleError(err)
	}

	fmt.Println("Server running on port 1234")
}
