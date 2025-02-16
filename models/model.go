package models

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name    string             `json:"name" validate:"required"`
	Email   string             `json:"email" validate:"required,email"`
	Age     int                `json:"age"`
	Gender  string             `json:"gender"`
	Balance float64            `json:"balance"`
}

// Create a validator instance
var Validate = validator.New()
