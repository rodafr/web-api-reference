package user

import (
	"encoding/json/v2"
	"errors"
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
	service Service
}

func NewHTTPHandler(s Service) HTTPHandler {
	return HTTPHandler{service: s}
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

		u, err := h.service.Lookup(ctx, ID(parsedID))
		if err != nil {
			switch {
			case errors.Is(ErrNotFound, err):
				http.Error(w, "user not found", http.StatusNotFound)
			default:
				http.Error(w, "user lookup failed", http.StatusInternalServerError)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.MarshalWrite(w, u)
		if err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
		}
	}
}
