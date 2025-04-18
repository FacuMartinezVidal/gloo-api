package models

// RecipeExampleResponse - Modelo específico para ejemplos de Swagger
type RecipeExampleResponse struct {
	Status     string `json:"status" example:"success"`
	StatusCode int    `json:"status_code" example:"200"`
	Success    bool   `json:"success" example:"true"`
	Message    string `json:"message,omitempty" example:"Operación completada con éxito"`
	Data       struct {
		SearchParam string `json:"searchParam" example:"chicken"`
		Category    string `json:"category" example:"Chicken"`
		Recipes     []struct {
			ID            int     `json:"id" example:"1"`
			Title         string  `json:"title" example:"Chicken Wings"`
			Description   string  `json:"description" example:"A feast for the senses"`
			EstimatedTime string  `json:"estimated_time" example:"15 minutes"`
			ImageUrl      string  `json:"image_url" example:"https://example.com/image.png"`
			CreatedBy     string  `json:"created_by" example:"@facu.potti"`
			TotalComments int     `json:"total_comments" example:"10"`
			TotalLikes    int     `json:"total_likes" example:"20"`
			TotalStars    float64 `json:"total_stars" example:"4.5"`
			Category      string  `json:"category" example:"Chicken"`
		} `json:"recipes"`
	} `json:"data"`
}

// ErrorResponse - Modelo específico para ejemplos de errores en Swagger
type ErrorResponse struct {
	Status     string      `json:"status" example:"error"`
	StatusCode int         `json:"status_code" example:"400"`
	Success    bool        `json:"success" example:"false"`
	Message    string      `json:"message" example:"Solicitud inválida"`
	Data       interface{} `json:"data,omitempty"`
}
