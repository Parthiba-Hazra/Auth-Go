package models

import (
	"errors"
	"net/http"
)

type Type string

var (
	ErrInvalidRequest          Type = "INVALIDREQUEST"  // Validation errors
	ErrUnauthorizedClient      Type = "AUTHORIZATION"   // Authentication failures
	ErrAccessDenied            Type = "ACESSDENY"       // Database acess denied
	ErrUnsupportedResponseType Type = "RESPONSE"        // Invalid response type
	ErrInvalidScope            Type = "SCOPE"           // Invalid scope
	ErrServerError             Type = "SERVERERR"       // Server fallback (500) error
	ErrNotFound                Type = "NOTFOUND"        // Resources not found
	ErrConflict                Type = "CONFLICT"        // Already exist account
	ErrPayloadTooLarge         Type = "PAYLOADTOOLARGE" // Uploading a lot of JSON, image too big size
)

// It holds the error type and error message
type Error struct {
	Type    Type   `json:"errType"`
	Message string `json:"message"`
}

// Function that return the error message
func (e *Error) Error() string {
	return e.Message
}

func (e *Error) ErrorStatus() int {
	switch e.Type {
	case ErrInvalidRequest:
		return http.StatusBadRequest
	case ErrUnauthorizedClient:
		return http.StatusUnauthorized
	case ErrAccessDenied:
		return http.StatusForbidden
	case ErrUnsupportedResponseType:
		return http.StatusUnauthorized
	case ErrInvalidScope:
		return http.StatusBadRequest
	case ErrServerError:
		return http.StatusInternalServerError
	case ErrNotFound:
		return http.StatusNotFound
	case ErrConflict:
		return http.StatusConflict
	case ErrPayloadTooLarge:
		return http.StatusRequestEntityTooLarge
	default:
		return http.StatusInternalServerError
	}
}

// Check model error
func ErrorStatus(err error) int {
	var e *Error
	if errors.As(err, &e) {
		return e.ErrorStatus()
	}
	return http.StatusInternalServerError
}

// Error functions -- >>
