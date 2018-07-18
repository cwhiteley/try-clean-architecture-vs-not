package controllers

import "fmt"

// Error - error struct
type Error struct {
	Message string
	Code    int
}

// Error - set error interface to Error
func (e *Error) Error() string {
	return fmt.Sprintf("Error: %s [code=%d]", e.Message, e.Code)
}

// NewError - init error
func NewError(err error) *Error {
	return &Error{
		Message: err.Error(),
	}
}
