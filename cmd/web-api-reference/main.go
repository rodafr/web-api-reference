package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/rodafr/web-api-reference/internal/api"
)

func main() {
	ctx := context.Background()
	if err := api.Run(ctx); err != nil {
		// to be able to log structured logs from here we need to create our own instance of slog.Logger.
		logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
		logger.Error("[FATAL] startup failed", "error", err)
		os.Exit(1)
	}
}
