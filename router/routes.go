package router

import (
	"github.com/LouisFernando1204/yuccAI-backend/controller"
	"github.com/gofiber/fiber/v2"
)

// configures the routes for the application
func SetUp(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Successful ping!",
		})
	})

	// route to get all information
	app.Get("/api/get_all_information", controllers.GetAllInformation)
	
	// route to add new information
	app.Post("/api/add_information", controllers.AddInformation)
}