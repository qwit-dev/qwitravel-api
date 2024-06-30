package main

import (
	"github.com/gofiber/fiber/v2"
	"qwit.dev/qwitravel-api/handlers"
)

func main() {
	app := fiber.New()

	// base routes
	app.Get("/", handlers.Home)
	app.Get("/endpoints", handlers.Home)

	err := app.Listen(":3002")
	if err != nil {
		panic(err)
	}
}