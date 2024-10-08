package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"checkmarx/internal/domain/entity"
)

type AuthRepository struct {
	db *Client
}

func NewAuthRepository(db *Client) *AuthRepository {
	return &AuthRepository{db: db}
}

func (ar *AuthRepository) Get(email string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	user := new(entity.User)

	query := `SELECT id, email, password, created_at, updated_at FROM users WHERE email = $1`
	if err := ar.db.client.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		ar.db.observ.Error(ctx, err)
		if err == sql.ErrNoRows {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return user, nil
}

func (ar *AuthRepository) Insert(u *entity.User, t *entity.Token) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var user entity.User

	args := []any{u.Email, u.Password, t.Scope, t.Hash, t.ExpiresAt}
	query :=
		`WITH inserted_user AS (
			INSERT INTO users (email, password, created_at, updated_at)
			VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
			RETURNING id, email, created_at, updated_at
		)
		INSERT INTO tokens (user_id, scope, hash, created_at, updated_at, expires_at)
		SELECT id, $3, $4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, $5
		FROM inserted_user
		RETURNING id, (SELECT email FROM inserted_user), created_at, updated_at;
		`

	err := ar.db.client.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		ar.db.observ.Error(ctx, err)
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return nil, ErrDuplicateEmail
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrEditConflict
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (ar *AuthRepository) InsertToken(t *entity.Token) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	args := []any{t.Hash, t.UserID, t.Scope, t.ExpiresAt}
	query := `
		INSERT INTO tokens (hash, user_id, scope, expires_at)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (user_id)
		DO UPDATE SET
			hash = EXCLUDED.hash,
			expires_at = EXCLUDED.expires_at,
			updated_at = CURRENT_TIMESTAMP
	`

	_, err := ar.db.client.ExecContext(ctx, query, args...)
	if err != nil {
		ar.db.observ.Error(ctx, err)
		return err
	}

	return nil
}

func (ar *AuthRepository) GetByToken(t [32]byte) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var user entity.User
	query := `
		SELECT u.*
		FROM users u
		JOIN tokens t ON u.id = t.user_id
		WHERE t.hash = $1
		AND t.expires_at > CURRENT_TIMESTAMP
	`

	err := ar.db.client.QueryRowContext(ctx, query, t[:]).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		ar.db.observ.Error(ctx, err)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
