package httpadapter

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	router := NewRouter(slog.New(slog.NewTextHandler(io.Discard, nil)))

	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}

	if response.Header().Get(requestIDHeader) == "" {
		t.Fatal("expected request id header")
	}
}

func TestRequestIDMiddlewareKeepsIncomingRequestID(t *testing.T) {
	router := NewRouter(slog.New(slog.NewTextHandler(io.Discard, nil)))

	request := httptest.NewRequest(http.MethodGet, "/ready", nil)
	request.Header.Set(requestIDHeader, "test-request-id")
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Header().Get(requestIDHeader) != "test-request-id" {
		t.Fatalf("expected incoming request id to be preserved, got %q", response.Header().Get(requestIDHeader))
	}
}

func TestCORSMiddlewareHandlesPreflight(t *testing.T) {
	router := NewRouter(slog.New(slog.NewTextHandler(io.Discard, nil)))

	request := httptest.NewRequest(http.MethodOptions, "/health", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusNoContent {
		t.Fatalf("expected status 204, got %d", response.Code)
	}

	if response.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Fatal("expected CORS allow origin header")
	}
}
