package handlers

import (
	"github.com/critiq17/critiqal-site/internal/domain/user"
	"github.com/critiq17/critiqal-site/internal/service"
)

type Handlers struct {
	userRepo user.Repository
	service  *service.UserService
}

func NewHandlers(userRepo user.Repository, service *service.UserService) *Handlers {
	return &Handlers{
		userRepo: userRepo, service: service,
	}
}
