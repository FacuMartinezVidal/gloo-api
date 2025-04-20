package routes

import (
	"gloo-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupMultipleRoutes(router fiber.Router) {
	multiple := router.Group("/multiple/:userId/:recipeId")
	multiple.Get("/", handlers.GetMultiple)
	multiple.Post("/", handlers.CreateMultiple)
	multiple.Put("/", handlers.UpdateMultiple)
}
