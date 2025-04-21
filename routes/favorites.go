package routes

import (
	"gloo-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupFavoritesRoutes(router fiber.Router) {
	favorites := router.Group("/favorites/:userId/:collectionId")
	favorites.Get("/", handlers.GetFavorites)
	favorites.Post("/:recipeId", handlers.CreateFavorite)
	favorites.Put("/:favoriteId", handlers.UpdateFavorite)
	favorites.Delete("favoriteId", handlers.DeleteFavorite)
}
