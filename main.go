package main

import (
	"log"

	"github.com/tecolotedev/stori_back/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.SetUpConfig()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal((app.Listen(":" + config.EnvVars.PORT)))

}
