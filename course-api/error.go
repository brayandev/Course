package api

import "fmt"

// ErrorType error types.
type ErrorType int

// ErrorType error type values.
const (
	ErrorUnknown ErrorType = iota
	ErrorInvalidRequest
)

// ErrorCodes
const (
	ErrorCodeUnknown string = "CRS0000"

	// Common error CRS1xxx.
	ErrorCodeInvalidRequest string = "CRS1000"
)

// Error representation of error.
type Error struct {
	ErrCode string    `json:"error,omitempty"`
	Message string    `json:"message,omitempty"`
	ErrType ErrorType `json:"-"`
}

// NewError constructor of error.
func NewError(code, message string, errType ErrorType) *Error {
	return &Error{ErrCode: code, Message: message, ErrType: errType}
}

// NewUnknownError constructor of unknown error.
func NewUnknownError(message string) *Error {
	return NewError(ErrorCodeUnknown, message, ErrorUnknown)
}

// NewInvalidRequestError constructor of Invalid Request error.
func NewInvalidRequestError(message string) *Error {
	return NewError(ErrorCodeInvalidRequest, message, ErrorInvalidRequest)
}

// Error return a string representation of and Error.
func (e Error) Error() string {
	return fmt.Sprintf("%s - %s", e.ErrCode, e.Message)
}

// Version represents a version of error.
func (e Error) Version() string {
	return "course.error.v1"
}
