package routes

import (
	"gloo-api/handlers"

	"github.com/gofiber/fiber/v2"
)


func SetupLikesRoutes(router fiber.Router) {
	likes := router.Group("/likes/:userId/:recipeId")
	likes.Post("/", handlers.CreateLike)
	likes.Delete("/", handlers.DeleteLike)
}


