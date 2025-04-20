package handlers

import "github.com/gofiber/fiber/v2"

// @Summary Get a multiple
// @Description Get a multiple
// @Tags Multiple
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Param recipeId path string true "Recipe ID"
// @Success 200 {object} models.GetMultipleResponse
// @Failure 400 {object} models.ErrorBadRequestResponse
// @Failure 404 {object} models.ErrorNotFoundResponse
// @Failure 500 {object} models.ErrorInternalServerErrorResponse
// @Router /multiple/{userId}/{recipeId} [get]
func GetMultiple(c *fiber.Ctx) error {
	return c.SendString("GetMultiple endpoint - To be implemented")
}	

// @Summary Create a multiple
// @Description Create a multiple
// @Tags Multiple
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Param recipeId path string true "Recipe ID"
// @Param multiple body models.CreateMultipleRequest true "Multiple"
// @Success 201 {object} models.CreateResponse
// @Failure 400 {object} models.ErrorBadRequestResponse
// @Failure 404 {object} models.ErrorNotFoundResponse
// @Failure 500 {object} models.ErrorInternalServerErrorResponse
// @Router /multiple/{userId}/{recipeId} [post]
func CreateMultiple(c *fiber.Ctx) error {
	return c.SendString("CreateMultiple endpoint - To be implemented")
}

// @Summary Update a multiple
// @Description Update a multiple
// @Tags Multiple
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Param recipeId path string true "Recipe ID"
// @Param multiple body models.CreateMultipleRequest true "Multiple"
// @Success 200 {object} models.UpdateResponse
// @Failure 400 {object} models.ErrorBadRequestResponse
// @Failure 404 {object} models.ErrorNotFoundResponse
// @Failure 500 {object} models.ErrorInternalServerErrorResponse
// @Router /multiple/{userId}/{recipeId} [put]
func UpdateMultiple(c *fiber.Ctx) error {
	return c.SendString("UpdateMultiple endpoint - To be implemented")
}
