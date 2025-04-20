package handlers

import (
	"github.com/gofiber/fiber/v2"
)

//	@Summary		Get all collections
//	@Description	Get all collections
//	@Tags			Collections
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		string									true	"User ID"
//	@Success		200		{object}	models.GetAllCollectionsResponse		"Successfully retrieved collections"
//	@Failure		400		{object}	models.ErrorBadRequestResponse			"Bad Request"
//	@Failure		404		{object}	models.ErrorNotFoundResponse			"Not Found"
//	@Failure		500		{object}	models.ErrorInternalServerErrorResponse	"Internal Server Error"
//	@Router			/collections/{userId} [get]
func GetAllCollections(c *fiber.Ctx) error {
	return c.SendString("GetAllCollections endpoint - To be implemented")
}


//	@Summary		Get collection by ID
//	@Description	Get collection by ID
//	@Tags			Collections
//	@Accept			json
//	@Produce		json
//	@Param			collectionId	path		string									true	"Collection ID"
//	@Success		200				{object}	models.GetCollectionByIdResponse		"Successfully retrieved collection"
//	@Failure		400				{object}	models.ErrorBadRequestResponse			"Bad Request"
//	@Failure		404				{object}	models.ErrorNotFoundResponse			"Not Found"
//	@Failure		500				{object}	models.ErrorInternalServerErrorResponse	"Internal Server Error"
//	@Router			/collections/{userId}/{collectionId} [get]
func GetCollectionById(c *fiber.Ctx) error {
	return c.SendString("GetCollectionById endpoint - To be implemented")
}




//	@Summary		Create a collection
//	@Description	Create a collection
//	@Tags			Collections
//	@Accept			json
//	@Produce		json
//	@Param			collection	body		models.CreateCollectionRequest			true	"Collection"
//	@Success		201			{object}	models.CreateResponse					"Collection created successfully"
//	@Failure		400			{object}	models.ErrorBadRequestResponse			"Bad Request"
//	@Failure		500			{object}	models.ErrorInternalServerErrorResponse	"Internal Server Error"
//	@Router			/collections/{userId} [post]
func CreateCollection(c *fiber.Ctx) error {
	return c.SendString("CreateCollection endpoint - To be implemented")
}






