package handlers

import (
	"github.com/critiq17/critiqal-site/internal/domain/user"
)

type Handlers struct {
	userRepo user.Repository
	service  user.UserService
}

func NewHandlers(userRepo user.Repository, service user.UserService) *Handlers {
	return &Handlers{
		userRepo: userRepo, service: service,
	}
}
