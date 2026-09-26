package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type storer interface {
	Insert(context.Context, User) error
	Update(context.Context, User) (User, error)
	Delete(context.Context, ID) error
	Lookup(context.Context, ID) (User, error)
}

type Store struct {
	storer storer
}

func NewStore(s storer) Store {
	return Store{storer: s}
}

var ErrNotFound = errors.New("user not found")

func (s Store) Create(ctx context.Context, user User) (User, error) {
	return User{}, nil
}

func (s Store) Read(ctx context.Context, userID ID) (User, error) {
	u, err := s.storer.Lookup(ctx, userID)
	if errors.Is(err, sql.ErrNoRows) {
		return User{}, ErrNotFound
	}
	if err != nil {
		return User{}, fmt.Errorf("lookup user from store: %w", err)
	}
	return u, nil
}

func (s Store) Update(ctx context.Context, user User) (User, error) {
	return User{}, nil
}

func (s Store) Delete(ctx context.Context, userID ID) error {
	return nil
}
