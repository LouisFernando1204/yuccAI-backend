package errors

import "github.com/gofiber/fiber/v2"

// reusable function to show error
func GetError(c *fiber.Ctx, err string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
}