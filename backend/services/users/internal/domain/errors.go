package domain

import "errors"

var (
	ErrInvalidEmail = errors.New("invalid email")
	ErrInvalidName = errors.New("invalid name")
	ErrInvalidUsername = errors.New("invalid username")
	ErrInvalidRole = errors.New("invalid role")
	ErrInvalidID = errors.New("invalid id")
)