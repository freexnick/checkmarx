package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"checkmarx/internal/domain/entity"
)

type PostRepository struct {
	db *Client
}

func NewPostRepository(db *Client) *PostRepository {
	return &PostRepository{db: db}
}

func (pr *PostRepository) Insert(p *entity.Post) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	args := []any{p.Title, p.Content, p.AuthorID}
	query := "INSERT INTO posts (title, content, author_id, created_at) VALUES ($1, $2, $3, CURRENT_TIMESTAMP)"

	_, err := pr.db.client.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (pr *PostRepository) Get(id int) (*entity.Post, []*entity.Comment, error) {
	var post entity.Post
	var comments []*entity.Comment

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	postQuery := `
		SELECT
			id, title, content, author_id, created_at, updated_at
		FROM
			posts
		WHERE
			id = $1
	`

	err := pr.db.client.QueryRowContext(ctx, postQuery, id).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.AuthorID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, ErrRecordNotFound
		}
		return nil, nil, fmt.Errorf("failed to fetch post: %w", err)
	}

	commentsQuery := `
		SELECT
			id, content, author_id, post_id, created_at, updated_at
		FROM
			comments
		WHERE
			post_id = $1
		ORDER BY
			created_at DESC
	`

	rows, err := pr.db.client.QueryContext(ctx, commentsQuery, id)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch comments: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var comment entity.Comment
		err := rows.Scan(
			&comment.ID,
			&comment.Content,
			&comment.AuthorID,
			&comment.PostID,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to scan comment: %w", err)
		}
		comments = append(comments, &comment)
	}

	if err := rows.Err(); err != nil {
		return nil, nil, fmt.Errorf("error iterating comments rows: %w", err)
	}

	return &post, comments, nil
}

func (pr *PostRepository) Update(p *entity.Post) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	args := []any{p.Title, p.Content, p.ID}
	query := "UPDATE posts SET title = $1, content = $2 WHERE id = $3"

	result, err := pr.db.client.ExecContext(ctx, query, args...)
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

func (pr *PostRepository) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := "DELETE FROM posts WHERE id = $1"

	result, err := pr.db.client.ExecContext(ctx, query, id)
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

func (pr *PostRepository) GetAll() ([]*entity.Post, error) {
	var posts []*entity.Post
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := "SELECT id, title, content, author_id, created_at FROM posts ORDER BY created_at DESC"

	rows, err := pr.db.client.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		post := new(entity.Post)
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)

	}
	return posts, nil
}
