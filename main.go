package main

import (
	"github.com/gofiber/fiber/v2"
	"qwit.dev/qwitravel-api/handlers"
)

func main() {
	app := fiber.New()

	// base routes
	app.Get("/", handlers.Home)
	app.Get("/endpoints", handlers.Endpoints)
	app.Get("/version", handlers.Version)

	// start the server
	err := app.Listen(":3002")
	if err != nil {
		panic(err)
	}
}
