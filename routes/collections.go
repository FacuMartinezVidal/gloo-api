package routes

import (
	"gloo-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupCollectionsRoutes(router fiber.Router) {
	
	collections := router.Group("/collections/:userId")
	collections.Get("/", handlers.GetAllCollections)
	collections.Get("/:collectionId", handlers.GetCollectionById)
	collections.Post("/", handlers.CreateCollection)
}
