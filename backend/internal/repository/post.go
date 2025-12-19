package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/critiq17/critiqal-site/internal/domain/post"
	"github.com/critiq17/critiqal-site/internal/domain/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostModel struct {
	ID          string `gorm:"primaryKey;not null"`
	OwnerID     string `gorm:"index;not null"`
	Title       *string
	PhotoURL    *string
	Description string `gorm:"not null"`
	CreatedAt   *time.Time
	DeletedAt   *time.Time `gorm:"index"`

	Owner User `gorm:"foreignKey:OwnerID;references:ID" json:"author"`
}

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

// toDomainPost converts database model to domain model
func toDomainPost(p *PostModel) *post.Post {
	domainPost := &post.Post{
		ID:          p.ID,
		OwnerID:     p.OwnerID,
		Title:       p.Title,
		PhotoURL:    p.PhotoURL,
		Description: p.Description,
		CreatedAt:   p.CreatedAt,
		DeletedAt:   p.DeletedAt,
	}

	// Map Owner if exists
	if p.Owner.ID != "" {
		domainPost.Owner = user.User{
			ID:        p.Owner.ID,
			Username:  p.Owner.Username,
			Email:     p.Owner.Email,
			FirstName: p.Owner.FirstName,
			LastName:  p.Owner.LastName,
			PhotoURL:  p.Owner.PhotoURL,
		}
	}

	return domainPost
}

// toDomainPosts converts slice of database models to domain models
func toDomainPosts(models []*PostModel) []*post.Post {
	posts := make([]*post.Post, len(models))
	for i, m := range models {
		posts[i] = toDomainPost(m)
	}
	return posts
}

// toModelPost converts domain model to database model
func toModelPost(p *post.Post) *PostModel {
	return &PostModel{
		ID:          p.ID,
		OwnerID:     p.OwnerID,
		Title:       p.Title,
		PhotoURL:    p.PhotoURL,
		Description: p.Description,
		CreatedAt:   p.CreatedAt,
		DeletedAt:   p.DeletedAt,
	}
}

// BeforeCreate generates UUID and sets timestamp
func (p *PostModel) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == "" {
		p.ID = uuid.NewString()
	}
	now := time.Now()
	if p.CreatedAt == nil {
		p.CreatedAt = &now
	}
	return nil
}

// Create inserts a new post using userID directly from JWT
func (r *PostRepository) Create(ctx context.Context, p *post.Post) error {
	model := toModelPost(p)

	if model.ID == "" {
		model.ID = uuid.NewString()
	}

	if model.CreatedAt == nil {
		now := time.Now()
		model.CreatedAt = &now
	}

	err := r.db.WithContext(ctx).Create(model).Error
	if err != nil {
		return err
	}

	p.ID = model.ID
	p.CreatedAt = model.CreatedAt

	return nil
}

// Get retrieves a single post by ID
func (r *PostRepository) Get(ctx context.Context, id string) (*post.Post, error) {
	var model PostModel
	err := r.db.WithContext(ctx).
		Preload("Owner").
		Where("id = ? AND deleted_at IS NULL", id).
		First(&model).Error

	if err != nil {
		return nil, err
	}

	return toDomainPost(&model), nil
}

// Update modifies an existing post
func (r *PostRepository) Update(ctx context.Context, id string, p *post.Post) error {
	updates := map[string]interface{}{}

	if p.Title != nil {
		updates["title"] = p.Title
	}
	if p.PhotoURL != nil {
		updates["photo_url"] = p.PhotoURL
	}
	if p.Description != "" {
		updates["description"] = p.Description
	}

	return r.db.WithContext(ctx).
		Model(&PostModel{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(updates).
		Error
}

// Delete soft deletes a post
func (r *PostRepository) Delete(ctx context.Context, id string) error {
	now := time.Now()
	return r.db.WithContext(ctx).
		Model(&PostModel{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Update("deleted_at", now).
		Error
}

// GetPostsByUserID retrieves all posts by userID
func (r *PostRepository) GetPostsByUserID(ctx context.Context, userID string) ([]*post.Post, error) {
	var models []*PostModel

	err := r.db.WithContext(ctx).
		Preload("Owner").
		Where("owner_id = ? AND deleted_at IS NULL", userID).
		Order("created_at DESC").
		Find(&models).Error

	if err != nil {
		return nil, err
	}

	return toDomainPosts(models), nil
}

// GetRecent retrieves most recent posts
// В repository/post.go - GetRecent
func (r *PostRepository) GetRecent(ctx context.Context, limit int) ([]*post.Post, error) {
	var models []*PostModel

	fmt.Printf("GetRecent called with limit=%d\n", limit)

	err := r.db.WithContext(ctx).
		Debug().
		Preload("Owner").
		Where("deleted_at IS NULL").
		Order("created_at DESC").
		Limit(limit).
		Find(&models).Error

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	fmt.Printf("Found %d models\n", len(models))
	for i, m := range models {
		fmt.Printf("  Model %d: ID=%s, OwnerID=%s, Owner.ID='%s', Owner.Username='%s'\n",
			i, m.ID, m.OwnerID, m.Owner.ID, m.Owner.Username)
	}

	posts := toDomainPosts(models)
	fmt.Printf("Converted to %d domain posts\n", len(posts))

	if len(posts) > 0 {
		fmt.Printf("  First post Owner: ID='%s', Username='%s'\n",
			posts[0].Owner.ID, posts[0].Owner.Username)
	}

	return posts, nil
}
