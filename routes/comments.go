package routes

import (
	"gloo-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupCommentsRoutes(router fiber.Router) {
	comments := router.Group("/comments/:recipeId/:commentId")
	comments.Get("/", handlers.GetCommentById)
	comments.Put("/", handlers.UpdateComment)
	comments.Delete("/", handlers.DeleteComment)
	comments.Post("/:userId/:fatherCommentId", handlers.CreateComment)
}
