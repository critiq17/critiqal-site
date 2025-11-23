package post

import "time"

type Post struct {
	ID          string
	OwnerID     string
	PhotoURL    *string
	Description string
	CreatedAt   *time.Time
	DeletedAt   *time.Time
}
