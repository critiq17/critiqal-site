package dto

import (
	"time"

	"github.com/critiq17/critiqal-site/internal/domain/post"
)

type PostCreateDTO struct {
	OwnerID     string  `json:"owner_id" binding:"required"`
	PhotoURL    *string `json:"photo_url"`
	Description string  `json:"description" binding:"required"`
}

type PostResponseDTO struct {
	ID        string     `json:"id"`
	Author    UserApi    `json:"author"`
	Title     *string    `json:"title,omitempty"`
	Body      string     `json:"body"`
	CreatedAt *time.Time `json:"created_at"`
	ImageURL  *string    `json:"image_url,omitempty"`
}

type PostUpdateDTO struct {
	PhotoURL    *string `json:"photo_url"`
	Description string  `json:"description" binding:"required"`
}

func ToPostDomain(p *PostCreateDTO) *post.Post {
	if p == nil {
		return nil
	}
	return &post.Post{
		OwnerID:     p.OwnerID,
		PhotoURL:    p.PhotoURL,
		Description: p.Description,
	}
}

func ToPostDTO(p *post.Post) *PostResponseDTO {
	return &PostResponseDTO{
		ID: p.ID,
		Author: UserApi{
			Username:  p.Owner.Username,
			Email:     p.Owner.Email,
			FirstName: p.Owner.FirstName,
			LastName:  p.Owner.LastName,
			PhotoURL:  p.Owner.PhotoURL,
		},
		Title:     p.Title,
		Body:      p.Description,
		CreatedAt: p.CreatedAt,
		ImageURL:  p.PhotoURL,
	}
}

func ToPostsDTO(posts []*post.Post) []PostResponseDTO {
	dtos := make([]PostResponseDTO, len(posts))
	for i, u := range posts {
		dtos[i] = *ToPostDTO(u)
	}
	return dtos
}

func ToPostDomainFromUpdateDTO(p *PostUpdateDTO) *post.Post {
	if p == nil {
		return nil
	}
	return &post.Post{
		PhotoURL:    p.PhotoURL,
		Description: p.Description,
	}
}
