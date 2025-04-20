package handlers

import "github.com/gofiber/fiber/v2"

//	@Summary		Get all notifications
//	@Description	Get all notifications for a user
//	@Tags			notifications
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		string	true	"User ID"
//	@Success		200		{object}	models.GetAllNotificationsResponse
//	@Failure		400		{object}	models.ErrorBadRequestResponse
//	@Failure		404		{object}	models.ErrorNotFoundResponse
//	@Failure		500		{object}	models.ErrorInternalServerErrorResponse
//	@Router			/notifications/{userId} [get]
func GetAllNotifications(c *fiber.Ctx) error {
	return c.SendString("GetAllNotifications endpoint - To be implemented")
}

//	@Summary		Get notification by ID
//	@Description	Get a notification by its ID
//	@Tags			notifications
//	@Accept			json
//	@Produce		json			
//	@Param			userId			path		string	true	"User ID"
//	@Param			notificationId	path		string	true	"Notification ID"
//	@Success		200				{object}	models.GetNotificationsByIdResponse
//	@Failure		400				{object}	models.ErrorBadRequestResponse
//	@Failure		404				{object}	models.ErrorNotFoundResponse
//	@Failure		500				{object}	models.ErrorInternalServerErrorResponse
//	@Router			/notifications/{userId}/{notificationId} [get]
func GetNotificationById(c *fiber.Ctx) error {
	return c.SendString("GetNotificationById endpoint - To be implemented")
}

//	@Summary		Create a notification
//	@Description	Create a new notification
//	@Tags			Notifications
//	@Accept			json
//	@Produce		json
//	@Param			userId			path		string								true	"User ID"
//	@Param			notification	body		models.CreateNotificationRequest	true	"Notification"
//	@Success		201				{object}	models.CreateResponse
//	@Failure		400				{object}	models.ErrorBadRequestResponse
//	@Failure		404				{object}	models.ErrorNotFoundResponse
//	@Failure		500				{object}	models.ErrorInternalServerErrorResponse
//	@Router			/notifications/{userId} [post]
func CreateNotification(c *fiber.Ctx) error {
	return c.SendString("CreateNotification endpoint - To be implemented")
}


