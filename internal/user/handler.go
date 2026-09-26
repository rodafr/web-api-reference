package user

import (
	"encoding/json/v2"
	"fmt"
	"net/http"
	"uuid"
)

// CreateRequest is the unvalidated request from a user to register
// i.e. POST request model
type CreateRequest struct {
	Email string
	Name  string
}

// ReadRequest is the unvalidated request from a user to read
// i.e. GET request model
type ReadRequest struct {
	UserID string
}

type Handler struct {
	svc *Service
}

func NewHandler(s Service) Handler {
	return Handler{svc: &s}
}

func (h Handler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() { _ = r.Body.Close() }()

		ctx := r.Context()

		parsedID, err := uuid.Parse(r.PathValue("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("parse user ID as UUID: %s", err.Error()), http.StatusBadRequest)
			return
		}

		u, err := h.svc.reader.Read(ctx, ID(parsedID))

		err = json.MarshalWrite(w, u)
	}
}
