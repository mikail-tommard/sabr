package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"sabr/backend/services/auth/internal/domain"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user domain.User) (domain.User, error) {
	const query = `
		INSERT INTO users (id, name, username, email, password_hash, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, name, username, email, password_hash, role, created_at, updated_at
	`

	var created domain.User
	err := r.db.QueryRow(ctx, query,
		user.ID,
		user.Name,
		user.Username,
		user.Email,
		user.PasswordHash,
		user.Role,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(
		&created.ID,
		&created.Name,
		&created.Username,
		&created.Email,
		&created.PasswordHash,
		&created.Role,
		&created.CreatedAt,
		&created.UpdatedAt,
	)

	return created, err
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	return r.getOne(ctx, `SELECT id, name, username, email, password_hash, role, created_at, updated_at FROM users WHERE email = $1`, email)
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (domain.User, error) {
	return r.getOne(ctx, `SELECT id, name, username, email, password_hash, role, created_at, updated_at FROM users WHERE username = $1`, username)
}

func (r *UserRepository) GetByID(ctx context.Context, userID string) (domain.User, error) {
	return r.getOne(ctx, `SELECT id, name, username, email, password_hash, role, created_at, updated_at FROM users WHERE id = $1`, userID)
}

func (r *UserRepository) getOne(ctx context.Context, query string, arg string) (domain.User, error) {
	var user domain.User
	err := r.db.QueryRow(ctx, query, arg).Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.User{}, domain.ErrUserNotFound
	}

	return user, err
}
