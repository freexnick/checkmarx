package postgres

import (
	"context"
	"errors"
	"time"

	"checkmarx/internal/domain/entity"
)

type CommentRepository struct {
	db *Client
}

func NewCommentRepository(db *Client) *CommentRepository {
	return &CommentRepository{db: db}
}

func (cr *CommentRepository) Insert(c *entity.Comment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	args := []any{c.Content, c.AuthorID, c.PostID}
	query := "INSERT INTO comments (content, author_id, post_id, created_at, updated_at) VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"

	_, err := cr.db.client.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (cr *CommentRepository) Update(c *entity.Comment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	args := []any{c.Content, c.ID}
	query := "UPDATE comments SET content = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2"

	result, err := cr.db.client.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("couldn't find the comment")
	}

	return nil
}

func (cr *CommentRepository) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := "DELETE FROM comments WHERE id = $1"

	result, err := cr.db.client.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("couldn't find the post")
	}

	return nil
}
