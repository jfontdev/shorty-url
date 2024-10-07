package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/shorten/:url", func(c *fiber.Ctx) error {
		params := c.Params("url")

		log.Print("URL to parse " + params)

		return c.JSON(params)
	})
}
