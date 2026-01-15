package handlers

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/critiq17/critiqal-site/internal/api/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const (
	accessTokenCookieName  = "access_token"
	refreshTokenCookieName = "refresh_token"
)

func cookieSecureFromEnv() bool {
	v := strings.ToLower(strings.TrimSpace(os.Getenv("COOKIE_SECURE")))
	return v == "1" || v == "true" || v == "yes"
}

func cookieSameSiteFromEnv() string {
	v := strings.TrimSpace(os.Getenv("COOKIE_SAMESITE"))
	if v == "" {
		return "Lax"
	}
	return v
}

func cookieDomainFromEnv() string {
	return strings.TrimSpace(os.Getenv("COOKIE_DOMAIN"))
}

func (h *Handlers) setAuthCookies(c *fiber.Ctx, accessToken string, accessExpires time.Time, refreshToken string, refreshExpires time.Time) {
	secure := cookieSecureFromEnv()
	sameSite := cookieSameSiteFromEnv()
	domain := cookieDomainFromEnv()

	c.Cookie(&fiber.Cookie{
		Name:     accessTokenCookieName,
		Value:    accessToken,
		Path:     "/",
		Domain:   domain,
		Expires:  accessExpires,
		HTTPOnly: true,
		Secure:   secure,
		SameSite: sameSite,
	})

	c.Cookie(&fiber.Cookie{
		Name:     refreshTokenCookieName,
		Value:    refreshToken,
		Path:     "/",
		Domain:   domain,
		Expires:  refreshExpires,
		HTTPOnly: true,
		Secure:   secure,
		SameSite: sameSite,
	})
}

type SignInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UserID   string `json:"user_id"`
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

	if err := h.userService.CreateUser(u); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "error creating user",
		})
	}

	created, err := h.userService.GetByUsername(input.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "error loading created user",
		})
	}

	accessExpires := time.Now().Add(15 * time.Minute)
	refreshExpires := time.Now().Add(7 * 24 * time.Hour)

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  created.ID,
		"username": created.Username,
		"type":     "access",
		"exp":      accessExpires.Unix(),
	})
	accessTokenStr, err := accessToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to generate token")
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  created.ID,
		"username": created.Username,
		"type":     "refresh",
		"exp":      refreshExpires.Unix(),
	})
	refreshTokenStr, err := refreshToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to generate refresh token")
	}

	h.setAuthCookies(c, accessTokenStr, accessExpires, refreshTokenStr, refreshExpires)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user":          dto.ToUserApi(created),
		"token":         accessTokenStr,
		"refresh_token": refreshTokenStr,
	})
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

	user, ok, err := h.userService.Auth(input.Username, input.Password)
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

	accessExpires := time.Now().Add(15 * time.Minute)
	refreshExpires := time.Now().Add(7 * 24 * time.Hour)

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"type":     "access",
		"exp":      accessExpires.Unix(),
	})

	accessTokenStr, err := accessToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to generate token")
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"type":     "refresh",
		"exp":      refreshExpires.Unix(),
	})
	refreshTokenStr, err := refreshToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to generate refresh token")
	}

	h.setAuthCookies(c, accessTokenStr, accessExpires, refreshTokenStr, refreshExpires)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user":          dto.ToUserApi(user),
		"token":         accessTokenStr,
		"refresh_token": refreshTokenStr,
	})
}

func (h *Handlers) Refresh(c *fiber.Ctx) error {
	refreshTokenStr := c.Cookies(refreshTokenCookieName)
	if refreshTokenStr == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "missing refresh token",
		})
	}

	token, err := jwt.ParseWithClaims(refreshTokenStr, jwt.MapClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "invalid token signing method")
		}
		return []byte(jwtSecret), nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid or expired refresh token",
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid refresh token claims",
		})
	}

	if fmt.Sprint(claims["type"]) != "refresh" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid refresh token type",
		})
	}

	userID := fmt.Sprint(claims["user_id"])
	username := fmt.Sprint(claims["username"])
	if userID == "" || username == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid refresh token payload",
		})
	}

	accessExpires := time.Now().Add(15 * time.Minute)
	newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"type":     "access",
		"exp":      accessExpires.Unix(),
	})

	newAccessTokenStr, err := newAccessToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to generate token")
	}

	secure := cookieSecureFromEnv()
	sameSite := cookieSameSiteFromEnv()
	domain := cookieDomainFromEnv()
	c.Cookie(&fiber.Cookie{
		Name:     accessTokenCookieName,
		Value:    newAccessTokenStr,
		Path:     "/",
		Domain:   domain,
		Expires:  accessExpires,
		HTTPOnly: true,
		Secure:   secure,
		SameSite: sameSite,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": newAccessTokenStr,
	})
}

func (h *Handlers) SignOut(c *fiber.Ctx) error {
	secure := cookieSecureFromEnv()
	sameSite := cookieSameSiteFromEnv()
	domain := cookieDomainFromEnv()

	expired := time.Unix(0, 0)
	c.Cookie(&fiber.Cookie{
		Name:     accessTokenCookieName,
		Value:    "",
		Path:     "/",
		Domain:   domain,
		Expires:  expired,
		MaxAge:   -1,
		HTTPOnly: true,
		Secure:   secure,
		SameSite: sameSite,
	})
	c.Cookie(&fiber.Cookie{
		Name:     refreshTokenCookieName,
		Value:    "",
		Path:     "/",
		Domain:   domain,
		Expires:  expired,
		MaxAge:   -1,
		HTTPOnly: true,
		Secure:   secure,
		SameSite: sameSite,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}

func (h *Handlers) AuthMe(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(string)
	if !ok || userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "missing user context",
		})
	}

	u, err := h.userService.GetByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.ToUserApi(u))
}
