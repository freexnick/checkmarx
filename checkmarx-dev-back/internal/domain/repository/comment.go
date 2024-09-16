package repository

import (
	"checkmarx/internal/domain/entity"
)

type CommentRepository interface {
	Insert(comment *entity.Comment) error
	Update(comment *entity.Comment) error
	Delete(id int) error
}
