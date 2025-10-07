package routes

import (
	"github.com/critiq17/critiqal-site/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", handlers.HandleMessage)
}
