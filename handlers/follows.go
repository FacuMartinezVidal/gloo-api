package handlers

import "github.com/gofiber/fiber/v2"

//	@Summary		Create a follow
//	@Description	Create a new follow
//	@Tags			Follows
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		string						true	"User ID"
//	@Param			follow	body		models.CreateFollowRequest	true	"Follow"
//	@Success		201		{object}	models.CreateResponse
//	@Failure		400		{object}	models.ErrorBadRequestResponse
//	@Failure		404		{object}	models.ErrorNotFoundResponse
//	@Failure		500		{object}	models.ErrorInternalServerErrorResponse
//	@Router			/follows/{userId} [post]
func CreateFollow(c *fiber.Ctx) error {
	return c.SendString("CreateFollow endpoint - To be implemented")
}

//	@Summary		Delete a follow
//	@Description	Delete a follow
//	@Tags			Follows
//	@Accept			json
//	@Produce		json
//	@Param			userId		path		string	true	"User ID"
//	@Param			followId	path		string	true	"Follow ID"	
//	@Success		200			{object}	models.DeleteResponse
//	@Failure		400			{object}	models.ErrorBadRequestResponse
//	@Failure		404			{object}	models.ErrorNotFoundResponse
//	@Failure		500			{object}	models.ErrorInternalServerErrorResponse
//	@Router			/follows/{userId}/{followId} [delete]
func DeleteFollow(c *fiber.Ctx) error {
	return c.SendString("DeleteFollow endpoint - To be implemented")
}

