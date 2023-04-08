package response

type Response struct {
	Message string `json:"message,omitempty"`
}

type ErrorResponse struct {
	StatusCode       int    `json:"code,omitempty"`
	DeveloperMessage string `json:"developerMessage"`
}
