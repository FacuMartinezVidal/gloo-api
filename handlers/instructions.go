package handlers

import "github.com/gofiber/fiber/v2"

//	@Summary		Update instruction
//	@Description	Update a specific instruction by its ID
//	@Tags			Instructions
//	@Accept			json
//	@Produce		json
//	@Param			recipeId		path		int										true	"Recipe ID"			minimum(1)
//	@Param			instructionId	path		int										true	"Instruction ID"	minimum(1)
//	@Param			instruction		body		models.UpdateInstructionRequest			true	"Instruction"
//	@Success		200				{object}	models.UpdateResponse					"Successfully updated instruction"
//	@Failure		400				{object}	models.ErrorBadRequestResponse			"Bad Request"
//	@Failure		404				{object}	models.ErrorNotFoundResponse			"Instruction not found"
//	@Failure		500				{object}	models.ErrorInternalServerErrorResponse	"Internal Server Error"
//	@Router			/instructions/{recipeId}/{instructionId} [put]
func UpdateInstruction(c *fiber.Ctx) error {
	return c.SendString("UpdateInstruction endpoint - To be implemented")
}


//	@Summary		Create instruction
//	@Description	Create a new instruction
//	@Tags			Instructions
//	@Accept			json
//	@Produce		json
//	@Param			recipeId		path		int										true	"Recipe ID"			minimum(1)
//	@Param			instructionId	path		int										true	"Instruction ID"	minimum(1)
//	@Param			instruction		body		models.CreateInstructionRequest			true	"Instruction"
//	@Success		200				{object}	models.CreateResponse					"Successfully created instruction"
//	@Failure		400				{object}	models.ErrorBadRequestResponse			"Bad Request"
//	@Failure		500				{object}	models.ErrorInternalServerErrorResponse	"Internal Server Error"
//	@Failure		404				{object}	models.ErrorNotFoundResponse			"Instruction not found"
//	@Router			/instructions/{recipeId}/{instructionId} [post]
func CreateInstruction(c *fiber.Ctx) error {
	return c.SendString("CreateInstruction endpoint - To be implemented")
}



//	@Summary		Delete instruction
//	@Description	Delete a specific instruction by its ID
//	@Tags			Instructions
//	@Accept			json
//	@Produce		json
//	@Param			recipeId		path		int										true	"Recipe ID"			minimum(1)
//	@Param			instructionId	path		int										true	"Instruction ID"	minimum(1)
//	@Success		200				{object}	models.DeleteResponse					"Successfully deleted instruction"
//	@Failure		400				{object}	models.ErrorBadRequestResponse			"Bad Request"
//	@Failure		404				{object}	models.ErrorNotFoundResponse			"Instruction not found"
//	@Failure		500				{object}	models.ErrorInternalServerErrorResponse	"Internal Server Error"
//	@Router			/instructions/{recipeId}/{instructionId} [delete]
func DeleteInstruction(c *fiber.Ctx) error {
	return c.SendString("DeleteInstruction endpoint - To be implemented")
}


//	@Summary		Get instruction by ID
//	@Description	Get a specific instruction by its ID
//	@Tags			Instructions
//	@Accept			json	
//	@Produce		json
//	@Param			recipeId										path		int										true	"Recipe ID"			minimum(1)
//	@Param			instructionId									path		int										true	"Instruction ID"	minimum(1)
//	@Success		200												{object}	models.GetInstructionByIdResponse		"Successfully retrieved instruction"
//	@Failure		400												{object}	models.ErrorBadRequestResponse			"Bad Request"
//	@Failure		404												{object}	models.ErrorNotFoundResponse			"Instruction not found"
//	@Failure		500												{object}	models.ErrorInternalServerErrorResponse	"Internal Server Error"
//	@Router			/instructions/{recipeId}/{instructionId} [get]	
func GetInstructionById(c *fiber.Ctx) error {
	return c.SendString("GetInstructionById endpoint - To be implemented")
}



