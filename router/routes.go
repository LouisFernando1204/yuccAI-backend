package router

import (
	"github.com/LouisFernando1204/yuccAI-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	app.Get("/api/get_all_information", controllers.GetAllInformation)
	app.Post("/api/add_information", controllers.AddInformation)
}