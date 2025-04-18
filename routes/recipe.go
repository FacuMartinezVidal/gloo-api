package routes

import (
	"gloo-api/models"

	"github.com/gofiber/fiber/v2"
)

type GetAllRecipeResponseData struct {
	SearchParam string          `json:"searchParam" example:"chicken"`
	Category    string          `json:"category" example:"Chicken"`
	Recipes     []models.Recipe `json:"recipes" `
}

//	@Summary		Get all recipes
//	@Description	Get all recipes by search and category
//	@Tags			Recipes
//	@Accept			json
//	@Produce		json
//	@Param			searchParam	query		string							false	"Search term for recipes"
//	@Param			category	query		string							false	"Category filter for recipes"
//	@Success		200			{object}	models.BaseResponse{data=GetAllRecipeResponseData}	"Successfully retrieved recipes"
//	@Failure		400			{object}	models.ErrorBadRequestResponse	"Bad Request"
//	@Failure		500			{object}	models.ErrorInternalServerErrorResponse	"Internal Server Error"
//  @Failure		404			{object}	models.ErrorNotFoundResponse			"Not Found"
//	@Router			/recipes [get]
func GetAllRecipe(c *fiber.Ctx) error {
	return c.SendString("GetAllRecipe endpoint - To be implemented")
}

//	@Summary		Get recipe by ID
//	@Description	Get a specific recipe by its ID
//	@Tags			Recipes
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int										true	"Recipe ID"	minimum(1)
//	@Success		200	{object}	models.BaseResponse{data=models.Recipe}	"Successfully retrieved recipe"
//	@Failure		400	{object}	models.ErrorBadRequestResponse					"Bad Request"
//	@Failure		404	{object}	models.ErrorNotFoundResponse					"Recipe not found"
//	@Failure		500	{object}	models.ErrorInternalServerErrorResponse					"Internal Server Error"
//	@Router			/recipes/{id} [get]
func GetRecipeById(c *fiber.Ctx) error {
	return c.SendString("GetRecipeById endpoint - To be implemented")
}
