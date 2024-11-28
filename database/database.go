package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// declare as a global variable, so it can accessed by all function
var Client *mongo.Client

// logic to return database on MongoDB
func GetDatabase() *mongo.Database {

	// if Client has an empty value, it will give this output
	if Client == nil {
		log.Fatalf("MongoDB is not initialized.")
	}

	// if Client doesn't have an empty value, it will return the database
	return Client.Database("yuccAI_db")
}

// logic to connect the database using MONGO_URL from the .env file
func ConnectDatabase() {

	// try load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// load MONGO_URL from .env file
	MONGO_URI := os.Getenv("MONGO_URI")

	// configure MongoDB with MONGO_URI from .env file
	clientOption := options.Client().ApplyURI(MONGO_URI)
	
	// try to connect to MongoDB
	// contextBackground() is a default context that doesn't have a duration or deadline for its process and cannot be canceled
	result, err := mongo.Connect(context.Background(), clientOption) 
	if err != nil {
		log.Fatalf("Error while connecting to MongoDB: %s", err)
	}

	// save result to global variable (Client)
	Client = result

	// ping MongoDB
	err = Client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error while ping to MongoDB: %s", err)
	}

	// it will display this output if the connection to MongoDB is successful
	log.Printf("Successfully connected to MongoDB!")
}

func DisconnectDatabase() {

	// if Client has an empty value, it will give this output
	if Client == nil {
		log.Fatalf("Mongo DB is not initialized.")
		return
	}

	// try to disconenct from MongoDB
	err := Client.Disconnect(context.Background())
	if err != nil {
		log.Fatalf("Error while disconnecting MongoDB: %s", err)
	}

	// it will display this output if the disconnection process from MongoDB is successful
	log.Printf("Successfully disconnected from MongoDB!")
}