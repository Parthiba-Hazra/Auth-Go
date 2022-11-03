package error

import (
	"errors"
	"fmt"
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

func NewInvalidReq(cause string) *Error {
	return &Error{
		Type:    ErrInvalidRequest,
		Message: fmt.Sprintf("Invalid request, cause: %v", cause),
	}
}

func NewAuthorizationErr(cause string) *Error {
	return &Error{
		Type:    ErrUnauthorizedClient,
		Message: cause,
	}
}

func NewAcessDenied(cause string) *Error {
	return &Error{
		Type:    ErrAccessDenied,
		Message: fmt.Sprintf("Server refuse to authorize the request, cause: %v", cause),
	}
}

func NewUnsupportedResonse(cause string) *Error {
	return &Error{
		Type:    ErrUnsupportedResponseType,
		Message: fmt.Sprintf("Unsupported response type, cause: %v", cause),
	}
}

func NewInvalidScope(cause string) *Error {
	return &Error{
		Type:    ErrInvalidScope,
		Message: fmt.Sprintf("Bad request , cause: %v", cause),
	}
}

func NewServerErr() *Error {
	return &Error{
		Type:    ErrServerError,
		Message: fmt.Sprintf("Internal server error"),
	}
}

func NewErrNotFound(name string, value string) *Error {
	return &Error{
		Type:    ErrNotFound,
		Message: fmt.Sprintf("cant found %v resourses with value %v", name, value),
	}
}

func NewErrConflict(name string, value string) *Error {
	return &Error{
		Type:    ErrConflict,
		Message: fmt.Sprintf("%v resourses with value %v, already exist", name, value),
	}
}

func NewErrPauloadLarge(maxSize int64, currentSize int64) *Error {
	return &Error{
		Type:    ErrPayloadTooLarge,
		Message: fmt.Sprintf("Current playload size %v, exceed max size %v", currentSize, maxSize),
	}
}
