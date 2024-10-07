package main

import (
	"shorty-url/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.RegisterRoutes(app)

	app.Listen("localhost:3000")
}
