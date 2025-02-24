package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-sample-post/controller"
)

// mendaftarkan endpoint API.
func SetupRoutes(app *fiber.App, productController *controller.ProductController, customerController *controller.CustomerController, employeeController *controller.EmployeeController) {

	//product routes
	productGroup := app.Group("/products")
	productGroup.Post("/", productController.CreateProduct)
	productGroup.Post("/bulk", productController.CreateProductsBulk)
	productGroup.Get("/", productController.GetAllProducts)
	productGroup.Get("/:id", productController.GetProductsByID)
	productGroup.Put("/:id", productController.UpdateProduct)
	productGroup.Delete("/:id", productController.DeleteProduct)

	//customer routes
	customerGroup := app.Group("/customers")
	customerGroup.Post("/", customerController.CreateCustomer)
	customerGroup.Get("/", customerController.GetAllCustomer)
	customerGroup.Get("/:id", customerController.GetCustomerByID)
	customerGroup.Put("/:id", customerController.UpdateCustomer)
	customerGroup.Delete("/:id", customerController.DeleteCustomer)

}
