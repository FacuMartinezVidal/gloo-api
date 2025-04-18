package routes

import "github.com/gofiber/fiber/v2"

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
		"data": fiber.Map{
			"searchParam": searchParam,
			"recipes": []fiber.Map{
				{
					"id": 1,
					"name": "Chicken Wings",
					"description": "A feast for the senses",
					"likes": 100,
					"comments": 50,
					"stars": 4.5,
					"estimated_time": "1 hour",
					"category": category,
					"ingredients": []string{"chicken", "wings", "salt", "pepper"},
					"instructions": []string{"1. Mix the chicken with salt and pepper", "2. Bake the chicken in the oven for 1 hour", "3. Serve the chicken with wings"},
				},
			},
		},
	})
}

func GetRecipeById(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}


