package routes

import (
	//_ "github.com/critiq17/critiqal-site/backend/docs"
	_ "github.com/critiq17/critiqal-site/docs"
	"github.com/critiq17/critiqal-site/internal/api/handlers"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func InitRoutes(app *fiber.App, handlers *handlers.Handlers) {

	// Swagger
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	api := app.Group("/api")

	users := api.Group("/users")
	{
		users.Post("/", handlers.AddUser)
		users.Delete("/:id", handlers.DeleteUser)
		users.Get("/", handlers.GetUsers)
	}

	auth := api.Group("/auth")
	{
		auth.Post("/sign-up", handlers.SignUp)
	}
}
