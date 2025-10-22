package db

import (
	"fmt"
	"log"

	"github.com/critiq17/critiqal-site/config"
	"github.com/critiq17/critiqal-site/internal/domain/user/dto"
	pgdriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func Must(cfg *config.DatabaseConfig) *DB {
	db, err := setupDB(cfg)
	if err != nil {
		log.Fatalf("Failed init db: %v", err)
	}

	if err := migrate(db); err != nil {
		log.Fatalf("Migrations failed: %v", err)
	}

	return db
}

// ----- Helpers -----

// open connection to postgres
func setupDB(cfg *config.DatabaseConfig) (*DB, error) {

	db, err := gorm.Open(pgdriver.Open(cfg.DSN()), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed connect to DB: %v", err)
	}
	log.Println("Connected to db successfully")

	return &DB{db}, nil
}

// migrating models for DB
func migrate(db *DB) error {

	if err := db.AutoMigrate(&dto.User{}); err != nil {
		return fmt.Errorf("error migrating models: %v", err)
	}

	log.Println("Models migrated successfully!")
	return nil
}
