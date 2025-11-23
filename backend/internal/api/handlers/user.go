package handlers

import (
	"fmt"

	"github.com/critiq17/critiqal-site/internal/api/dto"
	"github.com/critiq17/critiqal-site/internal/domain/user"
	"github.com/gofiber/fiber/v2"
)

// AddUser godoc
// @Summary      Created User
// @Description  Created new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user            body      dto.UserApi  true  "User data"
// @Success      201  {object}   map[string]interface{}    "user created successfuly"
// @Failure      400  {object}   map[string]string         "bad request"
// @Failure      500  {object}   map[string]string         "internal server error"
// @Router       /api/users [post]
func (h *Handlers) AddUser(c *fiber.Ctx) error {
	var req dto.CreateRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid json body",
		})
	}

	u := user.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	if err := h.userService.CreateUser(&u); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(u)
}

// DeleteUser godoc
// @Summary      Delete user
// @Description  Deletes a user by ID
// @Tags         users
// @Param        id   path      string  true  "User ID"
// @Success      204  "No Content"
// @Failure      400  {object}  map[string]string  "Invalid ID"
// @Failure      500  {object}  map[string]string  "Server error"
// @Router       /users/{id} [delete]
func (h *Handlers) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	err := h.userService.DeleteUser(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("failed to delete user: %v", err),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)

}

// GetUsers godoc
// @Summary 		Get all users
// @Description Return all users from db
// @Tags users
// @Success 200 "StatusOK"
// @Failure  500 {object} map[string]string "Server error"
// @Router /users [get]
func (h *Handlers) GetUsers(c *fiber.Ctx) error {

	users, err := h.userService.GetUsers()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "error get users",
		})
	}

	dtos := dto.ToUsersApi(users)

	return c.JSON(dtos)
}

func (h *Handlers) GetByUsername(c *fiber.Ctx) error {

	username := c.Params("username")

	user, err := h.userService.GetByUsername(username)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.ToUserApi(user))
}

func (h *Handlers) Get(c *fiber.Ctx) error {

	id := c.Params("id")

	user, err := h.userService.GetByID(id)

	userApi := dto.ToUserApi(user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error found user by id",
		})
	}

	return c.Status(fiber.StatusOK).JSON(userApi)
}

func (h *Handlers) UploadPhoto(c *fiber.Ctx) error {
	username := c.Params("username")
	if username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid username",
		})
	}

	file, err := c.FormFile("photo")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid file",
		})
	}

	url, err := h.userService.UploadUserPhoto(c.Context(), username, file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "error uploading photo",
		})
	}

	return c.JSON(fiber.Map{"url": url})
}

func (h *Handlers) SearchUsers(c *fiber.Ctx) error {
	username := c.Params("username")
	if username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid username",
		})
	}

	users, err := h.userService.SearchUsers(username)
	if err != nil {
		return err
	}

	response := dto.ToUsersApi(users)

	return c.JSON(response)

}

func (h *Handlers) GetMe(c *fiber.Ctx) error {
	username := c.Locals("username").(string)

	user, err := h.userService.GetByUsername(username)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	return c.JSON(dto.ToUserApi(user))
}
