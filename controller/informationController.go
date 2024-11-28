package controllers

import (
	"context"
	"time"

	"github.com/LouisFernando1204/yuccAI-backend/database"
	"github.com/LouisFernando1204/yuccAI-backend/error"
	"github.com/LouisFernando1204/yuccAI-backend/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// logic to add new information
func AddInformation(c *fiber.Ctx) error {

	// create a context with a 5-second timeout and ensure it's cancelled after use
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// parse the request body to 'information' struct
	var information model.Information
	err := c.BodyParser(&information)
	if err != nil {
		return errors.GetError(c, "Error while parsing data.")
	}

	// try to insert new document into database
	collection := database.GetDatabase().Collection("information")
	result, err := collection.InsertOne(ctx, information)
	if err != nil {
		return errors.GetError(c, "Error while add new information.")
	}

	// it will return JSON response with a status code of 200 (OK) and the result data
	return c.Status(fiber.StatusOK).JSON(result)
}

// logic to fetch all information
func GetAllInformation(c *fiber.Ctx) error {

	// create a context with a 5-second timeout and ensure it's cancelled after use
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var informations []model.Information

	// query to fetch all documents from the collection
	collection := database.GetDatabase().Collection("information")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return errors.GetError(c, "Error while fetch informations.")
	}

	// ensure the cursor is closed once the operation is finished
	defer cursor.Close(ctx)

	// decode the results into the 'informations' slice
	err = cursor.All(ctx, &informations)
	if err != nil {
		return errors.GetError(c, "Error while decoding data.")
	}

	// it will return JSON response with a status code of 200 (OK) and the result data
	return c.Status(fiber.StatusOK).JSON(informations)
}