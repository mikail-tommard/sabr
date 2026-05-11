package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"sabr/backend/services/auth/internal/domain"
)

type IdentityRepository struct {
	db *pgxpool.Pool
}

func NewIdentityRepository(db *pgxpool.Pool) *IdentityRepository {
	return &IdentityRepository{db: db}
}

func (r *IdentityRepository) Create(ctx context.Context, identity domain.Identity) (domain.Identity, error) {
	const query = `
		INSERT INTO auth_identities (id, email, password_hash, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, email, password_hash, role, created_at, updated_at
	`

	var created domain.Identity
	err := r.db.QueryRow(ctx, query,
		identity.ID,
		identity.Email,
		identity.PasswordHash,
		identity.Role,
		identity.CreatedAt,
		identity.UpdatedAt,
	).Scan(
		&created.ID,
		&created.Email,
		&created.PasswordHash,
		&created.Role,
		&created.CreatedAt,
		&created.UpdatedAt,
	)

	return created, err
}

func (r *IdentityRepository) GetByEmail(ctx context.Context, email string) (domain.Identity, error) {
	return r.getOne(ctx, `SELECT id, email, password_hash, role, created_at, updated_at FROM auth_identities WHERE email = $1`, email)
}

func (r *IdentityRepository) GetByID(ctx context.Context, userID string) (domain.Identity, error) {
	return r.getOne(ctx, `SELECT id, email, password_hash, role, created_at, updated_at FROM auth_identities WHERE id = $1`, userID)
}

func (r *IdentityRepository) getOne(ctx context.Context, query string, arg string) (domain.Identity, error) {
	var identity domain.Identity
	err := r.db.QueryRow(ctx, query, arg).Scan(
		&identity.ID,
		&identity.Email,
		&identity.PasswordHash,
		&identity.Role,
		&identity.CreatedAt,
		&identity.UpdatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.Identity{}, domain.ErrUserNotFound
	}

	return identity, err
}
