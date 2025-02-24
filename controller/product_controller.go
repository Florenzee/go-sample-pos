package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-sample-post/models"
	"go-sample-post/services"
)

// mengatur request dan response dari client (memanggil repository)
type ProductController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{productService}
}

func (c *ProductController) CreateProduct(ctx *fiber.Ctx) error {
	var product models.Product
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.productService.CreateProduct(&product); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(product)
}

func (c *ProductController) CreateProductV2(ctx *fiber.Ctx) error {
	var product []models.Product
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	var productErrors []models.ProductError

	for _, p := range product {
		if err := c.productService.CreateProduct(&p); err != nil {
			productErrors = append(productErrors, models.ProductError{
				Product: p,
				Error:   err,
			})
		}
	}

	if len(productErrors) > 0 {
		return ctx.Status(500).JSON(fiber.Map{"error": productErrors})
	}
	return ctx.Status(201).JSON(product)
}

// bulk insert products
func (c *ProductController) CreateProductsBulk(ctx *fiber.Ctx) error {
	var products []models.Product

	// Parse JSON request
	if err := ctx.BodyParser(&products); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	// Insert products
	if err := c.productService.CreateProductBulk(&products); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert products",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Products added successfully",
	})
}

func (c *ProductController) GetAllProducts(ctx *fiber.Ctx) error {
	products, err := c.productService.GetAllProducts()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(products)
}

func (c *ProductController) GetProductsByID(ctx *fiber.Ctx) error {
	products, err := c.productService.GetAllProducts()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(products)
}

func (c *ProductController) UpdateProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var product models.Product
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.productService.UpdateProduct(id, &product); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(product)
}

func (c *ProductController) DeleteProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.productService.DeleteProduct(id); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}
