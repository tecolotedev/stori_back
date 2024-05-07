package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/stori_back/config"
	"github.com/tecolotedev/stori_back/email"
	"github.com/tecolotedev/stori_back/models"
	"github.com/tecolotedev/stori_back/routes"
)

func main() {
	// init env vars
	config.InitConfig()

	// init db
	models.InitDB()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"ok": true, "message": "API is working"})
	})

	routes.SetRoutes(app)

	// concurrency for emails
	email.EmailHandler.NewsletterEmailChan = make(chan email.NewsletterEmail)
	email.EmailHandler.DoneChan = make(chan bool)
	email.EmailHandler.InitDialer()
	go email.EmailHandler.ListenEmails()

	log.Fatal(app.Listen(":8000"))
}
