package service

import (
	"context"

	"github.com/critiq17/critiqal-site/internal/domain/post"
)

type PostService struct {
	postRepo post.Repository
}

func NewPostService(postRepo post.Repository) *PostService {
	return &PostService{
		postRepo: postRepo,
	}
}

func (s *PostService) Create(ctx context.Context, post *post.Post) error {
	return s.postRepo.Create(ctx, post)
}

func (s *PostService) Get(ctx context.Context, id string) (*post.Post, error) {
	return s.postRepo.Get(ctx, id)
}

func (s *PostService) Update(ctx context.Context, id string, post *post.Post) error {
	return s.postRepo.Update(ctx, id, post)
}

func (s *PostService) Delete(ctx context.Context, id string) error {
	return s.postRepo.Delete(ctx, id)
}

func (s *PostService) GetPostsByUserID(ctx context.Context, user_id string) ([]*post.Post, error) {
	return s.postRepo.GetPostsByUserID(ctx, user_id)
}

func (s *PostService) GetRecentPosts(ctx context.Context, limit int) ([]*post.Post, error) {
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}

	return s.postRepo.GetRecent(ctx, limit)
}
