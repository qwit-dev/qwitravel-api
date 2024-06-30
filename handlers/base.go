package handlers

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "QwITravel API v1.0.0 is up and running!",
		"status":  "success",
	})
}
