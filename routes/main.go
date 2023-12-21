package routes

import "github.com/gofiber/fiber/v2"

func SetUpRoutes(app *fiber.App) {
	routerUser := app.Group("/api")
	SetUpUserRoutes(routerUser)
}
