package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/stori_back/controllers"
)

func SetRecipientRoutes(app *fiber.App) {
	app.Get("/recipient", controllers.GetRecipients)
}
