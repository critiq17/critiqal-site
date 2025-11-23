package handlers

import (
	"github.com/critiq17/critiqal-site/internal/service"
)

type Handlers struct {
	userService *service.UserService
	postService *service.PostService
}

func NewHandlers(userService *service.UserService, postService *service.PostService) *Handlers {
	return &Handlers{
		userService: userService, postService: postService,
	}
}
