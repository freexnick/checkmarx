package repository

import (
	"checkmarx/internal/domain/entity"
)

type PostRepository interface {
	Insert(post *entity.Post) error
	Get(id int) (*entity.Post, []*entity.Comment, error)
	GetAll() ([]*entity.Post, error)
	Update(post *entity.Post) error
	Delete(id int) error
}
