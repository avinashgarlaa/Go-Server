package controllers

import (
	"context"
	"time"

	"server/database"
	"server/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterUser(c *fiber.Ctx) error {
	collection := database.GetCollection("createdUser")

	var user models.User

	// Parse JSON request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Validate required fields
	if err := models.Validate.Struct(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"validation_error": err.Error()})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, bson.M{
		"name":    user.Name,
		"email":   user.Email,
		"age":     user.Age,
		"gender":  user.Gender,
		"balance": user.Balance,
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	// Extract the inserted ID and convert to string
	insertedID := result.InsertedID.(primitive.ObjectID).Hex()

	return c.JSON(fiber.Map{
		"message": "User Created Successfully",
		"user": fiber.Map{
			"id":      insertedID,
			"name":    user.Name,
			"email":   user.Email,
			"age":     user.Age,
			"gender":  user.Gender,
			"balance": user.Balance,
		},
	})
}

func GetUsers(c *fiber.Ctx) error {
	collection := database.GetCollection("createdUser")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Find all users
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	defer cursor.Close(ctx)

	var users []models.User

	if err := cursor.All(ctx, &users); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"users": users})

}

func GetUserById(c *fiber.Ctx) error {
	collection := database.GetCollection("createdUser")

	id := c.Params("id")

	ObjId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid user ID format"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User

	err = collection.FindOne(ctx, bson.M{"_id": ObjId}).Decode(&user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "cannot find user"})
	}
	return c.JSON(user)
}
