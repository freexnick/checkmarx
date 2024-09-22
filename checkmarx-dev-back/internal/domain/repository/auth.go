package repository

import (
	"checkmarx/internal/domain/entity"
)

type AuthRepository interface {
	Insert(user *entity.User, token *entity.Token) (*entity.User, error)
	InsertToken(user *entity.Token) error
	Get(email string) (*entity.User, error)
	GetByToken(plaintext [32]byte) (*entity.User, error)
}
