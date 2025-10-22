package app

import (
	"github.com/critiq17/critiqal-site/config"
	"github.com/critiq17/critiqal-site/internal/api/handlers"
	"github.com/critiq17/critiqal-site/internal/api/routes"
	"github.com/critiq17/critiqal-site/internal/db"
	"github.com/critiq17/critiqal-site/internal/logger"
	"github.com/critiq17/critiqal-site/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupApp() (*fiber.App, error) {
	log := logger.New(logger.DEBUG)

	cfg := config.LoadConfig()

	db := db.Must(&cfg.DatabaseConfig)

	userRepo := repository.NewRepository(db.DB)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origins, Content-Type, Accept",
	}))

	handlers := handlers.NewHandlers(userRepo)
	routes.InitRoutes(app, handlers)

	log.Info("Success init db, handlers, and more")

	return app, nil
}
