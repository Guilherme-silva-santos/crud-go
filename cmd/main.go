package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	usecase "go-api/useCase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	productController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/product", productController.GetProducts)
	server.POST("/products", productController.CreateProduct)

	server.Run(":8080")
}
