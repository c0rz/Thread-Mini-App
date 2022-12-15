package models

import "time"

type Comment struct {
	ID        uint64    `json:"id,omitempty"`
	ID_Post   uint64    `json:"id_post,omitempty"`
	ID_User   string    `json:"comment_by"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}
