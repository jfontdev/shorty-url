package routes

import (
	"crypto/sha256"
	"encoding/base64"
	"log"

	"github.com/gofiber/fiber/v2"
)

type shortenBody struct {
	Url string `json:"url"`
}

func RegisterRoutes(app *fiber.App) {
	app.Post("/shorten", func(c *fiber.Ctx) error {
		body := new(shortenBody)

		if err := c.BodyParser(body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}

		if body.Url == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "URL cannot be empty"})
		}

		log.Printf("URL to parse: %s", body.Url)

		hashed := sha256.Sum256([]byte(body.Url))
		log.Printf("Hashed: %x", hashed)

		encoded := base64.StdEncoding.EncodeToString(hashed[:])
		log.Println("Encoded: ", encoded)

		shortenedUrl := encoded[:12]
		log.Println("Shortened URL: ", shortenedUrl)

		return c.JSON(fiber.Map{"shortened_url": shortenedUrl, "original_url": body.Url})
	})
}
