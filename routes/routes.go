package routes

import (
	"crypto/sha256"
	"encoding/base64"
	"log"
	"shorty-url/database"
	"shorty-url/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type shortenBody struct {
	Url string `json:"url"`
}

type shortenResponse struct {
	OriginalURL  string `json:"original_url"`
	ShortenedURL string `json:"shortened_url"`
}

func RegisterRoutes(app *fiber.App) {
	app.Post("/shorten", func(c *fiber.Ctx) error {
		body := new(shortenBody)
		url := new(models.Url)

		if err := c.BodyParser(body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}

		originalUrl := body.Url

		if originalUrl == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "URL cannot be empty"})
		}

		log.Printf("URL to parse: %s", originalUrl)

		hashed := sha256.Sum256([]byte(originalUrl))
		log.Printf("Hashed: %x", hashed)

		encoded := base64.StdEncoding.EncodeToString(hashed[:])
		log.Println("Encoded: ", encoded)

		shortenedUrl := encoded[:12]
		log.Println("Shortened URL: ", shortenedUrl)

		url.OriginalURL = originalUrl
		url.ShortenedURL = shortenedUrl

		db := database.DB.Db

		if err := db.Create(&url).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create url", "data": err})
		}

		return c.JSON(shortenResponse{
			OriginalURL:  originalUrl,
			ShortenedURL: shortenedUrl,
		})
	})

	app.Get("/shorten/:hash", func(c *fiber.Ctx) error {
		hash := c.Params("hash")
		url := new(models.Url)

		println(hash)

		db := database.DB.Db

		result := db.Where("shortened_url = ?", hash).First(&url).Error

		if result != nil {
			if result == gorm.ErrRecordNotFound {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "URL not found"})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
		}

		return c.JSON(fiber.Map{"originalUrl": url.OriginalURL})
	})
}
