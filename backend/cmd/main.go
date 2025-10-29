package main

import (
	"log"

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

	// run init server and DB
	app, err := app.SetupApp()
	if err != nil {
		log.Fatalf("Error setup app")
	}

	app.Listen(":8080")
}
