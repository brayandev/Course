package api

import "fmt"

// ErrorType error types.
type ErrorType int

// ErrorType error type values.
const (
	ErrorUnknown ErrorType = iota
)

// ErrorCodes
const (
	ErrorCodeUnknown string = "CRS0000"

	// Common error CRS1xxx.
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

// Error return a string representation of and Error.
func (e Error) Error() string {
	return fmt.Sprintf("%s - %s", e.ErrCode, e.Message)
}

// Version represents a version of error.
func (e Error) Version() string {
	return "vnd.catho.company.error.v1"
}
