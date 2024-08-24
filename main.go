package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"telkomsel-technical-test.com/config"
	"telkomsel-technical-test.com/endpoints"
)

func main() {
	app := fiber.New()

	app.Use(helmet.New(helmet.ConfigDefault))
	// Connect db
	config.ConnectDB()
	defer config.DB.Close()

	// Endpoints
	endpoints.ProductRoute(app)

	// Start Server
	log.Fatal(app.Listen(":3000"))
}
