package domain

import (
	"time"
)

// Posts - post slice
type Posts []Post

// Post - blog article
type Post struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	Author    Author    `json:"author"`
	Image     string    `json:"image"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
