package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/stori_back/controllers"
)

func SetUpUserRoutes(router fiber.Router) {
	router.Post("/login", controllers.Login)
	router.Post("/signup", controllers.Signup)
	router.Get("/verifyAccount", controllers.VerifyAccount)
	router.Get("/verifyToken", controllers.VerifyToken)
}
