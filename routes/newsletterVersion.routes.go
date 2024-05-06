package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/stori_back/controllers"
)

func SetNewsletterVersionRoutes(app *fiber.App) {
	// app.Get("/newsletter-version", controllers.GetAllNewsletters)
	app.Post("/newsletter-version/:newsletter_id", controllers.CreateNewsletterVersion)
}
