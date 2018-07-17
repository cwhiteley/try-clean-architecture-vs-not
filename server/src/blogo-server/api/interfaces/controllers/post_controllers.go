package controllers

import (
	"blogo-server/api/domain"
	"blogo-server/api/interfaces/database"
	"blogo-server/api/usecase"
	"net/http"
	"strconv"
)

// PostController - post controller
type PostController struct {
	Interactor usecase.PostInteractor
}

// NewPostController - init post contoroller
func NewPostController(SQLHandler database.SQLHandler) *PostController {
	return &PostController{
		Interactor: usecase.PostInteractor{
			PostRepository: &database.PostRepository{
				SQLHandler: SQLHandler,
			},
		},
	}
}

// Create - create new post controller
func (controller *PostController) Create(c Context) {
	p := domain.Post{}
	c.Bind(&p)
	post, err := controller.Interactor.Add(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusCreated, post)
}

// Index - get all controller
func (controller *PostController) Index(c Context) {
	posts, err := controller.Interactor.Posts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusOK, posts)
}

// Show - get post controller by ID
func (controller *PostController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	post, err := controller.Interactor.PostByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusOK, post)
}

// Delete - delete controller post
func (controller *PostController) Delete(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := controller.Interactor.Remove(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
