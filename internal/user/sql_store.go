package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type SQLStore struct {
	connection *sql.DB
}

func NewSQLStore(pool *sql.DB) SQLStore {
	return SQLStore{connection: pool}
}

func (s SQLStore) Create(ctx context.Context, user User) (User, error) {
	return User{}, nil
}

func (s SQLStore) Read(ctx context.Context, userID ID) (User, error) {
	// Insert SELECT statement here
	// row, err := sql.QuerySelect()
	var err error

	u := User{
		ID: userID,
		Registration: Registration{
			Email: "test@example.com",
			Name:  "Name Name",
		},
	}

	// when doing a real SELECT, the returned error may be a sql.ErrNoRows
	// which should be translated to the service definition's general ErrNotFound.
	if errors.Is(err, sql.ErrNoRows) {
		return User{}, ErrNotFound
	}
	if err != nil {
		return User{}, fmt.Errorf("lookup user from store: %w", err)
	}
	return u, nil
}

func (s SQLStore) Update(ctx context.Context, user User) (User, error) {
	return User{}, nil
}

func (s SQLStore) Delete(ctx context.Context, userID ID) error {
	return nil
}
