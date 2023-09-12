package errors

import "fmt"

var _ error = (*Error)(nil)

// New returns a new Error
func New(code int, message string) *Error {
	return &Error{code, message, nil}
}

// GetCode gets the Error code
func GetCode(e error) int {
	if err, ok := e.(*Error); ok {
		return err.Code
	}
	return 0
}

type Error struct {
	Code    int
	Message string
	Cause   error
}

// Error implements error.
func (e *Error) Error() string {
	return fmt.Sprintf("code: %d message: %v cause: %v", e.Code, e.Message, e.Cause)
}

// Wrap
func (e *Error) Wrap(cause error) *Error {
	e.Cause = cause
	return e
}
