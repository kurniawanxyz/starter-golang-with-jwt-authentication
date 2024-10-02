package helper

import "github.com/gofiber/fiber/v2"

func HandleResponse(c *fiber.Ctx, status int, message string, data any) error {

	var success bool = false

	if status >= 200 && status <= 299 {
		success = true
	}

	return c.Status(status).JSON(fiber.Map{
		"success": success,
		"message": message,
		"data":    data,
	})
}