-- +goose Up
CREATE TABLE auth_identities (
    id UUID PRIMARY KEY,
    email TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    role TEXT NOT NULL DEFAULT 'Student',
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);

CREATE UNIQUE INDEX auth_identities_email_unique_idx ON auth_identities (LOWER(email));

-- +goose Down
DROP INDEX IF EXISTS auth_identities_email_unique_idx;
DROP TABLE IF EXISTS auth_identities;
