package utils

import (
	"net/http"
)

type RestErr struct {
	// Id         string `json: "id"`
	Message    string `json: "message"`
	StatusCode int    `json: "status_code"`
	Code       string `json: "code"`
}

//Save new data error
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Code:       "bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
