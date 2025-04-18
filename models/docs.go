package models


type ErrorBadRequestResponse struct {
	Status     string      `json:"status" example:"error"`
	StatusCode int         `json:"status_code" example:"400"`
	Success    bool        `json:"success" example:"false"`
	Message    string      `json:"message" example:"Bad Request"`
	Data       interface{} `json:"data,omitempty"`
}

type ErrorInternalServerErrorResponse struct {
	Status     string      `json:"status" example:"error"`
	StatusCode int         `json:"status_code" example:"500"`
	Success    bool        `json:"success" example:"false"`
	Message    string      `json:"message" example:"Internal Server Error"`
	Data       interface{} `json:"data,omitempty"`
}

type ErrorNotFoundResponse struct {
	Status     string      `json:"status" example:"error"`
	StatusCode int         `json:"status_code" example:"404"`
	Success    bool        `json:"success" example:"false"`
	Message    string      `json:"message" example:"Not Found"`
	Data       interface{} `json:"data,omitempty"`
}


