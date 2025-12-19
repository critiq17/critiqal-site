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

	// endpoint where saves all users profile picture
	app.Static("/uploads", "./uploads")

	api := app.Group("/api")

	// login, register
	auth := api.Group("/auth")
	{
		auth.Post("/sign-up", handlers.SignUp)
		auth.Post("/sign-in", handlers.SignIn)
	}

	// for search, get, profile, photo
	users := api.Group("/users", handlers.UserIdentity)
	{
		// CRUD (without update)
		users.Post("/", handlers.CreateUser)
		users.Get("/", handlers.GetUsers)
		users.Get("/:username", handlers.GetByUsername)
		users.Delete("/:id", handlers.DeleteUser)

		// search by username
		users.Get("/search/:username", handlers.SearchUsers)

		// upload photo
		users.Post("/:username/photo", handlers.UploadPhoto)

		// retrieves full user information, without password, id
		users.Get("/me", handlers.GetMe)

	}

	posts := api.Group("/posts", handlers.UserIdentity)
	{
		// CRUD
		posts.Post("/", handlers.CreatePost)

		// retrieves 50 recents posts, by created_at
		posts.Get("/recent", handlers.GetRecentPosts)

		posts.Get("/:id", handlers.GetPost)
		posts.Put("/:id", handlers.UpdatePost)
		posts.Delete("/:id", handlers.DeletePost)

		// retrieves all posts by username
		posts.Get("/users/:username", handlers.GetPostsByUserName)
	}

}
