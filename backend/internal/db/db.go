package db

import (
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func (db *DB) Migrate() error {
	return nil
}
