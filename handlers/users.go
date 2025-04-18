package handlers

import (
	"github.com/gofiber/fiber/v2"
)

//	@Summary		Get user by ID
//	@Description	Get user by ID
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		string	true	"User ID"
//	@Success		200	{object}	models.GetUserByIdResponse	"Successfully retrieved user"
//	@Failure		400	{object}	models.ErrorBadRequestResponse	"Bad Request"
//	@Failure		500	{object}	models.ErrorInternalServerErrorResponse	"Internal Server Error"
//	@Failure		404	{object}	models.ErrorNotFoundResponse	"Not Found"
//	@Router			/users/{userId} [get]
func GetUserById(c *fiber.Ctx) error {
	return c.SendString("GetUserById endpoint - To be implemented")
}	

//	@Summary		Get profile
//	@Description	Get profile
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		string	true	"User ID"
//	@Success		200	{object}	models.GetProfileResponse	"Successfully retrieved profile"
//	@Failure		400	{object}	models.ErrorBadRequestResponse	"Bad Request"
//	@Failure		500	{object}	models.ErrorInternalServerErrorResponse	"Internal Server Error"
//	@Failure		404	{object}	models.ErrorNotFoundResponse	"Not Found"
//	@Router			/users/profile [get]
func GetProfile(c *fiber.Ctx) error {
	return c.SendString("getProfile endpoint - To be implemented")
}

//	@Summary		Get change
//	@Description	Get change
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		string	true	"User ID"
//	@Success		200	{object}	models.GetChangeResponse	"Successfully retrieved change"
//	@Failure		400	{object}	models.ErrorBadRequestResponse	"Bad Request"
//	@Failure		500	{object}	models.ErrorInternalServerErrorResponse	"Internal Server Error"
//	@Failure		404	{object}	models.ErrorNotFoundResponse	"Not Found"
//	@Router			/users/change [get]
func GetChange(c *fiber.Ctx) error {
	return c.SendString("getChange endpoint - To be implemented")
}


//	@Summary		Update user
//	@Description	Update user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		string	true	"User ID"
//	@Param			user	body	models.UpdateUserRequest	true	"User"
//	@Success		200	{object}	models.UpdateResponse	"Successfully updated user"
//	@Failure		400	{object}	models.ErrorBadRequestResponse	"Bad Request"
//	@Failure		500	{object}	models.ErrorInternalServerErrorResponse	"Internal Server Error"
//	@Failure		404	{object}	models.ErrorNotFoundResponse	"Not Found"
//	@Router			/users [put]
func UpdateUser(c *fiber.Ctx) error {
	return c.SendString("UpdateUser endpoint - To be implemented")
}



