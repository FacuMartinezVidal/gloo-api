package handlers

import (
	"github.com/gofiber/fiber/v2"
)

//	@Summary		Create a rate
//	@Description	Create a rate for a recipe
//	@Tags			Rates
//	@Accept			json
//	@Produce		json
//	@Param			recipeId	path		int	true	"Recipe ID"
//	@Param			userId		path		int	true	"User ID"
//	@Success		201			{object}	models.CreateResponse
//	@Failure		400			{object}	models.ErrorBadRequestResponse
//	@Failure		500			{object}	models.ErrorInternalServerErrorResponse
//	@Failure		404			{object}	models.ErrorNotFoundResponse
//	@Router			/rates/{recipeId}/{userId} [post]
func CreateRate(c *fiber.Ctx) error {
	return c.SendString("CreateRate endpoint - To be implemented")
}

//	@Summary		Get a rate
//	@Description	Get a rate for a recipe
//	@Tags			Rates
//	@Accept			json
//	@Produce		json
//	@Param			recipeId	path		int	true	"Recipe ID"
//	@Param			userId		path		int	true	"User ID"
//	@Success		200			{object}	models.GetRatesResponse
//	@Failure		400			{object}	models.ErrorBadRequestResponse
//	@Failure		500			{object}	models.ErrorInternalServerErrorResponse
//	@Failure		404			{object}	models.ErrorNotFoundResponse
//	@Router			/rates/{recipeId}/{userId} [get]
func GetRates(c *fiber.Ctx) error {
	return c.SendString("GetRate endpoint - To be implemented")
}	

//	@Summary		Update a rate
//	@Description	Update a rate for a recipe
//	@Tags			Rates
//	@Accept			json
//	@Produce		json
//	@Param			recipeId	path		int							true	"Recipe ID"
//	@Param			userId		path		int							true	"User ID"
//	@Param			rateId		path		int							true	"Rate ID"
//	@Param			rate		body		models.UpdateRateRequest	true	"Rate"
//	@Success		200			{object}	models.UpdateResponse
//	@Failure		400			{object}	models.ErrorBadRequestResponse
//	@Failure		500			{object}	models.ErrorInternalServerErrorResponse
//	@Failure		404			{object}	models.ErrorNotFoundResponse
//	@Router			/rates/{recipeId}/{userId}/{rateId} [put]
func UpdateRate(c *fiber.Ctx) error {
	return c.SendString("UpdateRate endpoint - To be implemented")
}

//	@Summary		Delete a rate
//	@Description	Delete a rate for a recipe
//	@Tags			Rates
//	@Accept			json
//	@Produce		json
//	@Param			recipeId	path		int	true	"Recipe ID"
//	@Param			userId		path		int	true	"User ID"
//	@Param			rateId		path		int	true	"Rate ID"
//	@Success		200			{object}	models.DeleteResponse
//	@Failure		400			{object}	models.ErrorBadRequestResponse
//	@Failure		500			{object}	models.ErrorInternalServerErrorResponse
//	@Failure		404			{object}	models.ErrorNotFoundResponse
//	@Router			/rates/{recipeId}/{userId}/{rateId} [delete]
func DeleteRate(c *fiber.Ctx) error {
	return c.SendString("DeleteRate endpoint - To be implemented")
}
