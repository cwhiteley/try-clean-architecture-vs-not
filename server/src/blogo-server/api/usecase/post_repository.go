package usecase

import "blogo-server/api/domain"

// PostRepository - post interface
type PostRepository interface {
	FindAll() (domain.Posts, error)
	FindByID(int) (domain.Post, error)
	Store(p domain.Post) (id int, err error)
}
