package handlers

import "github.com/gofiber/fiber/v2"

//	@Summary		Create a like
//	@Description	Create a like
//	@Tags			Likes
//	@Accept			json
//	@Produce		json
//	@Param			userId		path		string	true	"User ID"
//	@Param			recipeId	path		string	true	"Recipe ID"
//	@Success		201			{object}	models.CreateResponse
//	@Failure		400			{object}	models.ErrorBadRequestResponse
//	@Failure		404			{object}	models.ErrorNotFoundResponse
//	@Failure		500			{object}	models.ErrorInternalServerErrorResponse
//	@Router			/likes/{userId}/{recipeId} [post]
func CreateLike(c *fiber.Ctx) error {
	return c.SendString("CreateLike endpoint - To be implemented")
}

//	@Summary		Delete a like
//	@Description	Delete a like
//	@Tags			Likes
//	@Accept			json
//	@Produce		json
//	@Param			userId		path		string	true	"User ID"
//	@Param			recipeId	path		string	true	"Recipe ID"
//	@Success		200			{object}	models.DeleteResponse
//	@Failure		400			{object}	models.ErrorBadRequestResponse
//	@Failure		404			{object}	models.ErrorNotFoundResponse
//	@Failure		500			{object}	models.ErrorInternalServerErrorResponse
//	@Router			/likes/{userId}/{recipeId} [delete]
func DeleteLike(c *fiber.Ctx) error {
	return c.SendString("DeleteLike endpoint - To be implemented")
}


