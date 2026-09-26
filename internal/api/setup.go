package api

import (
	"database/sql"
	"net/http"

	"github.com/rodafr/web-api-reference/internal/health"
	"github.com/rodafr/web-api-reference/internal/user"
)

// SetupServerHandler wires up infra, services and routes
func SetupServerHandler(pool *sql.DB) (http.Handler, error) {
	userDatabase := user.NewUserDB(pool)
	userStore := user.NewStore(userDatabase)
	userService := user.NewService(userStore)
	userHandler := user.NewHandler(userService)

	healthCheckHandler := health.Healthz(
		health.HealthCheck{
			Service:         "user-service",
			HealthCheckFunc: userService.Healthcheck,
			Status:          "",
		},
	)

	mux := http.NewServeMux()

	mux.Handle("/users/{id}", http.HandlerFunc(userHandler.Get()))

	mux.Handle(
		"/healthz",
		healthCheckHandler,
	)
	return mux, nil
}
