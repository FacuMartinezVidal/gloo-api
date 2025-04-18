package handlers

import (
	"github.com/gofiber/fiber/v2"
)

//	@Summary		Update ingredient
//	@Description	Update a specific ingredient by its ID
//	@Tags			Ingredients
//	@Accept			json
//	@Produce		json
//	@Param			recipeId	path		int										true	"Recipe ID"	minimum(1)
//	@Param			ingredientId	path		int										true	"Ingredient ID"	minimum(1)
//	@Param			ingredient	body	models.UpdateIngredientRequest	true	"Ingredient"
//	@Success		200	{object}	models.UpdateResponse	"Successfully updated ingredient"
//	@Failure		400	{object}	models.ErrorBadRequestResponse			"Bad Request"
//	@Failure		404	{object}	models.ErrorNotFoundResponse			"Ingredient not found"
//	@Failure		500	{object}	models.ErrorInternalServerErrorResponse	"Internal Server Error"
//	@Router			/ingredients/{recipeId}/{ingredientId} [put]
func UpdateIngredient(c *fiber.Ctx) error {
	return c.SendString("UpdateIngredient endpoint - To be implemented")
}



//	@Summary		Create ingredient
//	@Description	Create a new ingredient
//	@Tags			Ingredients
//	@Accept			json
//	@Produce		json
//	@Param			recipeId	path		int										true	"Recipe ID"	minimum(1)
//	@Param			ingredientId	path		int										true	"Ingredient ID"	minimum(1)
//	@Param			ingredient	body	models.CreateIngredientRequest	true	"Ingredient"	
//	@Success		200	{object}	models.CreateResponse	"Successfully created ingredient"
//	@Failure		400	{object}	models.ErrorBadRequestResponse			"Bad Request"
//	@Failure		500	{object}	models.ErrorInternalServerErrorResponse	"Internal Server Error"
//	@Failure		404	{object}	models.ErrorNotFoundResponse			"Ingredient not found"
//	@Router			/ingredients/{recipeId}/{ingredientId} [post]
func CreateIngredient(c *fiber.Ctx) error {
	return c.SendString("CreateIngredient endpoint - To be implemented")
}




//	@Summary		Delete ingredient
//	@Description	Delete a specific ingredient by its ID
//	@Tags			Ingredients
//	@Accept			json
//	@Produce		json
//	@Param			recipeId	path		int										true	"Recipe ID"	minimum(1)
//	@Param			ingredientId	path		int										true	"Ingredient ID"	minimum(1)
//	@Success		200	{object}	models.DeleteResponse	"Successfully deleted ingredient"
//	@Failure		400	{object}	models.ErrorBadRequestResponse			"Bad Request"
//	@Failure		404	{object}	models.ErrorNotFoundResponse			"Ingredient not found"
//	@Failure		500	{object}	models.ErrorInternalServerErrorResponse	"Internal Server Error"
//	@Router			/ingredients/{recipeId}/{ingredientId} [delete]
func DeleteIngredient(c *fiber.Ctx) error {
	return c.SendString("DeleteIngredient endpoint - To be implemented")
}






