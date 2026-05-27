package repository

import (
	"context"
	"errors"
	"sabr/backend/services/users/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Close() {
	r.db.Close()
}

func (r *Repository) Create(ctx context.Context, user domain.User) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO users (id, username, email, name, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, user.ID, user.Username, user.Email, user.Name, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (domain.User, error) {
	const query = `
		SELECT id, email, name, username, avatar_url, bio, campus_id, company, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var u domain.User
	err := r.db.QueryRow(ctx, query, id).Scan(
		&u.ID,
		&u.Email,
		&u.Name,
		&u.Username,
		&u.AvatarURL,
		&u.Bio,
		&u.CampusID,
		&u.Company,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.User{}, domain.ErrUserNotFound
	}
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}
