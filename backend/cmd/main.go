package main

import (
	_ "github.com/critiq17/critiqal-site/docs"
	"github.com/critiq17/critiqal-site/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title Critiqal API
// @version 1.0
// @description This is the API for Critiqal web-site
// @host localhost:8080
// @BasePath /api
// @schemes http
func main() {
	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)
	routes.InitRoutes(app)
	app.Listen(":8080")

}
