package models

type Recipe struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	EstimatedTime string `json:"estimated_time"`
	ImageUrl      string `json:"image_url"`
	CreatedBy     string `json:"created_by"`
}
