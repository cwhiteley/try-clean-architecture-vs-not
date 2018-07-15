package controllers

// Error - error struct
type Error struct {
	Message string
}

// NewError - init error
func NewError(err error) *Error {
	return &Error{
		Message: err.Error(),
	}
}
