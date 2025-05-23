package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Uploading to render...")
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal("Error in listerning to port")
	}
}
