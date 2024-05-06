package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/stori_back/controllers"
)

func SetNewsletterRoutes(app *fiber.App) {
	app.Get("/newsletter", controllers.GetAllNewsletters)
	app.Post("/newsletter", controllers.CreateNewsletter)
	app.Put("/newsletter/:newsletter_id", controllers.UpdateNewsletter)
	app.Delete("/newsletter/:newsletter_id", controllers.DeleteNewsletter)
}
