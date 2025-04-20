package routes

import (
	"gloo-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupNotificationsRoutes(router fiber.Router) {
	notifications := router.Group("/notifications/:userId")
	notifications.Get("/", handlers.GetAllNotifications)
	notifications.Get("/:notificationId", handlers.GetNotificationById)
	notifications.Post("/", handlers.CreateNotification)
}