package post

import (
	"time"

	"github.com/critiq17/critiqal-site/internal/domain/user"
)

type Post struct {
	ID          string
	OwnerID     string
	Title       *string
	PhotoURL    *string
	Description string
	CreatedAt   *time.Time
	DeletedAt   *time.Time
	Owner       user.User
}
