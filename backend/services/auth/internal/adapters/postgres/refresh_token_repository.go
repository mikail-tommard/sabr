package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"sabr/backend/services/auth/internal/domain"
)

type RefreshTokenRepository struct {
	db *pgxpool.Pool
}

func NewRefreshTokenRepository(db *pgxpool.Pool) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}

func (r *RefreshTokenRepository) Create(ctx context.Context, token domain.RefreshToken) error {
	const query = `
		INSERT INTO refresh_tokens (id, user_id, token_hash, expires_at, created_at, revoked_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.Exec(ctx, query,
		token.ID,
		token.UserID,
		token.TokenHash,
		token.ExpiresAt,
		token.CreatedAt,
		token.RevokedAt,
	)

	return err
}

func (r *RefreshTokenRepository) GetByHash(ctx context.Context, tokenHash string) (domain.RefreshToken, error) {
	const query = `
		SELECT id, user_id, token_hash, expires_at, created_at, revoked_at
		FROM refresh_tokens
		WHERE token_hash = $1
	`

	var token domain.RefreshToken
	err := r.db.QueryRow(ctx, query, tokenHash).Scan(
		&token.ID,
		&token.UserID,
		&token.TokenHash,
		&token.ExpiresAt,
		&token.CreatedAt,
		&token.RevokedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.RefreshToken{}, domain.ErrInvalidRefreshToken
	}

	return token, err
}

func (r *RefreshTokenRepository) Revoke(ctx context.Context, tokenID string, revokedAt time.Time) error {
	_, err := r.db.Exec(ctx, `UPDATE refresh_tokens SET revoked_at = $2 WHERE id = $1`, tokenID, revokedAt)
	return err
}
