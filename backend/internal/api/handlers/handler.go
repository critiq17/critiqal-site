package handlers

import (
	"github.com/critiq17/critiqal-site/internal/service"
)

type Handlers struct {
	service *service.UserService
}

func NewHandlers(service *service.UserService) *Handlers {
	return &Handlers{
		service: service,
	}
}
