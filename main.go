package main

import (
	"log"

	"github.com/LouisFernando1204/yuccAI-backend/database"
	"github.com/LouisFernando1204/yuccAI-backend/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New();

	database.ConnectDatabase()
	defer database.DisconnectDatabase()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.SetUp(app)

	log.Fatal(app.Listen(":4050"));
}