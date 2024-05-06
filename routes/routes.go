package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	SetNewsletterRoutes(app)
	SetNewsletterVersionRoutes(app)
	SetRecipientRoutes(app)
}
