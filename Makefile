APP_NAME := travel-risk-radar

.PHONY: help run-api run-worker test vet docker-up docker-down docker-build compose-config

help:
	@echo "$(APP_NAME) commands:"
	@echo "  make run-api          Run API locally"
	@echo "  make run-worker       Run worker locally"
	@echo "  make test             Run tests"
	@echo "  make vet              Run go vet"
	@echo "  make docker-build     Build Docker images"
	@echo "  make docker-up        Start services with Docker Compose"
	@echo "  make docker-down      Stop services with Docker Compose"
	@echo "  make compose-config   Validate and render Docker Compose config"

run-api:
	go run ./cmd/api

run-worker:
	go run ./cmd/worker

test:
	go test ./...

vet:
	go vet ./...

docker-build:
	docker compose build

docker-up:
	docker compose up --build

docker-down:
	docker compose down

compose-config:
	docker compose config
