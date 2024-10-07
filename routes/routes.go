package routes

import (
	"crypto/sha256"
	"encoding/base64"
	"log"

	"github.com/gofiber/fiber/v2"
)

// TODO: Change this to a POST method and use the body
func RegisterRoutes(app *fiber.App) {
	app.Get("/shorten/:url", func(c *fiber.Ctx) error {
	params := c.Params("url")
    log.Printf("URL to parse: %s", params)

    hashed := sha256.Sum256([]byte(params))
    log.Printf("Hashed: %x", hashed)

    encoded := base64.StdEncoding.EncodeToString(hashed[:])
    log.Println("Encoded: ", encoded)

    shortenedUrl := encoded[:12] 
    log.Println("Shortened URL: ", shortenedUrl)

    return c.JSON(fiber.Map{"shortened_url": shortenedUrl, "original_url": params})
	})
}
