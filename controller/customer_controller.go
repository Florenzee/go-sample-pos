package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-sample-post/models"
	"go-sample-post/services"
)

type CustomerController struct {
	customerService services.CustomerService
}

func NewCustomerController(services services.CustomerService) *CustomerController {
	return &CustomerController{services}
}

func (c *CustomerController) CreateCustomer(ctx *fiber.Ctx) error {
	var customer models.Customer
	if err := ctx.BodyParser(&customer); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.customerService.CreateCustomer(&customer); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(customer)
}

func (c *CustomerController) GetAllCustomer(ctx *fiber.Ctx) error {
	customers, err := c.customerService.GetAllCustomers()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(customers)
}

func (c *CustomerController) GetCustomerByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	customer, err := c.customerService.GetCustomerByID(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Customer not found"})
	}
	return ctx.JSON(customer)
}

func (c *CustomerController) UpdateCustomer(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var customer models.Customer
	if err := ctx.BodyParser(&customer); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.customerService.UpdateCustomer(id, &customer); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(customer)
}

func (c *CustomerController) DeleteCustomer(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.customerService.DeleteCustomer(id); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}
