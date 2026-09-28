package user

import (
	"fmt"
	"net/mail"
	"uuid"
)

// User is the registered and validated user info
// i.e domain model
type User struct {
	ID ID
	Registration
}

// Registration is the unvalidated request from a user to register
// i.e. POST request model
type Registration struct {
	Email Email
	Name  string
}

type Email string

func NewEmail(raw string) (Email, error) {
	valid, err := mail.ParseAddress(raw)
	if err != nil {
		return "", fmt.Errorf("parse email: %w", err)
	}

	return Email(valid.Address), nil
}

func (e Email) String() string {
	return string(e)
}

type ID uuid.UUID

func (id ID) MarshalText() ([]byte, error) {
	return uuid.UUID(id).MarshalText()
}

func (id *ID) UnmarshalText(b []byte) error {
	return (*uuid.UUID)(id).UnmarshalText(b)
}

func (id ID) String() string {
	return uuid.UUID(id).String()
}
