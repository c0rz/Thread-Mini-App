package models

import "time"

type Post struct {
	ID          uint64    `json:"id,omitempty"`
	Title       string    `json:"title"`
	ID_Category string    `json:"category"`
	Text        string    `json:"text"`
	ID_User     string    `json:"author,omitempty"`
	Comment     []Comment `json:"comment,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}
