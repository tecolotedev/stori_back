package main

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tecolotedev/stori_back/config"
	"github.com/tecolotedev/stori_back/db"
	"github.com/tecolotedev/stori_back/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.SetUpConfig()
	db.InitDb()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, https://stori-front.tecolotedev.com",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		ExposeHeaders: "Set-Cookie",
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"ok": true, "message": "api is working"})
	})

	routes.SetUpRoutes(app)

	log.Fatal((app.Listen(":" + config.EnvVars.PORT)))

}
