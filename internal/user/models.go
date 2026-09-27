package user

import (
	"net/mail"
	"uuid"
)

// User is the registered and validated user info
// i.e domain model
type User struct {
	ID    ID
	Email Email
	Name  string
}

type Email string

func NewEmail(raw string) (Email, error) {
	valid, err := mail.ParseAddress(raw)
	if err != nil {
		return "", ErrInvalidEmail
	}

	return Email(valid.String()), nil
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
