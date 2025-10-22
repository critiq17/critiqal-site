package routes

import (
	"github.com/critiq17/critiqal-site/internal/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App, handlers *handlers.Handlers) {
	api := app.Group("/api")

	users := api.Group("/users")
	{
		users.Post("/", handlers.AddUser)
		users.Post("/:id", handlers.DeleteUser)
	}
}
