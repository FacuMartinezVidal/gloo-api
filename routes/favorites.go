package routes

import (
	"gloo-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupFavoritesRoutes(router fiber.Router) {
	favorites := router.Group("/favorites/:userId/:recipeId")
	favorites.Get("/", handlers.GetFavorites)
	favorites.Post("/", handlers.CreateFavorite)
	favorites.Delete("/:favoriteId", handlers.DeleteFavorite)
}
