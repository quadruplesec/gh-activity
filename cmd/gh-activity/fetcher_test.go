package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFetchUsers(t *testing.T) {
	mockJSON := `[
		{
			"type": "PushEvent",
			"repo": {
				"name": "test/repo"
			}
		},
		{
			"type": "PushEvent",
			"repo": {
				"name": "test/repo"
			}
		}
	]`

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, mockJSON)
	}))
	defer mockServer.Close()

	originalBaseURL := apiBaseURL
	apiBaseURL = mockServer.URL
	defer func() { apiBaseURL = originalBaseURL }() // Set it back in the end

	var buf bytes.Buffer
	usernames := []string{"testuser1", "testuser2"}

	FetchUsers(usernames, nil, mockServer.Client(), &buf)

	output := buf.String()

	if !strings.Contains(output, "Recent GitHub Activity of testuser1:") {
		t.Errorf("Expected output for testuser1, got:\n%s", output)
	}

	if !strings.Contains(output, "Recent GitHub Activity of testuser2:") {
		t.Errorf("Expected output for testuser2, got:\n%s", output)
	}

	if !strings.Contains(output, "Pushed 2 commits to test/repo") {
		t.Errorf("Expected parsed event output, got:\n%s", output)
	}
}