package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func GetCollection(collectionName string) *mongo.Collection {
	return Client.Database("UserCreation").Collection(collectionName)
}

func ConnectMongo() {
	mongoURI := "mongodb://localhost:27017/?directConnection=true"

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("Error connecting to MongoDB : ", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Ping Error: ", err)
	}

	fmt.Println("Connected to MongoDB! ")

	Client = client
}
