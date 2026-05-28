package config

import (
	"testing"
	"time"
)

func TestLoadUsesDefaults(t *testing.T) {
	t.Setenv("APP_PORT", "")
	t.Setenv("WORKER_INTERVAL", "")
	t.Setenv("SHUTDOWN_TIMEOUT", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("load config: %v", err)
	}

	if cfg.App.Port != 8080 {
		t.Fatalf("expected default port 8080, got %d", cfg.App.Port)
	}

	if cfg.Worker.Interval != 30*time.Minute {
		t.Fatalf("expected default worker interval 30m, got %s", cfg.Worker.Interval)
	}
}

func TestLoadParsesEnv(t *testing.T) {
	t.Setenv("APP_ENV", "test")
	t.Setenv("APP_PORT", "9090")
	t.Setenv("LOG_LEVEL", "debug")
	t.Setenv("SHUTDOWN_TIMEOUT", "3s")
	t.Setenv("WORKER_INTERVAL", "1m")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("load config: %v", err)
	}

	if cfg.App.Env != "test" {
		t.Fatalf("expected env test, got %s", cfg.App.Env)
	}

	if cfg.App.Port != 9090 {
		t.Fatalf("expected port 9090, got %d", cfg.App.Port)
	}

	if cfg.Observability.LogLevel != "debug" {
		t.Fatalf("expected log level debug, got %s", cfg.Observability.LogLevel)
	}

	if cfg.App.ShutdownTimeout != 3*time.Second {
		t.Fatalf("expected shutdown timeout 3s, got %s", cfg.App.ShutdownTimeout)
	}
}

func TestLoadReturnsErrorForInvalidDuration(t *testing.T) {
	t.Setenv("WORKER_INTERVAL", "invalid")

	if _, err := Load(); err == nil {
		t.Fatal("expected error for invalid worker interval")
	}
}
