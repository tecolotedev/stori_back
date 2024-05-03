package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/stori_back/routes"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"ok": true, "message": "API is working"})
	})

	routes.SetRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
