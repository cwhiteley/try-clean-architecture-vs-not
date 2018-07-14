package controllers

import (
	"blogo-server/api/usecase"
	"strconv"
)

// PostController - post controller
type PostController struct {
	Interactor usecase.PostInteractor
}

// Index - get all posts
func (controller *PostController) Index(c Content) {
	posts, err := controller.Interactor.Posts()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, posts)
}

// Show - get post
func (controller *PostController) Show(c Content) {
	id, _ := strconv.Atoi(c.Param("id"))
	post, err := controller.Interactor.PostById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, post)
}
