package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/vlady/travel-risk-radar/internal/config"
	"github.com/vlady/travel-risk-radar/internal/observability"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("load config", "error", err)
		os.Exit(1)
	}

	logger := observability.NewLogger(cfg.Observability.LogLevel)
	slog.SetDefault(logger)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	logger.Info("starting worker", "interval", cfg.Worker.Interval.String(), "env", cfg.App.Env)

	if err := run(ctx, logger, cfg); err != nil {
		logger.Error("worker stopped", "error", err)
		os.Exit(1)
	}

	logger.Info("worker stopped")
}
