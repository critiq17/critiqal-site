package user

import (
	"github.com/critiq17/critiqal-site/internal/domain/user/dto"
)

type UserRepository interface {
	Create(u *dto.User) error
	Delete(id int) error
	GetUsers() ([]*dto.User, error)
}
