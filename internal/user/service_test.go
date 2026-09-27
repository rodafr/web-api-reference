package user

import (
	"context"
	"testing"
	"uuid"
)

type MockedStore struct{}

func NewMockedStore() MockedStore {
	return MockedStore{}
}

func (ms MockedStore) Create(ctx context.Context, user User) (User, error) {
	return User{}, nil
}

func (ms MockedStore) Read(ctx context.Context, userID ID) (User, error) {
	return User{
		ID:    userID,
		Email: "test@example.com",
		Name:  "Name Name",
	}, nil
}

func (ms MockedStore) Update(ctx context.Context, user User) (User, error) {
	return User{}, nil
}

func (ms MockedStore) Delete(ctx context.Context, userID ID) error {
	return nil
}

func TestLookup(t *testing.T) {
	svc := NewService(NewMockedStore())

	ctx := context.TODO()
	userID := ID(uuid.New())

	u, err := svc.Lookup(ctx, userID)
	if err != nil {
		t.Fatal(err)
	}

	if u.ID != userID {
		t.Fatal(err)
	}
}
