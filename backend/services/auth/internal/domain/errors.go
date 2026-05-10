package domain

import "errors"

var (
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrEmailTaken          = errors.New("email is already taken")
	ErrUsernameTaken       = errors.New("username is already taken")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
	ErrUserNotFound        = errors.New("user not found")
)
