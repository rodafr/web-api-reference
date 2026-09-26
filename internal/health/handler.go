package health

import (
	"context"
	"encoding/json/v2"
	"net/http"
)

type HealthCheck struct {
	Service         string                      `json:"service"`
	HealthCheckFunc func(context.Context) error `json:"-"`
	Status          string                      `json:"status"`
}

func Healthz(checks ...HealthCheck) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statuses := []HealthCheck{}
		for _, hc := range checks {
			if err := hc.HealthCheckFunc(context.Background()); err != nil {
				hc.Status = "DOWN"
			} else {
				hc.Status = "OK"
			}
			statuses = append(statuses, hc)
		}

		err := json.MarshalWrite(w, statuses)
		if err != nil {
			// write err
			http.Error(w, "json marshal health check", http.StatusInternalServerError)

			return
		}
	}
}
