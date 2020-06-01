package error_utils

import (
	"net/http"
)

const (
	InternalCustomError = "internal_custom_error"
	InternalServerError = "internal_server_error"
	BadRequestError     = "bad_request_error"
	UnAuthorizedError   = "unauthorized_error"
	ForbiddenError      = "forbidden_error"
)

type ApiError struct {
	HttpStatusCode     int    `json:"http_status_code"`
	InternalStatusCode int    `json:"internal_status_code"`
	ErrorConst         string `json:"error_const"`
	ErrorMessage       string `json:"error_message"`
}

func NewInternalCustomError(internalStatusCode int, message string) *ApiError {
	return &ApiError{
		HttpStatusCode:     http.StatusInternalServerError,
		InternalStatusCode: internalStatusCode,
		ErrorMessage:       message,
		ErrorConst:         InternalCustomError,
	}
}

func NewInternalServerError(message string) *ApiError {
	return &ApiError{
		HttpStatusCode:     http.StatusInternalServerError,
		InternalStatusCode: http.StatusInternalServerError,
		ErrorMessage:       message,
		ErrorConst:         InternalServerError,
	}
}

func NewBadRequestError(message string) *ApiError {
	return &ApiError{
		HttpStatusCode:     http.StatusBadRequest,
		InternalStatusCode: http.StatusBadRequest,
		ErrorMessage:       message,
		ErrorConst:         BadRequestError,
	}
}

func NewUnauthorizedError(message string) *ApiError {
	return &ApiError{
		HttpStatusCode:     http.StatusUnauthorized,
		InternalStatusCode: http.StatusUnauthorized,
		ErrorMessage:       message,
		ErrorConst:         UnAuthorizedError,
	}
}

func NewForbiddenError(message string) *ApiError {
	return &ApiError{
		HttpStatusCode:     http.StatusForbidden,
		InternalStatusCode: http.StatusForbidden,
		ErrorMessage:       message,
		ErrorConst:         ForbiddenError,
	}
}
