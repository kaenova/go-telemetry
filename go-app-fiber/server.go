package main

import (
	"github.com/gofiber/fiber/v2"
	"go.elastic.co/apm/module/apmfiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(apmfiber.Middleware())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3000")
}
