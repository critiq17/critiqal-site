package handlers

import (
	"fmt"

	"github.com/critiq17/critiqal-site/internal/api/dto"
	"github.com/gofiber/fiber/v2"
)

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"requireds"`
}

func (h *Handlers) SignUp(c *fiber.Ctx) error {

	var input dto.UserApi

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("%v", err),
		})
	}

	u := dto.ToDBModel(&input)

	if err := h.service.CreateUser(u); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "error creating user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(input)
}

func (h *Handlers) SignIn(c *fiber.Ctx) error {

	var input SignInInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	user, ok, err := h.service.Auth(input.Username, input.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to authenticate users",
		})
	}

	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid username or password",
		})
	}

	safeUser := &dto.UserApi{
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	return c.Status(fiber.StatusOK).JSON(safeUser)
}
