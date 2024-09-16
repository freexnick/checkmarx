package service

import (
	"checkmarx/internal/domain/entity"
	"checkmarx/internal/domain/repository"
)

type CommentService struct {
	repo repository.CommentRepository
}

func NewCommentService(repo repository.CommentRepository) *CommentService {
	return &CommentService{repo: repo}
}

func (cs *CommentService) CreateComment(comment *entity.Comment) error {
	return cs.repo.Insert(comment)
}

func (cs *CommentService) UpdateComment(comment *entity.Comment) error {
	return cs.repo.Update(comment)
}

func (cs *CommentService) DeleteComment(id int) error {
	return cs.repo.Delete(id)
}
