package main

import (
	"shorty-url/database"
	"shorty-url/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(logger.New())

	routes.RegisterRoutes(app)

	app.Listen("localhost:3000")
}
