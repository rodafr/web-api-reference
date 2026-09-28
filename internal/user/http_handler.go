package user

import (
	"encoding/json/v2"
	"errors"
	"fmt"
	"net/http"
	"uuid"
)

// ReadRequest is the unvalidated request from a user to read
// i.e. GET request model
type ReadRequest struct {
	UserID string
}

// HTTPHandler consumes a service
type HTTPHandler struct {
	Service
}

// NewHTTPHandler creates an instance of an HTTPHandler, using the service
func NewHTTPHandler(s Service) HTTPHandler {
	return HTTPHandler{Service: s}
}

// Get takes an HTTP GET request with an id (string), parses it as an UUID,
// requests a lookup for that ID from the service, and writes a json encoded
// User object or an appropriate error to w.
func (h HTTPHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		parsedID, err := uuid.Parse(r.PathValue("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("parse user ID as UUID: %s", err.Error()), http.StatusBadRequest)
			return
		}

		u, err := h.Service.Lookup(ctx, ID(parsedID))
		if err != nil {
			switch {
			case errors.Is(err, ErrNotFound):
				http.Error(w, "user not found", http.StatusNotFound)
			default:
				http.Error(w, "user lookup failed", http.StatusInternalServerError)
			}
			return
		}

		js, err := json.Marshal(u)
		if err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func (h HTTPHandler) Post() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "failed to parse form", http.StatusBadRequest)
			return
		}

		unameRaw := r.Form.Get("uname")
		emailRaw := r.Form.Get("email")

		validEmail, err := NewEmail(emailRaw)
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid email: %s", err.Error()), http.StatusBadRequest)
			return
		}

		// TODO: validate username
		validUname := unameRaw

		cr := Registration{
			Email: validEmail,
			Name:  validUname,
		}

		u, err := h.Service.Register(ctx, cr)
		if err != nil {
			switch {
			case errors.Is(err, ErrConflictUname):
				http.Error(w, fmt.Sprintf("user already registered: %q", validUname), http.StatusUnprocessableEntity)
			case errors.Is(err, ErrConflictEmail):
				http.Error(w, fmt.Sprintf("email already registered: %q", validEmail), http.StatusUnprocessableEntity)
			default:
				http.Error(w, "failed to register user", http.StatusInternalServerError)
			}
			return
		}

		js, err := json.Marshal(u)
		if err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}
