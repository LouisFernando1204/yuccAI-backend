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

func AddInformation(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var information model.Information
	err := c.BodyParser(&information)
	if err != nil {
		return errors.GetError(c, "Error while parsing data.")
	}

	collection := database.GetDatabase().Collection("information")
	result, err := collection.InsertOne(ctx, information)
	if err != nil {
		return errors.GetError(c, "Error while add new information.")
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func GetAllInformation(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var informations []model.Information

	collection := database.GetDatabase().Collection("information")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return errors.GetError(c, "Error while fetch informations.")
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &informations)
	if err != nil {
		return errors.GetError(c, "Error while decoding data.")
	}

	return c.Status(fiber.StatusOK).JSON(informations)
}