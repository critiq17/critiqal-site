package handlers

import (
	"context"

	"github.com/critiq17/critiqal-site/internal/api/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) CreatePost(ctx *fiber.Ctx) error {
	var req dto.PostCreateDTO

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid json body",
		})
	}

	if err := h.postService.Create(context.Background(), dto.ToPostDomain(&req)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(req)
}

func (h *Handlers) GetPostsByUserID(ctx *fiber.Ctx) error {
	user_id := ctx.Params("user_id")
	posts, err := h.postService.GetPostsByUserID(context.Background(), user_id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "not found posts",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.ToPostsDTO(posts))
}
