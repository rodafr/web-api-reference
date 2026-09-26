package user

import (
	"context"
	"database/sql"
)

type UserDB struct {
	userDB *sql.DB
}

func NewUserDB(db *sql.DB) UserDB {
	return UserDB{userDB: db}
}

func (udb UserDB) Insert(ctx context.Context, user User) error {
	return nil
}

func (udb UserDB) Update(ctx context.Context, user User) (User, error) {
	return User{}, nil
}

func (udb UserDB) Delete(ctx context.Context, userID ID) error {
	return nil
}

func (udb UserDB) Lookup(ctx context.Context, userID ID) (User, error) {
	return User{
		ID:    userID,
		Email: "test@example.com",
		Name:  "Name Name",
	}, nil
}
