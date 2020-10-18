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
// in the format of a bad request
func NewBadRequest(message string) *RestError {
	return &RestError{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Error:      "bad_request",
	}
}

// NewInternalServerError returns a RestError
// whenever something goes wrong in the server side
func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
		Error:      "internal_server_error",
	}
}

// NewUnAuthorizedError returns a RestError
// whenever some unauthorized user tries to
// get data from a secured endpoint
func NewUnAuthorizedError() *RestError {
	return &RestError{
		Message:    "invalid token",
		StatusCode: http.StatusUnauthorized,
		Error:      "unauthorized",
	}
}

// NewInvalidCredentialsError returns a RestError
// whenever some unauthorized user tries to
// get data from a secured endpoint
func NewInvalidCredentialsError() *RestError {
	return &RestError{
		Message:    "invalid credentials",
		StatusCode: http.StatusNonAuthoritativeInfo,
		Error:      "unauthorized",
	}
}
