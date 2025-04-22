package tests

import (
	"github.com/valivishy/bootdev-cicd/internal/auth"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		expectedKey string
		expectError bool
	}{
		{
			name:        "valid api key",
			headers:     http.Header{"Authorization": []string{"ApiKey abc123"}},
			expectedKey: "abc123",
			expectError: false,
		},
		{
			name:        "missing Authorization header",
			headers:     http.Header{},
			expectedKey: "",
			expectError: true,
		},
		{
			name:        "Authorization header is empty",
			headers:     http.Header{"Authorization": []string{""}},
			expectedKey: "",
			expectError: true,
		},
		{
			name:        "Authorization header with only prefix",
			headers:     http.Header{"Authorization": []string{"ApiKey "}},
			expectedKey: "",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := auth.GetAPIKey(tt.headers)
			if tt.expectError && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if key != tt.expectedKey {
				t.Errorf("expected key %q, got %q", tt.expectedKey, key)
			}
		})
	}
}
