package app

import (
	"os"
	"strings"

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

	normalizeCSV := func(s string) string {
		parts := strings.Split(s, ",")
		out := make([]string, 0, len(parts))
		for _, p := range parts {
			p = strings.TrimSpace(p)
			if p != "" {
				out = append(out, p)
			}
		}
		return strings.Join(out, ",")
	}

	allowOrigins := normalizeCSV(os.Getenv("CORS_ALLOW_ORIGINS"))
	if allowOrigins == "" {
		allowOrigins = "http://localhost:5173,http://localhost:3001,http://localhost:81"
	}

	allowHeaders := os.Getenv("CORS_ALLOW_HEADERS")
	if strings.TrimSpace(allowHeaders) == "" {
		allowHeaders = "Origin, Content-Type, Accept, Authorization"
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     allowHeaders,
		AllowCredentials: true,
	}))

	handlers := handlers.NewHandlers(userService, postService)
	routes.InitRoutes(app, handlers)

	log.Info("Success init db, handlers, and more")
	db.Debug()

	return app, nil
}
