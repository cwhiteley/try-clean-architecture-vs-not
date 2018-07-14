package database

import "blogo-server/api/domain"

// PostRepository - sql handler
type PostRepository struct {
	SQLHandler
}

// FindByID - select post
func (repo *PostRepository) FindByID(identifier int) (post domain.Post, err error) {
	row, err := repo.Query("SELECT id, title, content FROM posts WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int64
	var title string
	var content string
	row.Next()
	if err = row.Scan(&id, &title, &content); err != nil {
		return
	}
	post.ID = id
	post.Title = title
	post.Content = content
	return
}

// FindAll - select all
func (repo *PostRepository) FindAll() (posts domain.Posts, err error) {
	rows, err := repo.Query("SELECT id, title, context FROM posts")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int64
		var title string
		var content string
		if err := rows.Scan(&id, &title, &content); err != nil {
			continue
		}
		post := domain.Post{
			ID:      id,
			Title:   title,
			Content: content,
		}
		posts = append(posts, post)
	}
	return
}
