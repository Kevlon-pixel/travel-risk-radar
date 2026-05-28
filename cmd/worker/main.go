package main

import (
	"log/slog"
	"os"
	"time"
)

func main() {
	interval := getenvDuration("WORKER_INTERVAL", 30*time.Minute)

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	slog.Info("starting worker", "interval", interval.String())

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		slog.Info("worker tick")
		<-ticker.C
	}
}

func getenvDuration(key string, fallback time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	duration, err := time.ParseDuration(value)
	if err != nil {
		slog.Warn("invalid duration env, using fallback", "key", key, "value", value, "fallback", fallback.String())
		return fallback
	}

	return duration
}
