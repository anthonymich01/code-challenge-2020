package main

import (
	"backend/endpoints"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	// GET
	app.Get("/v1/gas-price", endpoints.GasPrice)

	app.Listen(":8080")
}
