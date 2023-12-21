package main

import (
	"log"

	"github.com/tecolotedev/stori_back/config"
	"github.com/tecolotedev/stori_back/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.SetUpConfig()
	db.InitDb()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")

	})

	log.Fatal((app.Listen(":" + config.EnvVars.PORT)))

}
