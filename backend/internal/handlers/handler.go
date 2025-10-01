package handlers

import "github.com/gofiber/fiber/v2"

type RequestBody struct {
	Message string `json:"message"`
}

func HandleMessage(c *fiber.Ctx) error {
	var body RequestBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid JSON",
		})
	}

	switch body.Message {
	case "hello":
		return c.JSON(fiber.Map{
			"reply": "Hi there!",
		})
	case "bye":
		return c.JSON(fiber.Map{
			"reply": "Goodbye!",
		})
	default:
		return c.JSON(fiber.Map{
			"reply": "I don't understand.",
		})
	}
}
