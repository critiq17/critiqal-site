package main

import (
	"log"

	"github.com/critiq17/critiqal-site/config"
	_ "github.com/critiq17/critiqal-site/docs"
	"github.com/critiq17/critiqal-site/internal/app"
)

// @title Critiqal API
// @version 1.0
// @description This is  the API for Critiqal web-site
// @host localhost:8080
// @BasePath /api
// @schemes http
func main() {

	cfg := config.LoadConfig()
	app, err := app.SetupApp()
	if err != nil {
		log.Fatal("Error setup app: ", err)
	}

	log.Println("Starting on port: " + cfg.Server.PORT)
	app.Listen(":" + cfg.Server.PORT)
}
