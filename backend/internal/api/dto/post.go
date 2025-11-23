package dto

import "github.com/critiq17/critiqal-site/internal/domain/post"

type PostCreateDTO struct {
	OwnerID     string  `json:"owner_id" binding:"required"`
	PhotoURL    *string `json:"photo_url"`
	Description string  `json:"description" binding:"required"`
}

type PostResponseDTO struct {
	ID          string  `json:"id"`
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

func ToPostDTO(p *post.Post) *PostResponseDTO {
	return &PostResponseDTO{
		ID:          p.ID,
		OwnerID:     p.OwnerID,
		PhotoURL:    p.PhotoURL,
		Description: p.Description,
	}
}

func ToPostsDTO(posts []*post.Post) []PostResponseDTO {
	dtos := make([]PostResponseDTO, len(posts))
	for i, u := range posts {
		dtos[i] = *ToPostDTO(u)
	}
	return dtos
}
