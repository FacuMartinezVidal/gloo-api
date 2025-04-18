package models

type Recipe struct {
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
}
