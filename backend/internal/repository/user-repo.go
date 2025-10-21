package repository

import (
	"log"

	"github.com/critiq17/critiqal-site/internal/domain/user/dto"
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

func (r *UserRepository) Create(user *dto.User) error {
	err := r.db.Create(user).Error
	if err != nil {
		log.Printf("error creating user: %s", err)
	}

	return nil
}

func (r *UserRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&dto.User{}).Error
}

func (r *UserRepository) GetUsers() (*[]dto.User, error) {
	var models []dto.User

	err := r.db.
		Where("deleted_at IS NULL").
		Order("created_at DESC").
		Find(&models).Error

	if err != nil {
		return nil, err
	}

	return &models, nil
}
