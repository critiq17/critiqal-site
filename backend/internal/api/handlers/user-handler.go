package handlers

import (
	"fmt"

	"github.com/critiq17/critiqal-site/internal/domain/user/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) AddUser(c *fiber.Ctx) error {
	var req dto.CreateRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid json body",
		})
	}

	u := dto.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	if err := h.userRepo.Create(&u); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(u)
}

func (h *Handlers) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	err := h.userRepo.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("failed to delete user: %v", err),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)

}

func (h *Handlers) GetUsers() *[]dto.UserApi {
	return nil
}
