package routes

import (
	"server/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	api.Post("/register", controllers.RegisterUser)

	api.Get("/welcome", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the server")
	})

	api.Get("/message", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "success",
			"name":    "Hello Avinash",
		})
	})

	api.Get("/user/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.SendString("Name Is " + name)
	})

}
