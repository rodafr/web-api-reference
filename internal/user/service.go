package user

import (
	"context"
	"fmt"
)

// storer defines Service's dependencies, i.e. what it needs from the world
type storer interface {
	Create(context.Context, User) (User, error)
	Read(context.Context, ID) (User, error)
	Update(context.Context, User) (User, error)
	Delete(context.Context, ID) error
}

// Service consumes a storer
type Service struct {
	storer storer
}

// NewService creates an instance of a Service, using a storer
func NewService(s storer) Service {
	return Service{
		storer: s,
	}
}

//

// Lookup looks up a user from the store based on a given UUID
// i.e. uses a GET request to READ a user
// It received an already validated User struct and can focus on
// pure business logic pertaining to looking up/searching for a
// user in the store.
func (s Service) Lookup(ctx context.Context, id ID) (User, error) {
	u, err := s.storer.Read(ctx, id)
	if err != nil {
		return User{}, fmt.Errorf("read user from store: %w", err)
	}

	return u, nil
}

// func Update

// func Delete

func (s Service) Healthcheck(ctx context.Context) error {
	return nil
}
