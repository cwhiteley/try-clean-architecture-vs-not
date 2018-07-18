package database

import (
	"blogo-server/api/domain"
	"time"
)

// PostRepository - sql handler
type PostRepository struct {
	SQLHandler
}

//StorePost - insert post
func (repo *PostRepository) StorePost(p domain.Post) (id int, err error) {
	result, err := repo.Execute(
		"INSERT INTO posts (title, content, created_at) VALUES (?,?,?)", p.Title, p.Content, time.Now(),
	)
	if err != nil {
		print("ww")
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		print("ww")
		return
	}
	id = int(id64)
	return
}

// FindAll - select all
func (repo *PostRepository) FindAll() (posts domain.Posts, err error) {
	rows, err := repo.Query("SELECT id, title, content FROM posts")
	defer rows.Close()
	if err != nil {
		print("ww")
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

// FindByID - select post
func (repo *PostRepository) FindByID(identifier int) (post domain.Post, err error) {
	row, err := repo.Query("SELECT id, title, content FROM posts WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		print("ww2")
		return
	}
	var id int64
	var title string
	var content string
	row.Next()
	if err = row.Scan(&id, &title, &content); err != nil {
		print("ww3")
		return
	}
	post.ID = id
	post.Title = title
	post.Content = content
	return
}

// DeletePost - remove post
func (repo *PostRepository) DeletePost(identifier int) (err error) {
	_, err = repo.Execute("DELETE FROM posts WHERE id = ?", identifier)
	if err != nil {
		print("ww")
	}
	return
}
