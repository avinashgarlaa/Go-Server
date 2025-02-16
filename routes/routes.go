package routes

import (
	"server/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	api.Post("/register", controllers.RegisterUser)

	api.Get("/users", controllers.GetUsers)

	api.Get("/user/:id", controllers.GetUserById)

}
