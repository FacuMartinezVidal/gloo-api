package routes

import (
	"gloo-api/handlers"

	"github.com/gofiber/fiber/v2"
)


func SetupUsersRoutes(app *fiber.App) {
	api := app.Group("/api")
	users := api.Group("/users/:userId")
	users.Get("/", handlers.GetUserById)
	users.Get("/profile", handlers.GetProfile)
	users.Get("/change", handlers.GetChange)
	users.Put("/", handlers.UpdateUser)
}


