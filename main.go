package main

import (
	"fmt"
	"log"
	"server/database"
	"server/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.ConnectMongo()

	app := fiber.New()

	routes.SetupRoutes(app)

	port := ":3030"
	fmt.Println("Server running on http://localhost" + port)
	log.Fatal(app.Listen(port))
}
