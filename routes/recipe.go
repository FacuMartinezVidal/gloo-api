package routes

import (
	"gloo-api/models"

	"github.com/gofiber/fiber/v2"
)

type GetAllRecipeResponseData struct {
	SearchParam string          `json:"searchParam"`
	Category    string          `json:"category"`
	Recipes     []models.Recipe `json:"recipes"`
}

//	@Summary		Get all recipes
//	@Description	Get all recipes by search and category
//	@Tags			Recipes
//	@Accept			json
//	@Produce		json
//	@Param			searchParam	query		string												false	"Search term for recipes"
//	@Param			category	query		string												false	"Category filter for recipes"
//	@Success		200			{object}	models.BaseResponse{data=GetAllRecipeResponseData}	"Successfully retrieved recipes"
//	@Failure		400			{object}	models.BaseResponse									"Bad Request"
//	@Failure		500			{object}	models.BaseResponse									"Internal Server Error"
//	@Router			/recipes [get]
func GetAllRecipe(c *fiber.Ctx) error {
	return c.SendString("GetAllRecipe endpoint - To be implemented")
}

// TODO: Add Swaggo documentation and implementation for GetRecipeById
func GetRecipeById(c *fiber.Ctx) error {
	return c.SendString("GetRecipeById endpoint - To be implemented")
}
