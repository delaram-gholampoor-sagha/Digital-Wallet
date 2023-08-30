package derror

import (
	"fmt"
	"net/http"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/derror/message"
)

// Error represents a custom error with a message, arguments, and an HTTP status code.
type Error struct {
	Message string
	Args    []interface{}
	Code    int
}

// NewError creates a generic error with a custom message, status code, and arguments.
func NewError(msg string, code int, args ...interface{}) error {
	return &Error{
		Message: msg,
		Code:    code,
		Args:    args,
	}
}

func NewNotFoundError(msg string, args ...any) error {
	return NewError(msg, http.StatusNotFound, args...)
}

// Error formats the error message.
func (e *Error) Error() string {
	return fmt.Sprintf(e.Message, e.Args...)
}

// NewInternalSystemError creates a 500 Internal Server Error.
func NewInternalSystemError() error {
	return NewError(message.InternalSystemError, http.StatusInternalServerError)
}

// NewBadRequestError creates a 400 Bad Request error.
func NewBadRequestError(msg string, args ...interface{}) error {
	return NewError(msg, http.StatusBadRequest, args...)
}

// IsHTTPError checks if the error corresponds to a specific HTTP status code.
func IsHTTPError(err error, code int) bool {
	derr, ok := err.(*Error)
	return ok && derr.Code == code
}

// IsInternalError checks whether the provided error is an internal error
func IsInternalError(err error) bool {
	return IsHTTPError(err, http.StatusInternalServerError)
}

func IsNotFound(err error) bool {
	derr, ok := err.(*Error)
	return ok && derr.Code == http.StatusNotFound
}
