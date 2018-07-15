package usecase

import (
	"blogo-server/api/domain"
)

// PostInteractor - post struct
type PostInteractor struct {
	PostRepository PostRepository
}

// Posts - find all users
func (interactor *PostInteractor) Posts() (posts domain.Posts, err error) {
	posts, err = interactor.PostRepository.FindAll()
	return
}

// PostByID - find post
func (interactor *PostInteractor) PostByID(identifier int) (post domain.Post, err error) {
	post, err = interactor.PostRepository.FindByID(identifier)
	return
}
