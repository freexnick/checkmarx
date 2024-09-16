package entity

import "time"

type Token struct {
	PlainText string    `json:"token"`
	Hash      []byte    `json:"-"`
	UserID    int       `json:"-"`
	ExpiresAt time.Time `json:"expires_at"`
	Scope     string    `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
