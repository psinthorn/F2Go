package utils

type ApplicationError struct {
	// Id         string `json: "id"`
	Message    string `json: "message"`
	StatusCode int    `json: "status_code"`
	Code       string `json: "code"`
}
