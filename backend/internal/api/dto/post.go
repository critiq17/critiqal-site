package dto

import "github.com/critiq17/critiqal-site/internal/domain/post"

type PostCreateDTO struct {
	OwnerID     string  `json:"owner_id" binding:"required"`
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
