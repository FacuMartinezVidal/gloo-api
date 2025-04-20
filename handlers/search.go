package handlers

import "github.com/gofiber/fiber/v2"

//	@Summary		Get all searches
//	@Description	Get all searches
//	@Tags			Search
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		string									true	"User ID"
//	@Success		200		{object}	models.GetAllSearchesResponse			"Successfully retrieved searches"
//	@Failure		400		{object}	models.ErrorBadRequestResponse			"Bad Request"
//	@Failure		500		{object}	models.ErrorInternalServerErrorResponse	"Internal Server Error"
//	@Failure		404		{object}	models.ErrorNotFoundResponse			"Not Found"
//	@Router			/search/{userId} [get]
func GetAllSearches(c *fiber.Ctx) error {
	return c.SendString("GetAllSearches endpoint - To be implemented")
}






//	@Summary		Get all suggestions
//	@Description	Get all suggestions
//	@Tags			Search
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		string									true	"User ID"
//	@Success		200		{object}	models.GetAllSearchesResponse			"Successfully retrieved searches"
//	@Failure		400		{object}	models.ErrorBadRequestResponse			"Bad Request"
//	@Failure		500		{object}	models.ErrorInternalServerErrorResponse	"Internal Server Error"
//	@Failure		404		{object}	models.ErrorNotFoundResponse			"Not Found"
//	@Router			/search/{userId}/suggestions [get]
func GetAllSuggestions(c *fiber.Ctx) error {
	return c.SendString("GetAllSuggestions endpoint - To be implemented")
}




