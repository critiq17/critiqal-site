package handlers

import (
	"fmt"

	"github.com/critiq17/critiqal-site/internal/domain/user/dto"
	"github.com/gofiber/fiber/v2"
)

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"username" binding:"requireds"`
}

func (h *Handlers) SignUp(c *fiber.Ctx) error {

	var input dto.UserApi

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("%v", err),
		})
	}

	u := dto.User{
		Username:  input.Username,
		Email:     input.Email,
		Password:  input.Password,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	if err := h.userRepo.Create(&u); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(u)
}
