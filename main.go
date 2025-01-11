package main

import (
	"log"

	"github.com/LouisFernando1204/yuccAI-backend/database"
	"github.com/LouisFernando1204/yuccAI-backend/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	// make a new instance of fiber
	app := fiber.New();

	// connect to database
	database.ConnectDatabase()

	// after the `main` function finished executing, it will start to disconnect the database
	defer database.DisconnectDatabase()

	// // handle CORS error
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: "https://0bf8-139-0-80-240.ngrok-free.app/",
	// 	AllowMethods: "*",
	// 	AllowHeaders: "*",
	// 	AllowCredentials: true,
	// }))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST",
		AllowHeaders: "*",
	}))

	// handle set up router
	router.SetUp(app)

	// run on local server
	log.Fatal(app.Listen(":3000"));
}