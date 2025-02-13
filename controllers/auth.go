package controllers

import (
	"context"

	"server/database"
	"server/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func RegisterUser(c *fiber.Ctx) error {
	collection := database.GetCollection("createdUser")

	var user models.User

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Inavalid Requets"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	_, err = collection.InsertOne(ctx, bson.M{
		"name":   user.Name,
		"email":  user.Email,
		"age":    user.Age,
		"gender": user.Gender,
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": ""})
	}

	return c.Status(201).JSON(fiber.Map{"message": "User Created Successfully", "user": user})
}
