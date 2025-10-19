package app

import (
	"fmt"
	"log"

	"github.com/critiq17/critiqal-site/config"
	"github.com/critiq17/critiqal-site/internal/api/routes"
	"github.com/critiq17/critiqal-site/internal/db"
	"github.com/critiq17/critiqal-site/internal/domain/user/dto"
	"github.com/critiq17/critiqal-site/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewServer() (*fiber.App, error) {

	cfg := config.LoadConfig()
	db := db.Must(&cfg.DatabaseConfig)

	user := &dto.User{
		Username:  "shlyapa",
		Email:     "zuzya@zzz",
		Password:  "jdsadbibif",
		FirstName: "Porion",
		LastName:  "Harison",
	}

	userRepo := repository.NewRepository(db.DB)

	err := userRepo.AddUser(user)
	if err != nil {
		log.Printf("error creating user: %v", err)
	}

	fmt.Println(userRepo.GetUsers())

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origins, Content-Type, Accept",
	}))

	routes.InitRoutes(app)

	return app, err
}
