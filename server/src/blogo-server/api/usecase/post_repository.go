package usecase

import "blogo-server/api/domain"

// PostRepository - post interface
type PostRepository interface {
	FindById(int) (domain.Post, error)
	FindAll() (domain.Posts, error)
}
