package user

import "context"

type fakeStore struct {
	createFn func(context.Context, User) (User, error)
	readFn   func(context.Context, ID) (User, error)
	updateFn func(context.Context, User) (User, error)
	deleteFn func(context.Context, ID) error
}

func (f fakeStore) Create(ctx context.Context, u User) (User, error) { return f.createFn(ctx, u) }
func (f fakeStore) Read(ctx context.Context, id ID) (User, error)    { return f.readFn(ctx, id) }
func (f fakeStore) Update(ctx context.Context, u User) (User, error) { return f.updateFn(ctx, u) }
func (f fakeStore) Delete(ctx context.Context, id ID) error          { return f.deleteFn(ctx, id) }
