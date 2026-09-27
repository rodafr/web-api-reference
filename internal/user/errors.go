package user

import "errors"

var (
	ErrNotFound      = errors.New("user not found")
	ErrConflictUname = errors.New("username already exists")
	ErrConflictEmail = errors.New("email already exists")
	ErrDeleted       = errors.New("user already deleted")
	ErrInvalidEmail  = errors.New("invalid email address")
)
