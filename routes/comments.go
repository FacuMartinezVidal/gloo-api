package routes

import (
	"gloo-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupCommentsRoutes(router fiber.Router) {
	comments := router.Group("/comments/:recipeId/")
	comments.Get("/:commentId", handlers.GetCommentById)
	comments.Put("/:commentId", handlers.UpdateComment)
	comments.Delete("/:commentId", handlers.DeleteComment)
	comments.Post("/:userId/:fatherCommentId", handlers.CreateComment)
	comments.Get("/:commentId/admin/:organizationId", handlers.GetAllCommentsAdmin)
	comments.Put("/:commentId/admin/:organizationId", handlers.UpdateCommentAdmin)
}
