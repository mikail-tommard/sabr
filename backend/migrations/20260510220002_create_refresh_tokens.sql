-- +goose Up
CREATE TABLE refresh_tokens (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token_hash TEXT NOT NULL UNIQUE,
    expires_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    revoked_at TIMESTAMPTZ NULL
);

CREATE INDEX refresh_tokens_user_id_idx ON refresh_tokens (user_id);
CREATE INDEX refresh_tokens_expires_at_idx ON refresh_tokens (expires_at);

-- +goose Down
DROP INDEX IF EXISTS refresh_tokens_expires_at_idx;
DROP INDEX IF EXISTS refresh_tokens_user_id_idx;
DROP TABLE IF EXISTS refresh_tokens;
