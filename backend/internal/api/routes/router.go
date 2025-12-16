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

	app.Static("/uploads", "./uploads")

	api := app.Group("/api")

	users := api.Group("/users", handlers.UserIdentity)
	{
		users.Post("/", handlers.AddUser)
		users.Delete("/:id", handlers.DeleteUser)
		users.Get("/", handlers.GetUsers)
		users.Post("/:username/photo", handlers.UploadPhoto)
		users.Get("/:username", handlers.GetByUsername)
		users.Get("/me", handlers.GetMe)

		search := users.Group("search")
		{
			search.Get("/:username", handlers.SearchUsers)
		}
	}

	posts := api.Group("/posts", handlers.UserIdentity)
	{
		posts.Post("/", handlers.CreatePost)
		//posts.Get("/:post_id", handlers.GetPost)
		posts.Get("/recent", handlers.GetRecentPosts)
		posts.Get("/:user_id", handlers.GetPostsByUserID)
		posts.Put("/:post_id", handlers.UpdatePost)
		posts.Delete("/:post_id", handlers.DeletePost)
	}

	auth := api.Group("/auth")
	{
		auth.Post("/sign-up", handlers.SignUp)
		auth.Post("/sign-in", handlers.SignIn)
	}
}
