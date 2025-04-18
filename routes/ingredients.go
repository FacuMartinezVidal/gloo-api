package routes

import (
	"gloo-api/handlers"

	"github.com/gofiber/fiber/v2"
)


func SetupIngredientsRoutes(router fiber.Router) {
	ingredients := router.Group("/ingredients/:recipeId/:ingredientId")
	ingredients.Put("/", handlers.UpdateIngredient)
	ingredients.Delete("/", handlers.DeleteIngredient)
	ingredients.Post("/", handlers.CreateIngredient)
}


