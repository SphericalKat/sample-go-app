package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sample/api/gen"
	"sample/pkg/di"
	"time"

	"log/slog"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	ctx := context.Background()
	serverImpl, err := di.InjectServer()
	if err != nil {
		slog.Error("error creating server impl", "err", err)
		os.Exit(1)
	}
	serverInterface := gen.NewStrictHandler(serverImpl, nil)
	gen.RegisterHandlersWithBaseURL(e, serverInterface, "/api/v1")

	// https://echo.labstack.com/cookbook/graceful-shutdown/
	// Start server
	go func() {
		slog.Info("Starting HTTP server", "port", "3000")
		if err := e.Start(fmt.Sprintf("0.0.0.0:%d", 3000)); err != nil && err != http.ErrServerClosed {
			slog.Error("shutting down the server", "err", err)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		slog.Error("shutdown err", "err", err)
		os.Exit(1)
	}
}
