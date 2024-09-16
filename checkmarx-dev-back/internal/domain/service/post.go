package service

import (
	"checkmarx/internal/domain/entity"
	"checkmarx/internal/domain/repository"
	"errors"
)

type PostService struct {
	repo repository.PostRepository
}

type PostWithComments struct {
	Post     *entity.Post      `json:"post"`
	Comments []*entity.Comment `json:"comments"`
}

func NewPostService(repo repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (ps *PostService) CreatePost(post *entity.Post) error {
	if post.Title == "" || post.Content == "" || post.AuthorID == 0 {
		return errors.New("missing fields")
	}

	return ps.repo.Insert(post)
}

func (ps *PostService) GetPost(id int) (*PostWithComments, error) {
	post, comments, err := ps.repo.Get(id)
	if err != nil {
		return nil, err
	}

	return &PostWithComments{
		Post:     post,
		Comments: comments,
	}, nil
}

func (ps *PostService) GetAll() ([]*entity.Post, error) {
	return ps.repo.GetAll()
}

func (ps *PostService) Update(post *entity.Post) error {
	return ps.repo.Update(post)
}

func (ps *PostService) Delete(id int) error {
	return ps.repo.Delete(id)
}
