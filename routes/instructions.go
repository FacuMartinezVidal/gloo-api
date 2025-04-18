package routes

import (
	"gloo-api/handlers"

	"github.com/gofiber/fiber/v2"
)	


func SetupInstructionsRoutes(router fiber.Router) {
	instructions := router.Group("/instructions/:recipeId/:instructionId")
	instructions.Get("/", handlers.GetInstructionById)
	instructions.Put("/", handlers.UpdateInstruction)
	instructions.Delete("/", handlers.DeleteInstruction)
	instructions.Post("/", handlers.CreateInstruction)
}

	
