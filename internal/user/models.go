package user

import "uuid"

// User is the registered and validated user info
// i.e domain model
type User struct {
	ID    ID
	Email string // TODO: valid.Email type
	Name  string
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
