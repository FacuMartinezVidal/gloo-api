package models

type BaseResponse struct {
	Status     string      `json:"status" example:"success"`
	StatusCode int         `json:"status_code" example:"200"`
	Success    bool        `json:"success" example:"true"`
	Data       interface{} `json:"data,omitempty"`
}
