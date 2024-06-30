package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"qwit.dev/qwitravel-api/clients"
)

func GetStopSchedule(c *fiber.Ctx) error {
	var provider string = c.Query("provider", "")
	var stopIdsQuery string = c.Query("stop_ids", "")
	var stopIds []string = strings.Split(stopIdsQuery, ",")
	var date string = c.Query("date", "")

	if stopIdsQuery == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Missing or invalid \"stop_ids\" parameter!",
		})
	}
	if provider == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Missing or invalid \"provider\" parameter!",
		})
	}

	if strings.ToLower(provider) == "bkk" {
		// get stop schedule from BKK
		status, bkkStopSchedule := clients.BKKGetStopSchedule(stopIds, strings.ReplaceAll(date, "-", ""))
		if status != fiber.StatusOK {
			return c.Status(status).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed to get stop schedule from BKK API!",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "success",
			"data":   bkkStopSchedule,
		})
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"status":  "error",
		"message": "Invalid or unsupported \"provider\" parameter!",
	})
}
