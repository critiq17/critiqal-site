package repository

import (
	"github.com/critiq17/critiqal-site/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUsers() (*[]models.User, error) {
	var models []models.User

	err := r.db.
		Where("deleted_at IS NULL").
		Order("created_at DESC").
		Find(&models).Error

	if err != nil {
		return nil, err
	}

	return &models, nil
}
