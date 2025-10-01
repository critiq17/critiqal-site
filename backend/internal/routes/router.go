package routes

import (
	"github.com/critiq17/critiqal-site/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func InitRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Use(cors.New())

	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{})
	})
	api.Post("/", func(c *fiber.Ctx) error {
		return handlers.HandleMessage(c)
	})
}
