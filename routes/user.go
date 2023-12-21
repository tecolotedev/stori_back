package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/stori_back/controllers"
)

func SetUpUserRoutes(router fiber.Router) {
	router.Post("/user", controllers.CreateUser)
}
