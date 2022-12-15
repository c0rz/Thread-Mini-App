package models

import "time"

type Account struct {
	ID        uint64    `json:"id"`
	Full_Name string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	Level     string    `json:"level,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
