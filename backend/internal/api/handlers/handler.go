package handlers

import (
	"github.com/critiq17/critiqal-site/internal/domain/user"
)

type Handlers struct {
	userRepo user.Repository
}

func NewHandlers(userRepo user.Repository) *Handlers {
	return &Handlers{
		userRepo: userRepo,
	}
}
