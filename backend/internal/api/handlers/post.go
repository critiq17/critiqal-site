package handlers

import (
	"context"
	"fmt"

	"github.com/critiq17/critiqal-site/internal/api/dto"
	"github.com/gofiber/fiber/v2"
)

// CreatePost creates a new post
// @Summary Create post
// @Description Create a new post
// @Tags posts
// @Accept json
// @Produce json
// @Param post body dto.PostCreateDTO true "Post data"
// @Success 201 {object} map[string]string "post created"
// @Failure 400 {object} map[string]string "invalid json"
// @Failure 500 {object} map[string]string "server error"
// @Router /api/posts [post]
func (h *Handlers) CreatePost(c *fiber.Ctx) error {
	var input dto.PostCreateDTO

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	userID, ok := c.Locals("user_id").(string)
	if !ok || userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user_id not found in context",
		})
	}

	input.OwnerID = userID
	post := dto.ToPostDomain(&input)

	err := h.postService.Create(context.Background(), post)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create post",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "post created successfully",
		"post_id": post.ID,
	})
}

// GetPost retrieves a single post by ID
// @Summary Get post
// @Description Get post by ID
// @Tags posts
// @Produce json
// @Param post_id path string true "Post ID"
// @Success 200 {object} dto.PostResponseDTO
// @Failure 400 {object} map[string]string "id required"
// @Failure 500 {object} map[string]string "post not found"
// @Router /api/posts/{post_id} [get]
func (h *Handlers) GetPost(ctx *fiber.Ctx) error {
	postID := ctx.Params("post_id")

	if postID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id required",
		})
	}

	post, err := h.postService.Get(context.Background(), postID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "post not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.ToPostDTO(post))
}

// UpdatePost updates an existing post
// @Summary Update post
// @Description Update post by ID
// @Tags posts
// @Accept json
// @Produce json
// @Param post_id path string true "Post ID"
// @Param post body dto.PostUpdateDTO true "Updated post data"
// @Success 200 {object} map[string]string "successfully updated"
// @Failure 400 {object} map[string]string "invalid body"
// @Failure 403 {object} map[string]string "not authorized"
// @Failure 404 {object} map[string]string "post not found"
// @Failure 500 {object} map[string]string "server error"
// @Router /api/posts/{post_id} [put]
func (h *Handlers) UpdatePost(ctx *fiber.Ctx) error {

	postID := ctx.Params("post_id")
	userID := ctx.Locals("user_id").(string)

	var req dto.PostUpdateDTO
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid body type",
		})
	}

	existingPost, err := h.postService.Get(context.Background(), postID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "post not found",
		})
	}

	if existingPost.OwnerID != userID {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "not authorized to update this post",
		})
	}

	if err := h.postService.Update(context.Background(), postID, dto.ToPostDomainFromUpdateDTO(&req)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "successfully updated post",
	})
}

// DeletePost deletes a post
// @Summary Delete post
// @Description Delete post by ID
// @Tags posts
// @Param post_id path string true "Post ID"
// @Success 200 {object} map[string]string "successfully deleted"
// @Failure 403 {object} map[string]string "not authorized"
// @Failure 404 {object} map[string]string "post not found"
// @Failure 500 {object} map[string]string "server error"
// @Router /api/posts/{post_id} [delete]
func (h *Handlers) DeletePost(ctx *fiber.Ctx) error {
	postID := ctx.Params("post_id")
	userID := ctx.Locals("user_id").(string)

	existingPost, err := h.postService.Get(context.Background(), postID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "post not found",
		})
	}

	if existingPost.OwnerID != userID {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "not authorized to delete this post",
		})
	}

	if err := h.postService.Delete(context.Background(), postID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "successfully deleted post",
	})
}

// GetPostsByUserID retrieves all posts by a specific user
// @Summary Get user posts
// @Description Get all posts by username
// @Tags posts
// @Produce json
// @Param username path string true "Username"
// @Success 200 {array} dto.PostResponseDTO
// @Failure 404 {object} map[string]string "user not found"
// @Failure 500 {object} map[string]string "server error"
// @Router /api/posts/{username} [get]
func (h *Handlers) GetPostsByUserID(c *fiber.Ctx) error {
	username := c.Params("username")

	// Get user by username to find their ID
	user, err := h.userService.GetByUsername(username)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	posts, err := h.postService.GetPostsByUserID(context.Background(), user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to get posts",
		})
	}
	return c.Status(fiber.StatusOK).JSON(dto.ToPostsDTO(posts))
}

// GetRecentPosts retrieves recent posts from all users
// @Summary Get recent posts
// @Description Get most recent posts with optional limit
// @Tags posts
// @Produce json
// @Param limit query int false "Limit (default 50, max 100)"
// @Success 200 {array} dto.PostResponseDTO
// @Failure 500 {object} map[string]string "server error"
// @Router /api/posts/recent [get]
func (h *Handlers) GetRecentPosts(c *fiber.Ctx) error {
	limit := 50

	if queryLimit := c.QueryInt("limit", 50); queryLimit > 0 {
		limit = queryLimit
	}

	posts, err := h.postService.GetRecentPosts(context.Background(), limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to get recent posts",
		})
	}

	fmt.Printf("Found %d posts\n", len(posts))
	for i, p := range posts {
		fmt.Printf("Post %d: ID=%s, Owner.Username=%s, Owner.ID=%s\n",
			i, p.ID, p.Owner.Username, p.Owner.ID)
	}

	dtos := dto.ToPostsDTO(posts)

	fmt.Printf("Sending %d DTOs\n", len(dtos))
	if len(dtos) > 0 {
		fmt.Printf("First DTO: %+v\n", dtos[0])
	}

	return c.Status(fiber.StatusOK).JSON(dtos)
}
