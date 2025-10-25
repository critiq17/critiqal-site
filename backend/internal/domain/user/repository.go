package user

import (
	"github.com/critiq17/critiqal-site/internal/domain/user/dto"
)

type Repository interface {
	Create(u *dto.User) error
	Delete(id string) error
	GetUsers() ([]dto.User, error)

	GetUserByUsername(username string) (*dto.User, error)
}
