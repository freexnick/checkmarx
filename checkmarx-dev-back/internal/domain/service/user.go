package service

import (
	"checkmarx/internal/domain/entity"
	"checkmarx/internal/domain/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (us *UserService) Get(email string) (*entity.User, error) {
	return us.repo.Get(email)
}
