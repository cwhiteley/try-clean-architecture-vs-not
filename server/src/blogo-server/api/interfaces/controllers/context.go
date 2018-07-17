package controllers

// Context - context interface
type Context interface {
	Param(string) string
	Bind(interface{}) error
	Status(int)
	JSON(int, interface{})
	Redirect(int, string)
}
