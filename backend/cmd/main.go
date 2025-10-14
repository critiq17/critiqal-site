package main

import (
	"github.com/critiq17/critiqal-site/config"
	_ "github.com/critiq17/critiqal-site/docs"
	"github.com/critiq17/critiqal-site/internal/db"
)

// @title Critiqal API
// @version 1.0
// @description This is the API for Critiqal web-site
// @host localhost:8080
// @BasePath /api
// @schemes http
func main() {

	cfg := config.LoadConfig()

	db.Must(&cfg.DatabaseConfig)
}
