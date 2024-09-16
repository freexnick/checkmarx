package postgres

import (
	"context"
	"database/sql"
	"time"

	"checkmarx/internal/domain/entity"
)

type UserRepository struct {
	db *Client
}

func NewUserRepository(db *Client) *UserRepository {
	return &UserRepository{db: db}
}

func (pr *UserRepository) Get(email string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	user := new(entity.User)

	query := `SELECT id, email, password, created_at, updated_at FROM users WHERE email = $1`

	err := pr.db.client.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return user, nil
}
