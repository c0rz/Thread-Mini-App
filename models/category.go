package models

import "time"

type Category struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"category_name"`
	CreatedAt time.Time `json:"created_at"`
}
