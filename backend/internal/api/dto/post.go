package dto

import (
	"time"

	"github.com/critiq17/critiqal-site/internal/domain/post"
)

type PostCreateDTO struct {
	OwnerID     string  `json:"owner_id" binding:"required"`
	PhotoURL    *string `json:"photo_url"`
	Description string  `json:"description" binding:"required"`
	Title       *string `json:"title"`
}

type PostResponseDTO struct {
	ID        string  `json:"id"`
	Author    UserApi `json:"author"`
	Title     *string `json:"title,omitempty"`
	Body      string  `json:"body"`
	CreatedAt string  `json:"created_at"`
	ImageURL  *string `json:"image_url,omitempty"`
}

type PostUpdateDTO struct {
	PhotoURL    *string `json:"photo_url"`
	Description string  `json:"description" binding:"required"`
	Title       *string `json:"title"`
}

func ToPostDomain(p *PostCreateDTO) *post.Post {
	if p == nil {
		return nil
	}
	return &post.Post{
		OwnerID:     p.OwnerID,
		PhotoURL:    p.PhotoURL,
		Description: p.Description,
		Title:       p.Title,
	}
}

func ToPostDTO(p *post.Post) *PostResponseDTO {
	dto := &PostResponseDTO{
		ID: p.ID,
		Author: UserApi{
			Username:  p.Owner.Username,
			Email:     p.Owner.Email,
			FirstName: p.Owner.FirstName,
			LastName:  p.Owner.LastName,
			PhotoURL:  p.Owner.PhotoURL,
		},
		Title:    p.Title,
		Body:     p.Description,
		ImageURL: p.PhotoURL,
	}

	if p.CreatedAt != nil {
		dto.CreatedAt = p.CreatedAt.Format(time.RFC3339)
	} else {
		dto.CreatedAt = time.Now().Format(time.RFC3339)
	}

	return dto
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
		Title:       p.Title,
	}
}
