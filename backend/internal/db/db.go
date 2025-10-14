package db

import (
	"fmt"
	"log"

	"github.com/critiq17/critiqal-site/config"
	"github.com/critiq17/critiqal-site/internal/models"
	pgdriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func InitDB(cfg *config.DatabaseConfig) (*DB, error) {

	db, err := gorm.Open(pgdriver.Open(cfg.DSN()), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed connect to DB: %v", err)
	}

	log.Println("Connected to db successfully")

	return &DB{db}, nil
}

func Migrate(db *DB) error {

	if err := db.AutoMigrate(&models.User{}); err != nil {
		return fmt.Errorf("error migrating models: %v", err)
	}

	log.Println("Models migrated successfully!")
	return nil
}

func Must(cfg *config.DatabaseConfig) *DB {
	db, err := InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed init db: %v", err)
	}

	if err := Migrate(db); err != nil {
		log.Fatalf("Migrations failed: %v", err)
	}

	return db
}
