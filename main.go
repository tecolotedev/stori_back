package main

import (
	"log"
	"time"

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
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		cookie := new(fiber.Cookie)
		cookie.Name = "john"
		cookie.Value = "doe"
		cookie.Expires = time.Now().Add(24 * time.Hour)

		// Set cookie
		c.Cookie(cookie)
		return c.JSON(fiber.Map{"ok": true, "message": "api is working"})

	})

	routes.SetUpRoutes(app)

	log.Fatal((app.Listen(":" + config.EnvVars.PORT)))

}
