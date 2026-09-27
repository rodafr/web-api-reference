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
func SetupServerHandler(pool *sql.DB) http.Handler {
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

	setupUserRoutes(mux, userHandler)
	// setupOrderRoutes(mux, orderHandler)
	// setupProductRoutes(mux, productHandler)

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

	return mux
}
