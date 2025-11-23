package repository

import (
	"context"
	"time"

	"github.com/critiq17/critiqal-site/internal/domain/post"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostModel struct {
	ID          string `gorm:"primaryKey;not null"`
	OwnerID     string `gorm:"index;not null"`
	PhotoURL    *string
	Description string `gorm:"not null"`
	CreatedAt   *time.Time
	DeletedAt   *time.Time `gorm:"index"`
}

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func toDomainPost(p *PostModel) *post.Post {
	return &post.Post{
		ID:          p.ID,
		OwnerID:     p.OwnerID,
		PhotoURL:    p.PhotoURL,
		Description: p.Description,
		CreatedAt:   p.CreatedAt,
		DeletedAt:   p.DeletedAt,
	}
}
func toDomainPosts(models []*PostModel) []*post.Post {
	posts := make([]*post.Post, len(models))
	for i, m := range models {
		posts[i] = toDomainPost(m)
	}
	return posts
}

func fromDomainPost(p *post.Post) *PostModel {
	return &PostModel{
		ID:          p.ID,
		OwnerID:     p.OwnerID,
		PhotoURL:    p.PhotoURL,
		Description: p.Description,
		CreatedAt:   p.CreatedAt,
		DeletedAt:   p.DeletedAt,
	}
}

func fromDomainPosts(posts []post.Post) []PostModel {
	models := make([]PostModel, len(posts))
	for i, p := range posts {
		models[i] = *fromDomainPost(&p)
	}
	return models
}

func (r *PostRepository) Create(ctx context.Context, post *post.Post) error {

	err := r.db.WithContext(ctx).Create(fromDomainPost(post)).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) Get(ctx context.Context, id string) (*post.Post, error) {
	var model PostModel
	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&model).Error
	if err != nil {
		return nil, err
	}
	return toDomainPost(&model), nil
}

func (r *PostRepository) Update(ctx context.Context, post *post.Post) error {
	updates := map[string]interface{}{
		"photo_url":   post.PhotoURL,
		"description": post.Description,
	}

	return r.db.
		Model(&PostModel{}).
		Where("id = ?", post.ID).
		Updates(updates).
		Error
}

func (r *PostRepository) Delete(ctx context.Context, id string) error {
	var model PostModel
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&model).Error
}

func (r *PostRepository) GetPostsByUserID(ctx context.Context, user_id string) ([]*post.Post, error) {
	var models []*PostModel

	err := r.db.WithContext(ctx).
		Where("owner_id = ?", user_id).
		Order("created_at DESC").
		Find(&models).Error

	if err != nil {
		return nil, err
	}

	return toDomainPosts(models), nil
}

func (p *PostModel) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == "" {
		p.ID = uuid.NewString()
	}
	return nil
}
