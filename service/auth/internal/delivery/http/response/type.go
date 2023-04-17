package response

type Response struct {
	Message string `json:"message,omitempty"`
}

type LoginResponse struct {
	Message string `json:"message,omitempty"`
	Token   string `json:"token"`
}

type ErrorResponse struct {
	StatusCode       int    `json:"code,omitempty"`
	DeveloperMessage string `json:"developerMessage"`
}
