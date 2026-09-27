package user

import "errors"

// the storer is expected to use these sentinel errors:
var (
	ErrNotFound = errors.New("user not found")
	ErrConflict = errors.New("user already exists")
	ErrDeleted  = errors.New("user already deleted")
)
