package router

import (
	"github.com/LouisFernando1204/yuccAI-backend/controller"
	"github.com/gofiber/fiber/v2"
)

// configures the routes for the application
func SetUp(app *fiber.App) {

	// route to get all information
	app.Get("/api/get_all_information", controllers.GetAllInformation)
	
	// route to add new information
	app.Post("/api/add_information", controllers.AddInformation)
}