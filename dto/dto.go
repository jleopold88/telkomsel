package dto

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Product struct {
	ProductID          uuid.UUID `db:"product_id" json:"product_id"`
	ProductName        string    `db:"product_name" json:"product_name"`
	ProductDescription string    `db:"product_description" json:"product_description"`
	ProductPrice       float32   `db:"product_price" json:"product_price"`
	ProductVariety     string    `db:"product_variety" json:"product_variety"`
	ProductRating      float32   `db:"product_rating" json:"product_rating"`
	ProductStock       int32     `db:"product_stock" json:"product_stock"`
	ProductURL         string    `db:"product_url" json:"product_url"`
	ProductBrand       string    `db:"product_brand" json:"product_brand"`
}

func ResponseError(c *fiber.Ctx, err error, error_message string, status int) error {
	return c.Status(500).JSON(fiber.Map{
		"error":         err.Error(),
		"error_message": error_message,
	})
}
