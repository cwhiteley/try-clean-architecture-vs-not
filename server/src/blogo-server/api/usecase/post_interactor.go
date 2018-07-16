package usecase

import (
	"blogo-server/api/domain"
)

// PostInteractor - post struct
type PostInteractor struct {
	PostRepository PostRepository
}

// Add - find all users
func (interactor *PostInteractor) Add(p domain.Post) (post domain.Post, err error) {
	identifier, err := interactor.PostRepository.StorePost(p)
	if err != nil {
		return
	}
	post, err = interactor.PostRepository.FindByID(identifier)
	return
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

// Remove - remove post
func (interactor *PostInteractor) Remove(id int) (err error) {
	err = interactor.PostRepository.DeletePost(id)
	return
}
