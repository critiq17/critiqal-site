package handlers

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const (
	authHeader = "Authorization"
	username   = "username"
	jwtSecret  = "super_secret_key_123"
)

func (h *Handlers) UserIdentity(c *fiber.Ctx) error {
	header := c.Get(authHeader)

	if header == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "missing auth header",
		})
	}

	parts := strings.Split(header, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid auth header format",
		})
	}

	tokenStr := parts[1]

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "invalid token signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid or expired token",
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid token claims",
		})
	}

	username, ok := claims["username"]
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "missing username in token",
		})
	}

	if exp, ok := claims["exp"].(float64); ok && int64(exp) < time.Now().Unix() {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "token expired",
		})
	}

	c.Locals("username", username)

	return c.Next()
}
