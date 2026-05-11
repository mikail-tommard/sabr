package domain

import "time"

const DefaultRole = "Student"

type Identity struct {
	ID           string
	Email        string
	PasswordHash string
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type RefreshToken struct {
	ID        string
	UserID    string
	TokenHash string
	ExpiresAt time.Time
	CreatedAt time.Time
	RevokedAt *time.Time
}

type UserRegistered struct {
	UserID     string
	Email      string
	Name       string
	Username   string
	OccurredAt time.Time
}
