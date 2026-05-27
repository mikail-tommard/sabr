package domain

import "time"

type Role string

const (
	RoleAdmin   Role = "admin"
	RoleStudent Role = "student"
	RoleMentor  Role = "mentor"
	RoleAlumni  Role = "alumni"
	DefaultRole      = RoleStudent
)

type Identity struct {
	ID           string
	Email        string
	PasswordHash string
	Role         Role
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
