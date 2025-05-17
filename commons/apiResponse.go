package commons

type ApiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type SwaggerApiResponseSuccessWithData struct {
	Success bool        `json:"success" default:"true"`
	Message string      `json:"message"`
	Data    interface{} `json:"data" default:{}`
}

type SwaggerApiResponseSuccessWithoutData struct {
	Success bool        `json:"success" default:"true"`
	Message string      `json:"message"`
	Data    interface{} `json:"data" default:nil`
}

type SwaggerApiResponseError struct {
	Success bool        `json:"success" default:"false"`
	Message string      `json:"message"`
	Data    interface{} `json:"data" default:nil`
}
