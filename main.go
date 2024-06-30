package main

import (
	"github.com/gofiber/fiber/v2"
	"qwit.dev/qwitravel-api/handlers"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:      true,
		ServerHeader: "Fiber",
		AppName:      "QwITravel API v1.0.0",
	})

	// base routes
	app.Get("/", handlers.Home)
	app.Get("/endpoints", handlers.Endpoints)
	app.Get("/version", handlers.Version)

	// stop related routes
	stops := app.Group("/stops")
	stops.Get("/schedule", handlers.GetStopSchedule)

	// start the server
	err := app.Listen(":3002")
	if err != nil {
		panic(err)
	}
}
