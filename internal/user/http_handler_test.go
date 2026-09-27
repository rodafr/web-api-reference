package user

import (
	"context"
	"encoding/json/v2"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"uuid"
)

type fakeService struct {
	user User
	err  error
}

func (f fakeService) Lookup(ctx context.Context, id ID) (User, error) {
	return f.user, f.err
}

func (f fakeService) Register(ctx context.Context, createReq CreateRequest) (User, error) {
	return f.user, f.err
}

// TestGet checks that the GET HTTP handler behaves as expected:
// - Content-Type should be "application/json" when returning a User
// - Content-Type shoudl NOT be "application/json" when returning an error
// - a valid UUID should parse correctly
// - an invalid UUID should return a Bad Request status
// - a valid UUID should return a valid json representation of a User when successful
func TestGet(t *testing.T) {
	testCases := []struct {
		desc       string
		req        ReadRequest
		wantErr    error
		wantStatus int
	}{
		{
			desc:       "valid user ID",
			req:        ReadRequest{UserID: uuid.New().String()},
			wantStatus: http.StatusOK,
		},
		{
			desc:       "user ID not found",
			req:        ReadRequest{UserID: uuid.New().String()},
			wantErr:    ErrNotFound,
			wantStatus: http.StatusNotFound,
		},
		{
			desc:       "invalid user ID",
			req:        ReadRequest{UserID: "asdf"},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			parsedID, _ := uuid.Parse(tc.req.UserID)
			want := User{ID: ID(parsedID), Email: "test@example.com", Name: "Name Surname"}

			h := NewHTTPHandler(fakeService{user: want, err: tc.wantErr})

			req := httptest.NewRequest(http.MethodGet, "/users/{id}", nil)
			req.SetPathValue("id", tc.req.UserID)

			rec := httptest.NewRecorder()

			h.Get()(rec, req)

			if rec.Code != tc.wantStatus {
				t.Fatalf("got %d, want %d", rec.Code, tc.wantStatus)
			}

			if tc.wantStatus != http.StatusOK {
				if ct := rec.Header().Get("Content-Type"); ct == "application/json" {
					t.Fatalf("expected a plain-text error body, got Content-Type %q", ct)
				}
				if rec.Body.Len() == 0 {
					t.Fatal("expected a non-empty error body")
				}
				return
			}

			if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
				t.Fatalf("Content-Type: got %q, want application/json", ct)
			}

			var got User
			if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
				t.Fatalf("decode response body: %s", err)
			}
			if got != want {
				t.Fatalf("body: got %+v, want %+v", got, want)
			}
		})
	}
}

func TestPost(t *testing.T) {
	testCases := []struct {
		desc       string
		req        CreateRequest
		wantErr    error
		wantStatus int
	}{
		{
			desc: "valid user name and email",
			req: CreateRequest{
				Email: "charles@bark.ey",
				Name:  "CharlesBarkley",
			},
			wantStatus: http.StatusOK,
		},
		{
			desc: "invalid user email",
			req: CreateRequest{
				Email: "asdf",
				Name:  "Per Jensen",
			},
			wantErr:    ErrInvalidEmail,
			wantStatus: http.StatusUnprocessableEntity,
		},
		{
			desc:       "missing email",
			req:        CreateRequest{Name: "Anon Ymous"},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			want := User{
				ID:    ID(uuid.New()),
				Email: Email(tc.req.Email),
				Name:  tc.req.Name,
			}

			form := url.Values{}
			form.Add("uname", tc.req.Name)
			form.Add("email", tc.req.Email)

			h := NewHTTPHandler(fakeService{user: want, err: nil})

			req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(form.Encode()))

			rec := httptest.NewRecorder()

			h.Post()(rec, req)

			if rec.Code != http.StatusOK {
				t.Fatalf("got %d, want %d", rec.Code, http.StatusOK)
			}

			var got User
			if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
				t.Fatalf("decode response body: %s", err)
			}
			if got != want {
				t.Fatalf("body: got %+v, want %+v", got, want)
			}
		})
	}
}
