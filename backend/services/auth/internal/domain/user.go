package domain

import "time"

const DefaultRole = "Student"

type User struct {
	ID           string
	Name         string
	Username     string
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

type TokenPair struct {
	AccessToken           string
	AccessTokenExpiresAt  time.Time
	RefreshToken          string
	RefreshTokenExpiresAt time.Time
}

type PublicUser struct {
	ID        string
	Name      string
	Username  string
	Email     string
	Role      string
	CreatedAt time.Time
}

func (u User) Public() PublicUser {
	return PublicUser{
		ID:        u.ID,
		Name:      u.Name,
		Username:  u.Username,
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
	}
}
