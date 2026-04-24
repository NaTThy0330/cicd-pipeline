package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("returns api key when authorization header is valid", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Authorization", "ApiKey test-api-key")

		got, err := GetAPIKey(headers)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if got != "test-api-key" {
			t.Fatalf("expected api key %q, got %q", "test-api-key", got)
		}
	})

	t.Run("returns error when authorization header is missing", func(t *testing.T) {
		got, err := GetAPIKey(http.Header{})
		if !errors.Is(err, ErrNoAuthHeaderIncluded) {
			t.Fatalf("expected ErrNoAuthHeaderIncluded, got %v", err)
		}
		if got != "" {
			t.Fatalf("expected empty api key, got %q", got)
		}
	})
}
