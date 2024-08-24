package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"telkomsel-technical-test.com/dto"
	"telkomsel-technical-test.com/repository"
)

type Handler struct {
	repo repository.Repository
}

func NewHandler(repo repository.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) Fetch(c *fiber.Ctx) error {
	variety := c.Query("variety")
	brand := c.Query("brand")
	id := c.Query("id")

	resp, err := h.repo.FetchProduct(brand, variety, id)
	if err != nil {
		return dto.ResponseError(c, err, err.Error(), fiber.StatusInternalServerError)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    resp,
	})
}

func (h *Handler) Create(c *fiber.Ctx) error {
	req := &dto.Product{}

	err := c.BodyParser(req)
	if err != nil {
		return dto.ResponseError(c, err, "failed parsing request", fiber.StatusBadRequest)
	}

	err = h.repo.Create(req)
	if err != nil {
		return dto.ResponseError(c, err, "failed creating product", fiber.StatusBadRequest)
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "success",
	})
}

func (h *Handler) Update(c *fiber.Ctx) error {
	req := &dto.Product{}

	err := c.BodyParser(req)
	if err != nil {
		return dto.ResponseError(c, err, "failed parsing request", fiber.StatusBadRequest)
	}

	err = h.repo.Update(req)
	if err != nil {
		return dto.ResponseError(c, err, "failed creating product", fiber.StatusBadRequest)
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "success",
	})
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	product_id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return dto.ResponseError(c, err, "failed parsing id", fiber.StatusBadRequest)
	}
	err = h.repo.Delete(product_id)
	if err != nil {
		return dto.ResponseError(c, err, "failed creating product", fiber.StatusBadRequest)
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "success",
	})
}
