package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/stori_back/controllers"
)

func SetRecipientRoutes(app *fiber.App) {
	app.Post("/recipient", controllers.CreateRecipients)
	app.Get("/recipient/:newsletter_id", controllers.GetRecipients)
	app.Delete("/recipient/:newsletter_id/:recipient_id", controllers.DeleteRecipient)
}
