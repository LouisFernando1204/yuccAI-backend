package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func GetDatabase() *mongo.Database {
	if Client == nil {
		log.Fatalf("MongoDB is not initialized.")
	}
	return Client.Database("yuccAI_db")
}

func ConnectDatabase() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	MONGO_URL := os.Getenv("MONGO_URL")

	clientOption := options.Client().ApplyURI(MONGO_URL)
	result, err := mongo.Connect(context.Background(), clientOption) 
	if err != nil {
		log.Fatalf("Error while connecting to MongoDB: %s", err)
	}
	Client = result

	err = Client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error while ping to MongoDB: %s", err)
	}

	log.Printf("Successfully connected to MongoDB!")
}

func DisconnectDatabase() {
	if Client == nil {
		log.Fatalf("Mongo DB is not initialized")
		return
	}

	err := Client.Disconnect(context.Background())
	if err != nil {
		log.Fatalf("Error while disconnecting MongoDB: %s", err)
	}

	log.Printf("Successfully disconnected from MongoDB!")
}