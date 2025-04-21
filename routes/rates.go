package routes

import (
	"gloo-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRatesRoutes(app *fiber.App) {
	rates := app.Group("/rates/:recipeId/:userId")
	rates.Post("/", handlers.CreateRate)
	rates.Get("/", handlers.GetRates)
	rates.Put("/:rateId", handlers.UpdateRate)
	rates.Delete("/:rateId", handlers.DeleteRate)
}
