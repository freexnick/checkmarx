package repository

import "checkmarx/internal/domain/entity"

type UserRepository interface {
	Get(email string) (*entity.User, error)
}
