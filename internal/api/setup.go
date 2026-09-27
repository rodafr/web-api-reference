package api

import (
	"database/sql"
	"net/http"

	"github.com/rodafr/web-api-reference/internal/health"
	"github.com/rodafr/web-api-reference/internal/user"
)

// SetupServerHandler wires up infra, services and routes
// instantiate user service, with a store and a handler
//
// userStore here is a concrete SQL database store, but it could be
// MongoDB, file, etc.
//
// userHandler here is a concrete HTTP handler, but it could be
// gRPC, etc.
func SetupServerHandler(pool *sql.DB) (http.Handler, error) {
	userStore := user.NewSQLStore(pool)
	userService := user.NewService(userStore)
	userHandler := user.NewHTTPHandler(userService)

	// orderStore := order.NewSQLStore(pool)
	// orderService := order.NewService(orderStore)
	// orderHandler := order.NewHTTPHandler(orderService)

	// productStore := product.NewSQLStore(pool)
	// productService := product.NewService(productStore)
	// productHandler := product.NewHTTPHandler(productService)

	// wire routes and handlers
	mux := http.NewServeMux()

	// mux.Handle("POST /users", http.HandlerFunc(userHandler.Post()))
	mux.Handle("GET /users/{id}", http.HandlerFunc(userHandler.Get()))
	// mux.Handle("PATCH /users/{id}", http.HandlerFunc(userHandler.Patch()))
	// mux.Handle("DELETE /users/{id}", http.HandlerFunc(userHandler.Delete()))

	// mux.Handle("POST /orders", http.HandlerFunc(orderHandler.Post()))
	// mux.Handle("GET /orders/{id}", http.HandlerFunc(orderHandler.Get()))
	// mux.Handle("PATCH /orders/{id}", http.HandlerFunc(orderHandler.Patch()))
	// mux.Handle("DELETE /orders/{id}", http.HandlerFunc(orderHandler.Delete()))

	// mux.Handle("POST /products", http.HandlerFunc(productHandler.Post()))
	// mux.Handle("GET /products/{id}", http.HandlerFunc(productHandler.Get()))
	// mux.Handle("PATCH /products/{id}", http.HandlerFunc(productHandler.Patch()))
	// mux.Handle("DELETE /products/{id}", http.HandlerFunc(productHandler.Delete()))

	healthCheckHandler := health.Healthz(
		health.HealthCheck{
			Service:         "user-service",
			HealthCheckFunc: userService.Healthcheck,
			Status:          "",
		},
		// health.HealthCheck{
		// 	Service:         "order-service",
		// 	HealthCheckFunc: orderService.Healthcheck,
		// 	Status:          "",
		// },
		// health.HealthCheck{
		// 	Service:         "product-service",
		// 	HealthCheckFunc: productService.Healthcheck,
		// 	Status:          "",
		// },
	)

	mux.Handle("/healthz", healthCheckHandler)

	return mux, nil
}
