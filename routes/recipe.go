package routes

import (
	"gloo-api/models"

	"github.com/gofiber/fiber/v2"
)

//	@Summary		Get all recipes
//	@Description	Get all recipes
//	@Tags			recipes
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	fiber.Map
//	@Failure		400	{object}	fiber.Map
//	@Failure		500	{object}	fiber.Map
func GetAllRecipe(c *fiber.Ctx) error {
	searchParam := c.Params("searchParam")
	category := c.Params("category")
	return c.JSON(fiber.Map{
		"status": "success",
		"status_code": 200,
		"success": true,
		"data": []models.Recipe{
			{
				ID: 1,
				Title: "Chicken Wings",
				Description: "A feast for the senses",
			},
		},
		},
	})
}

func GetRecipeById(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
