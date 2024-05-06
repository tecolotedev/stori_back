package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/stori_back/controllers"
)

func SetNewsletterVersionRoutes(app *fiber.App) {
	app.Get("/newsletter-version/:newsletter_id", controllers.GetAllNewslettersVersions)
	app.Post("/newsletter-version/:newsletter_id", controllers.CreateNewsletterVersion)
	app.Put("/newsletter-version/:newsletter_version_id", controllers.UpdateNewsletterVersion)
	app.Delete("/newsletter-version/:newsletter_version_id", controllers.DeleteNewsletterVersion)

	app.Post("/send-newsletter/:newsletter_version_id", controllers.SendNewsletter)

}
