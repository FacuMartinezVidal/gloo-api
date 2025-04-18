package handlers

import "github.com/gofiber/fiber/v2"

//	@Summary		Get favorites
//	@Description	Get favorites
//	@Tags			Favorites
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		string	true	"User ID"
//	@Success		200	{object}	models.GetFavoritesResponse	"Successfully retrieved favorites"
//	@Failure		400	{object}	models.ErrorBadRequestResponse	"Bad Request"
//	@Failure		500	{object}	models.ErrorInternalServerErrorResponse	"Internal Server Error"
//	@Failure		404	{object}	models.ErrorNotFoundResponse	"Not Found"
//	@Router			/favorites/{userId} [get]
func GetFavorites(c *fiber.Ctx) error {
	return c.SendString("getFavorites endpoint - To be implemented")
}
