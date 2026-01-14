package handlers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const (
	authHeader = "Authorization"
	jwtSecret  = "super_secret_key_123"
)

// UserIdentity extracts user info from JWT session
func (h *Handlers) UserIdentity(c *fiber.Ctx) error {
	tokenStr := ""

	header := c.Get(authHeader)
	if header != "" {
		parts := strings.Split(header, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid auth header format",
			})
		}
		tokenStr = parts[1]
	} else {
		tokenStr = c.Cookies("access_token")
	}

	if tokenStr == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "missing auth token",
		})
	}

	token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "invalid token signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid or expired token: " + err.Error(),
		})
	}

	claims := token.Claims.(jwt.MapClaims)

	userID, ok := claims["user_id"]
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "missing user_id in token",
		})
	}

	username, ok := claims["username"]
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "missing username in token",
		})
	}

	// Store both in context
	c.Locals("user_id", fmt.Sprint(userID))
	c.Locals("username", fmt.Sprint(username))

	return c.Next()
}
