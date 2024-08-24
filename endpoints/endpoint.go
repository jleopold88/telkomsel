package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"telkomsel-technical-test.com/config"
	"telkomsel-technical-test.com/handlers"
	"telkomsel-technical-test.com/repository"
)

func ProductRoute(app *fiber.App) {
	// Initialize the repository
	repo := repository.NewRepository(config.DB)

	// Create the handler with the repository
	handler := handlers.NewHandler(repo)
	userGroup := app.Group("/v1/product")

	// Create
	userGroup.Post("/create", handler.Create)

	// Read
	userGroup.Get("/fetch", handler.Fetch)

	// Update
	userGroup.Put("/update", handler.Update)

	// Delete
	userGroup.Delete("/delete/:id", handler.Delete)
}
