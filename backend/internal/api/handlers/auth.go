package handlers

import (
	"fmt"
	"time"

	"github.com/critiq17/critiqal-site/internal/api/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type SignInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Sign-up godoc
// @Summary      Register
// @Description  Register new user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user            body      dto.UserApi  true  "User data"
// @Success      201  {object}   map[string]interface{}    "user created successfuly"
// @Failure      400  {object}   map[string]string         "bad request"
// @Failure      500  {object}   map[string]string         "internal server error"
// @Router       /auth/sign-up [post]
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

// Sign-in godoc:
// @Summary      Auth
// @Description  Auth if user created
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user            body      SignInInput  true  "User data"
// @Success      201  {object}   map[string]interface{}    "user auth successfuly"
// @Failure      400  {object}   map[string]string         "bad request"
// @Failure      500  {object}   map[string]string         "internal server error"
// @Router       /auth/sign-in [post]
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": input.Username,
		"exp":      time.Now().Add(25 * time.Hour).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to generate token")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": tokenStr, "user": user})
}
