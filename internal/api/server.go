package api

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

var (
	Version string
	Tracer  trace.Tracer
	Meter   metric.Meter
)

// Run initializes the web api server
func Run(ctx context.Context) error {
	fmt.Printf("starting API %s ...\n", Version)

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	serverconfig, err := LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load server config: %w", err)
	}

	dbpool := &sql.DB{}

	httpHandler, err := SetupServerHandler(dbpool)
	if err != nil {
		return fmt.Errorf("failed to set up server handler: %w", err)
	}

	httpServer := http.Server{
		Addr:    net.JoinHostPort("", serverconfig.Port),
		Handler: httpHandler,
	}
	slog.Debug("configured http server")

	go func() {
		slog.Info("http server listening", "port", httpServer.Addr, "service", serverconfig.ServiceName)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Go(func() {
		<-ctx.Done()
		slog.Info("received shutdown signal ...")
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	})
	wg.Wait()

	slog.Info("server stopped gracefully", "service", serverconfig.ServiceName)

	return nil
}
