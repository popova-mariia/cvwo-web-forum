package models

import "time"

type Post struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	ThreadID  int       `json:"thread_id"`
	AuthorID  int       `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
}
