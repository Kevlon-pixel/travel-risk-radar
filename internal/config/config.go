package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	App           AppConfig
	Database      DatabaseConfig
	Auth          AuthConfig
	Weather       WeatherConfig
	Worker        WorkerConfig
	Observability ObservabilityConfig
}

type AppConfig struct {
	Env             string
	Port            int
	ShutdownTimeout time.Duration
}

type DatabaseConfig struct {
	URL string
}

type AuthConfig struct {
	JWTSecret string
}

type WeatherConfig struct {
	APIKey string
}

type WorkerConfig struct {
	Interval time.Duration
}

type ObservabilityConfig struct {
	LogLevel string
}

func Load() (Config, error) {
	port, err := getEnvInt("APP_PORT", 8080)
	if err != nil {
		return Config{}, err
	}

	shutdownTimeout, err := getEnvDuration("SHUTDOWN_TIMEOUT", 10*time.Second)
	if err != nil {
		return Config{}, err
	}

	workerInterval, err := getEnvDuration("WORKER_INTERVAL", 30*time.Minute)
	if err != nil {
		return Config{}, err
	}

	return Config{
		App: AppConfig{
			Env:             getEnv("APP_ENV", "local"),
			Port:            port,
			ShutdownTimeout: shutdownTimeout,
		},
		Database: DatabaseConfig{
			URL: getEnv("DATABASE_URL", ""),
		},
		Auth: AuthConfig{
			JWTSecret: getEnv("JWT_SECRET", ""),
		},
		Weather: WeatherConfig{
			APIKey: getEnv("WEATHER_API_KEY", ""),
		},
		Worker: WorkerConfig{
			Interval: workerInterval,
		},
		Observability: ObservabilityConfig{
			LogLevel: getEnv("LOG_LEVEL", "info"),
		},
	}, nil
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}

func getEnvInt(key string, fallback int) (int, error) {
	value := os.Getenv(key)
	if value == "" {
		return fallback, nil
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("parse %s as int: %w", key, err)
	}

	return parsed, nil
}

func getEnvDuration(key string, fallback time.Duration) (time.Duration, error) {
	value := os.Getenv(key)
	if value == "" {
		return fallback, nil
	}

	parsed, err := time.ParseDuration(value)
	if err != nil {
		return 0, fmt.Errorf("parse %s as duration: %w", key, err)
	}

	return parsed, nil
}
