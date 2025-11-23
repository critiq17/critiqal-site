package post

import "context"

type Repository interface {

	// CRUD
	Create(ctx context.Context, post *Post) error
	Get(ctx context.Context, id string) (*Post, error)
	Update(ctx context.Context, id string, post *Post) error
	Delete(ctx context.Context, id string) error

	// Getters
	GetPostsByUserID(ctx context.Context, user_id string) ([]*Post, error)
}
