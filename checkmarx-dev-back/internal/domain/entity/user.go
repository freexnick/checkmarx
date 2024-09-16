package entity

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Credentials struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	Token  Token  `json:"session"`
}
