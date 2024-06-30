package handlers

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "QwITravel API v1.0.0 is up and running!",
		"status":  "success",
	})
}

func Endpoints(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"api_base": "https://qwitravel-api.qwit.hu",
		"endpoints": []fiber.Map{
			{
				"path":        "/",
				"methods":     []string{"GET"},
				"description": "API status and version information",
			},
			{
				"path":        "/endpoints",
				"methods":     []string{"GET"},
				"description": "List of all available endpoints",
			},
		},
	})
}
