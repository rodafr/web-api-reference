package api

import (
	"net/http"

	"github.com/rodafr/web-api-reference/internal/user"
)

func setupUserRoutes(mux *http.ServeMux, h user.HTTPHandler) {
	// mux.Handle("POST /users", userHandler.Post())
	mux.Handle("GET /users/{id}", h.Get())
	// mux.Handle("PATCH /users/{id}", userHandler.Patch())
	// mux.Handle("DELETE /users/{id}", userHandler.Delete())
}

//func setupOrderRoutes(mux *http.ServeMux, h order.HTTPHandler) {
// mux.Handle("POST /orders", orderHandler.Post())
// mux.Handle("GET /orders/{id}", orderHandler.Get())
// mux.Handle("PATCH /orders/{id}", orderHandler.Patch())
// mux.Handle("DELETE /orders/{id}", orderHandler.Delete())
//}

//func setupProductRoutes(mux *http.ServeMux, h product.HTTPHandler) {
// mux.Handle("POST /products", productHandler.Post())
// mux.Handle("GET /products/{id}", productHandler.Get())
// mux.Handle("PATCH /products/{id}", productHandler.Patch())
// mux.Handle("DELETE /products/{id}", productHandler.Delete())
//}
