package user

import (
	"context"
	"fmt"
)

type creator interface {
	Create(context.Context, User) (User, error)
}

type reader interface {
	Read(context.Context, ID) (User, error)
}

type updater interface {
	Update(context.Context, User) (User, error)
}

type deleter interface {
	Delete(context.Context, ID) error
}

type persister interface {
	creator
	reader
	updater
	deleter
}

type Service struct {
	creator creator
	reader  reader
	deleter deleter
}

func NewService(p persister) Service {
	return Service{
		creator: p,
		reader:  p,
		deleter: p,
	}
}

// func Create

// GetUser looks up a user from the store based on a given UUID
// i.e. uses a GET request to READ a user
func (s Service) Get(ctx context.Context, id ID) (User, error) {
	u, err := s.reader.Read(ctx, id)
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
