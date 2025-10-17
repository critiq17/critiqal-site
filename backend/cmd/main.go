package main

import (
	"fmt"

	"github.com/critiq17/critiqal-site/config"
	_ "github.com/critiq17/critiqal-site/docs"
	"github.com/critiq17/critiqal-site/internal/db"
	"github.com/critiq17/critiqal-site/internal/dto/api"
	"github.com/critiq17/critiqal-site/internal/dto/models"
	"github.com/critiq17/critiqal-site/internal/repository"
)

// @title Critiqal API
// @version 1.0
// @description This is  the API for Critiqal web-site
// @host localhost:8080
// @BasePath /api
// @schemes http
func main() {

	cfg := config.LoadConfig()
	db := db.Must(&cfg.DatabaseConfig)

	user := &models.User{
		Username:  "shlyapa",
		Email:     "zuzya@zzz",
		Password:  "jdsadbibif",
		FirstName: "Porion",
		LastName:  "Harison",
	}

	userRepo := repository.NewRepository(db.DB)

	//err := userRepo.AddUser(user)
	//if err != nil {
	//	log.Printf("error creating user: %v", err)
	//}

	//err = userRepo.SoftDelete(user.ID)
	//if err != nil {
	//	log.Printf("error delete user: %v", err)
	//
	fmt.Println(api.ToUserApi(user))

	fmt.Println(userRepo.GetUsers())

}
