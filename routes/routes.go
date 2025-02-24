package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-sample-post/controller"
)

// mendaftarkan endpoint API.
func SetupRoutes(app *fiber.App,
	productController *controller.ProductController,
	customerController *controller.CustomerController,
	employeeController *controller.EmployeeController,
	receiptController *controller.ReceiptController) {

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

	//customer Routes
	employeeGroup := app.Group("/employees")
	employeeGroup.Post("/", employeeController.CreateEmployee)
	employeeGroup.Get("/", employeeController.DeleteEmployee)
	employeeGroup.Get("/:id", employeeController.GetEmployeeByID)
	employeeGroup.Put("/:id", employeeController.UpdateEmployee)
	employeeGroup.Delete("/:id", employeeController.DeleteEmployee)

	//receipt routes
	receiptGroup := app.Group("/employees")
	receiptGroup.Post("/", receiptController.CreateReceipt)
	receiptGroup.Get("/", receiptController.DeleteReceipt)
	receiptGroup.Get("/:id", receiptController.GetReceiptByID)
	receiptGroup.Put("/:id", receiptController.UpdateReceipt)
	receiptGroup.Delete("/:id", receiptController.DeleteReceipt)

}
