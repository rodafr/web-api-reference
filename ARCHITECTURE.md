# Architecture

This project's architecture is an example of a pragmatic and simplistic take on
the ports-and-adapters (or hexagonal) architecture, combined with vertical
slices to contain each distinct domain concern in it's own package. This
provides readability and maintainability.

## Handler

## Service

 comment on composability if storer needs to be split up
 so that it can be satistfied by several implementations, i.e.
 read/write split, cached reads, etc.

```go
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

 type storer interface {
  creator
  reader
  updater
  deleter
 }

 type Service struct {
  creator creator
  reader  reader
  updater updater
  deleter deleter
 }
```

## Store
