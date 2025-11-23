package app

import (
	"github.com/critiq17/critiqal-site/config"
	"github.com/critiq17/critiqal-site/internal/api/handlers"
	"github.com/critiq17/critiqal-site/internal/api/routes"
	"github.com/critiq17/critiqal-site/internal/db"
	"github.com/critiq17/critiqal-site/internal/storage"
	"github.com/critiq17/critiqal-site/pkg/logger"

	"github.com/critiq17/critiqal-site/internal/repository"
	"github.com/critiq17/critiqal-site/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupApp() (*fiber.App, error) {
	log := logger.New(logger.DEBUG)

	cfg := config.LoadConfig()

	db := db.Must(&cfg.DatabaseConfig)
	storage, err := storage.NewStorageFromEnv()

	if err != nil || storage == nil {
		log.Warn("Falling back to local storage...")
	}
	userRepo := repository.NewRepository(db.DB)
	postRepo := repository.NewPostRepository(db.DB)
	userService := service.NewUserService(userRepo, storage)
	postService := service.NewPostService(postRepo)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders: "Origins, Content-Type, Accept, Authorization",
	}))

	handlers := handlers.NewHandlers(userService, postService)
	routes.InitRoutes(app, handlers)

	log.Info("Success init db, handlers, and more")

	return app, nil
}
