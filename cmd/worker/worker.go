package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/vlady/travel-risk-radar/internal/config"
)

func run(ctx context.Context, logger *slog.Logger, cfg config.Config) error {
	ticker := time.NewTicker(cfg.Worker.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			logger.Info("worker tick")
		}
	}
}
