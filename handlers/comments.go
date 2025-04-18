package handlers

import (
	"github.com/gofiber/fiber/v2"
)

//	@Summary		Get comment by ID
//	@Description	Get a specific comment by its ID
//	@Tags			Comments
//	@Accept			json
//	@Produce		json
//	@Param			recipeId	path		int										true	"Recipe ID"	minimum(1)
//	@Param			commentId	path		int										true	"Comment ID"	minimum(1
//	@Success		200	{object}	models.GetCommentByIdResponse	"Successfully retrieved comment"
//	@Failure		400	{object}	models.ErrorBadRequestResponse	"Bad Request"
//	@Failure		404	{object}	models.ErrorNotFoundResponse	"Comment not found"
//	@Failure		500	{object}	models.ErrorInternalServerErrorResponse	"Internal Server Error"
//	@Router			/comments/{recipeId}/{commentId} [get]
func GetCommentById(c *fiber.Ctx) error {
	return c.SendString("GetCommentById endpoint - To be implemented")
}


//	@Summary		Update comment
//	@Description	Update a specific comment by its ID
//	@Tags			Comments
//	@Accept			json
//	@Produce		json
//	@Param			recipeId	path		int										true	"Recipe ID"	minimum(1)
//	@Param			commentId	path		int										true	"Comment ID"	minimum(1)
//	@Param			comment	body	models.UpdateCommentRequest	true	"Comment"
//	@Success		200	{object}	models.UpdateResponse	"Successfully updated comment"
//	@Failure		400	{object}	models.ErrorBadRequestResponse	"Bad Request"
//	@Failure		404	{object}	models.ErrorNotFoundResponse	"Comment not found"
//	@Failure		500	{object}	models.ErrorInternalServerErrorResponse	"Internal Server Error"
//	@Router			/comments/{recipeId}/{commentId} [put]	
func UpdateComment(c *fiber.Ctx) error {
	return c.SendString("UpdateComment endpoint - To be implemented")
}


//	@Summary		Delete comment
//	@Description	Delete a specific comment by its ID
//	@Tags			Comments
//	@Accept			json
//	@Produce		json
//	@Param			recipeId	path		int										true	"Recipe ID"	minimum(1)
//	@Param			commentId	path		int										true	"Comment ID"	minimum(1)
//	@Success		200	{object}	models.DeleteResponse	"Successfully deleted comment"
//	@Failure		400	{object}	models.ErrorBadRequestResponse	"Bad Request"
//	@Failure		404	{object}	models.ErrorNotFoundResponse	"Comment not found"
//	@Failure		500	{object}	models.ErrorInternalServerErrorResponse	"Internal Server Error"	
//	@Router			/comments/{recipeId}/{commentId} [delete]
func DeleteComment(c *fiber.Ctx) error {
	return c.SendString("DeleteComment endpoint - To be implemented")
}	


//	@Summary		Create comment
//	@Description	Create a new comment
//	@Tags			Comments
//	@Accept			json
//	@Produce		json
//	@Param			recipeId	path		int										true	"Recipe ID"	minimum(1)
//	@Param			userId	path		int										true	"User ID"	minimum(1)
//	@Param			fatherCommentId	path		int										true	"Father Comment ID"	minimum(1)
//	@Param			comment	body	models.CreateCommentRequest	true	"Comment"
//	@Success		200	{object}	models.CreateResponse	"Successfully created comment"
//	@Failure		400	{object}	models.ErrorBadRequestResponse	"Bad Request"
//	@Failure		500	{object}	models.ErrorInternalServerErrorResponse	"Internal Server Error"
//	@Router			/comments/{recipeId}/{userId}/{fatherCommentId} [post]
func CreateComment(c *fiber.Ctx) error {
	return c.SendString("CreateComment endpoint - To be implemented")
}
















