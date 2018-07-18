package handler

import (
	"blogo-server/server/repository"
	"net/http"
)

// Repository - SQL controller
type Repository struct {
	repo repository.Repository
}

// CreatePost - create new post handler
func CreatePost(c Context) {
	c.JSON(http.StatusCreated, "www")
}

// GetPost - get all post handler
func GetPost(c Context) {
	c.JSON(http.StatusOK, "www")
}

// ListPosts - get post handler by id
func ListPosts(c Context) {
	c.JSON(http.StatusOK, "www")
}

// DeletePost - delete post handler
func DeletePost(c Context) {
	c.JSON(http.StatusNoContent, "www")
}
