package routes

import (
	"gloo-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupFollowsRoutes(router fiber.Router) {
	follows := router.Group("/follows/:userId")
	follows.Post("/", handlers.CreateFollow)
	follows.Delete("/:followId", handlers.DeleteFollow)
}
