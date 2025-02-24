package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-sample-post/controller"
	"go-sample-post/database"
	"go-sample-post/repositories"
	"go-sample-post/routes"
	"go-sample-post/services"
)

func main() {
	app := fiber.New()

	//connet to database
	database.ConnectDatabase()

	//dependency injection
	//product
	productRepo := repositories.NewProductRepository(database.DB)
	productService := services.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	//customer
	customerRepo := repositories.NewCustomerRepository(database.DB)
	customerService := services.NewCustomerService(customerRepo)
	customerController := controller.NewCustomerController(customerService)

	//setup routes
	routes.SetupRoutes(app, productController, customerController)

	//start server
	err := app.Listen(":8080")
	if err != nil {
		fmt.Println(err)
	}

}
