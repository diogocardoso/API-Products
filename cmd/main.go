package main

import (
	"API-Products/internal/controller"
	"API-Products/internal/infra/database"
	"API-Products/internal/repository"
	"API-Products/internal/usecase"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// Inicia a conexão
	dbConnection, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConnection.Close()

	// Verifica se a conexão não é nil
	if dbConnection == nil {
		log.Fatal("dbConnection is nil")
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")

	// Repositories
	productRepository := repository.NewProductRepository(dbConnection)
	// Usecases
	productUsecase := usecase.NewProductUseCase(productRepository)
	// Controllers
	productController := controller.NewProductController(productUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	server.GET("/products", productController.GetProducts)
	server.POST("/product", productController.CreateProduct)
	server.GET("/product/:productId", productController.GetProductById)

	server.Run(":8000")
}
