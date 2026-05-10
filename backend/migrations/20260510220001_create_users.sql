-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    role TEXT NOT NULL DEFAULT 'Student',
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);

CREATE UNIQUE INDEX users_email_unique_idx ON users (LOWER(email));
CREATE UNIQUE INDEX users_username_unique_idx ON users (LOWER(username));

-- +goose Down
DROP INDEX IF EXISTS users_username_unique_idx;
DROP INDEX IF EXISTS users_email_unique_idx;
DROP TABLE IF EXISTS users;
