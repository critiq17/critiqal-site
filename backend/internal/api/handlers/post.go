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

func (h *Handlers) GetPost(ctx *fiber.Ctx) error {
	post_id := ctx.Params("post_id")

	if post_id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id required",
		})
	}

	post, err := h.postService.Get(context.Background(), post_id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "post not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.ToPostDTO(post))
}

func (h *Handlers) UpdatePost(ctx *fiber.Ctx) error {

	id := ctx.Params("post_id")

	var req dto.PostUpdateDTO

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid body type",
		})
	}

	if err := h.postService.Update(context.Background(), id, dto.ToPostDomainFromUpdateDTO(&req)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "successfuly updated post"})
}

func (h *Handlers) DeletePost(ctx *fiber.Ctx) error {
	id := ctx.Params("post_id")

	if err := h.postService.Delete(context.Background(), id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "successfuly deleted post"})
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
