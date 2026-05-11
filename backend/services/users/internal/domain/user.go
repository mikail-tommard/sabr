package domain

import (
	"strings"
	"time"
)

type Role string

const (
	RoleAdmin   Role = "admin"
	RoleStudent Role = "student"
	RoleMentor  Role = "mentor"
	RoleAlumni  Role = "alumni"
)

type User struct {
	ID        string
	Email     string
	Name      string
	Username  string
	AvatarURL string
	Bio       string
	CampusID  *string
	Company   *string
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
}

type NewUserParams struct {
	ID        string
	Email     string
	Name      string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpdateProfileParams struct {
	Name      string
	Username  string
	AvatarURL string
	Bio       string
	CampusID  *string
	Company   *string
	UpdatedAt time.Time
}

func NewUser(params NewUserParams) (*User, error) {
	u := &User{
		ID:        strings.TrimSpace(params.ID),
		Email:     normalizeEmail(params.Email),
		Name:      strings.TrimSpace(params.Name),
		Username:  normalizeUsername(params.Username),
		Role:      RoleStudent,
		CreatedAt: params.CreatedAt,
		UpdatedAt: params.UpdatedAt,
	}

	if err := u.validate(); err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) UpdateUser(params UpdateProfileParams) error {
	updated := User{
		ID:        u.ID,
		Email:     u.Email,
		Name:      strings.TrimSpace(params.Name),
		Username:  normalizeUsername(params.Username),
		AvatarURL: strings.TrimSpace(params.AvatarURL),
		Bio:       strings.TrimSpace(params.Bio),
		CampusID:  normalizeOptionalString(params.CampusID),
		Company:   normalizeOptionalString(params.Company),
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: params.UpdatedAt,
	}

	if err := updated.validate(); err != nil {
		return err
	}

	*u = updated
	return nil
}

func (u *User) validate() error {
	if u.ID == "" {
		return ErrInvalidID
	}
	if u.Email == "" || !strings.Contains(u.Email, "@") {
		return ErrInvalidEmail
	}
	if len(u.Name) < 2 {
		return ErrInvalidName
	}
	if len(u.Username) < 3 {
		return ErrInvalidUsername
	}
	if !isValidRole(u.Role) {
		return ErrInvalidRole
	}
	return nil
}

func (u *User) ChangeRole(role Role, updatedAt time.Time) error {
	if !isValidRole(role) {
		return ErrInvalidRole
	}

	u.Role = role
	u.UpdatedAt = updatedAt
	return nil
}

func isValidRole(role Role) bool {
	switch role {
	case RoleAdmin, RoleStudent, RoleMentor, RoleAlumni:
		return true
	default:
		return false
	}
}

func normalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

func normalizeUsername(username string) string {
	return strings.ToLower(strings.TrimSpace(username))
}

func normalizeOptionalString(value *string) *string {
	if value == nil {
		return nil
	}

	trimmed := strings.TrimSpace(*value)
	if trimmed == "" {
		return nil
	}

	return &trimmed
}
