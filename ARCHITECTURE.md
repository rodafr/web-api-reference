# Architecture

This project's architecture is a pragmatic, simplified take on ports-and-adapters
(hexagonal) architecture, combined with vertical slices: each domain concern
(`user`, `order`, ...) lives in its own package, containing everything needed
to serve that domain end to end.

## Layers

Each domain package has three layers, always in the same three files:

### Handler

 **i.e. `http_handler.go`**
 Translates wire format to domain and back.
 Parses and validates incoming requests, maps domain values to HTTP status
 codes, marshals responses.
 Knows nothing about how a `User` is fetched or what business rules apply,
 only how to move it in and out over HTTP.

### Service

 **`service.go`**
 Pure domain/business logic. No I/O of its own; it only
 orchestrates calls to an interface it owns.
 This is where the important logic lives, and it's kept supremely testable:
 a hand-written fake satisfying its interface is enough, no mocking framework,
 no database.

### Store

**i.e. `sql_store.go` or `mongo_store.go`**
Translates domain language to infrastructure language
and talks to the storage provider directly (e.g. `SQLStore` runs SQL against
Postgres).
This is also where infra-specific errors are translated into the
package's domain-level sentinel errors (see below).

### Dependency direction

Dependencies point inward on both sides: `Handler` depends on a small
interface it defines for what it needs from `Service`; `Service` depends on
a `storer` interface it defines for what it needs from the world. Concrete
types (`SQLStore`, `HTTPHandler`) satisfy these interfaces structurally;
`Service` never imports `database/sql`, `net/http`, or knows either exists.

## Interface ownership

Interfaces are always defined by the consumer, next to the type that uses
them, not bundled into a shared file by "kind." `storer` lives in
`service.go` because only `Service` depends on it; a handler-facing
interface lives in `http_handler.go` for the same reason.
This keeps each side's tests minimal: a handler test only needs a fake with the
one method it actually calls.

Interfaces start collapsed to a single `storer`-style contract.
Segregating it further into single-verb interfaces (`reader`, `creator`,
...) is a deliberate later step, taken only when something actually needs
to diverge per verb; CQRS (reads/writes on different backends), a
cache-aside reader, permission boundaries. Splitting ahead of that need is
cheap but not free; do it when a second implementation shows up, not before.

## Models

- **Domain model** (`models.go`): `User`, `ID`. Validated, trusted, shared
  across all three layers.
- **Wire model** (`http_handler.go`): `CreateRequest`, `ReadRequest`.
  Unvalidated, HTTP-shaped, exists only to get bytes off the network into
  something Go can hold. Lives with the handler that constructs it, not with
  the domain model — it should never leak past the handler boundary.

## Errors

Sentinel errors (`ErrNotFound`, `ErrConflict`, ...) are defined in
`errors.go`, owned by the domain package. Every adapter is responsible for
translating its own failure mode into these sentinels at the point the
original error occurs (e.g. `SQLStore.Read` turns `sql.ErrNoRows` into
`ErrNotFound`). Nothing above the store layer should ever see a driver-
specific error.

## Wiring

Each domain package is assembled in `internal/api/setup.go`:
store → service → handler, three lines. Routes are registered onto a shared
`*http.ServeMux` via a `setupXRoutes(mux, handler)` function per domain —
side-effecting registration, matching the standard library's own
`mux.Handle` idiom.

## Testing

- `Service` is tested against a fake `storer`; no database, no mocking
  framework.
- `Handler` is tested against a fake satisfying its own small interface,
  not a real `Service`, so handler tests fail only for handler reasons
  (wrong status code, wrong content type, malformed body), and service
  tests fail only for service reasons.
