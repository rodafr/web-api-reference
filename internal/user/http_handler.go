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

type HTTPHandler struct {
	svc *Service
}

func NewHTTPHandler(s Service) HTTPHandler {
	return HTTPHandler{svc: &s}
}

func (h HTTPHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() { _ = r.Body.Close() }()

		ctx := r.Context()

		parsedID, err := uuid.Parse(r.PathValue("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("parse user ID as UUID: %s", err.Error()), http.StatusBadRequest)
			return
		}

		u, err := h.svc.Lookup(ctx, ID(parsedID))

		err = json.MarshalWrite(w, u)
	}
}
