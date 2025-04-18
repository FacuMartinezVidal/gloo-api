package routes

import (
	"gloo-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupSearchRoutes(app *fiber.App) {
	search := app.Group("/search/:userId")
	search.Get("/", handlers.GetAllSearches)
	search.Get("/suggestions", handlers.GetAllSuggestions)
}
