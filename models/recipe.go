package models

type Recipe struct {
	ID             int     `json:"id" example:"1"`
	Title          string  `json:"title" example:"Chicken Wings"`
	Description    string  `json:"description" example:"A feast for the senses"`
	EstimatedTime  string  `json:"estimated_time" example:"20 minutes"`
	ImageUrl       string  `json:"image_url" example:"https://example.com/image.png"`
	CreatedBy      string  `json:"created_by" example:"@facu.potti"`
	TotalComments  int     `json:"total_comments" example:"64"`
	TotalFavorites int     `json:"total_favorites" example:"4445"`
	TotalStars     float64 `json:"total_stars" example:"4"`
	Category       string  `json:"category" example:"Chicken"`
}
