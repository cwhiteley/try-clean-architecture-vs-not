package controllers

import (
	"blogo-server/api/domain"
	"blogo-server/api/interfaces/database"
	"blogo-server/api/usecase"
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

// Create - create new post
func (controller *PostController) Create(c Context) {
	p := domain.Post{}
	c.Bind(&p)
	post, err := controller.Interactor.Add(p)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, post)
}

// Index - get all posts
func (controller *PostController) Index(c Context) {
	posts, err := controller.Interactor.Posts()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, posts)
}

// Show - get post by ID
func (controller *PostController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	post, err := controller.Interactor.PostByID(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, post)
}
