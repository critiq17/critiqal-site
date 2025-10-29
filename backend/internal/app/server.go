package app

import (
	"github.com/critiq17/critiqal-site/config"
	"github.com/critiq17/critiqal-site/internal/api/handlers"
	"github.com/critiq17/critiqal-site/internal/api/routes"
	"github.com/critiq17/critiqal-site/internal/db"
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

	userRepo := repository.NewRepository(db.DB)
	service := service.NewUserService(userRepo)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origins, Content-Type, Accept",
	}))

	handlers := handlers.NewHandlers(userRepo, *service)
	routes.InitRoutes(app, handlers)

	log.Info("Success init db, handlers, and more")

	return app, nil
}
