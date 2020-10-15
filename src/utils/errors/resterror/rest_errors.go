package resterror

import "net/http"

// RestError common error interface for all errors
// to be returned to the user
type RestError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
}

// NewNotFoundError returns a RestError
// in the format of not found
func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message:    message,
		StatusCode: http.StatusNotFound,
		Error:      "not_found",
	}
}

// NewBadRequest returns a RestError
// in the format of abd request
func NewBadRequest(message string) *RestError {
	return &RestError{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Error:      "bad_request",
	}
}
