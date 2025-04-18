package routes

import (
	"gloo-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRecipesRoutes(router fiber.Router) {
	recipes := router.Group("/recipes")
	recipes.Get("/", handlers.GetAllRecipe)
	recipes.Get("/:id", handlers.GetRecipeById)
	recipes.Put("/:id", handlers.UpdateRecipe)
	recipes.Post("/:userId", handlers.CreateRecipe)
	recipes.Delete("/:id", handlers.DeleteRecipe)
	recipes.Post("/admin/:recipeId/:adminId", handlers.UpdateRecipeAdmin)
	recipes.Get("/admin/:organizationId", handlers.GetAllRecipeAdmin)
}
