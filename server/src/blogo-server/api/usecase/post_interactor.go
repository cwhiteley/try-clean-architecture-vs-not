package usecase

import "blogo-server/api/domain"

// PostInteractor - post struct
type PostInteractor struct {
	PostRepository PostRepository
}

// Posts - find all users
func (interactor *PostInteractor) Posts() (Post domain.Posts, err error) {
	Post, err = interactor.PostRepository.FindAll()
	return
}

// PostByID - find post
func (interactor *PostInteractor) PostByID(identifier int) (Post domain.Post, err error) {
	Post, err = interactor.PostRepository.FindById(identifier)
	return
}
