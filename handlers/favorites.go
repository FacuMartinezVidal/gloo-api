package handlers

import "github.com/gofiber/fiber/v2"

//	@Summary		Get favorites
//	@Description	Get favorites
//	@Tags			Favorites
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		string									true	"User ID"
//	@Success		200		{object}	models.GetFavoritesResponse				"Successfully retrieved favorites"
//	@Failure		400		{object}	models.ErrorBadRequestResponse			"Bad Request"
//	@Failure		500		{object}	models.ErrorInternalServerErrorResponse	"Internal Server Error"
//	@Failure		404		{object}	models.ErrorNotFoundResponse			"Not Found"
//	@Router			/favorites/{userId} [get]
func GetFavorites(c *fiber.Ctx) error {
	return c.SendString("getFavorites endpoint - To be implemented")
}


// @Summary Create a favorite
// @Description Create a favorite
// @Tags Favorites
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Param recipeId path string true "Recipe ID"
// @Param favorite body models.CreateFavoriteRequest true "Favorite"
// @Success 201 {object} models.CreateResponse
// @Failure 400 {object} models.ErrorBadRequestResponse
// @Failure 404 {object} models.ErrorNotFoundResponse
// @Failure 500 {object} models.ErrorInternalServerErrorResponse
// @Router /favorites/{userId}/{recipeId} [post]	
func CreateFavorite(c *fiber.Ctx) error {
	return c.SendString("CreateFavorite endpoint - To be implemented")
}


// @Summary Delete a favorite
// @Description Delete a favorite
// @Tags Favorites
// @Accept json
// @Produce json	
// @Param userId path string true "User ID"
// @Param recipeId path string true "Recipe ID"
// @Success 200 {object} models.DeleteResponse
// @Failure 400 {object} models.ErrorBadRequestResponse
// @Failure 404 {object} models.ErrorNotFoundResponse
// @Failure 500 {object} models.ErrorInternalServerErrorResponse
// @Router /favorites/{userId}/{recipeId}/{favoriteId} [delete]
func DeleteFavorite(c *fiber.Ctx) error {
	return c.SendString("DeleteFavorite endpoint - To be implemented")
}




