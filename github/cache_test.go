package github

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestURLToCacheKey(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Standard URL",
			input:    "https://api.github.com/users/torvalds/events",
			expected: "users_torvalds_events.json",
		},
		{
			name:     "Different Endpoint",
			input:    "https://api.github.com/repos/golang/go/commits",
			expected: "repos_golang_go_commits.json",
		},
		{
			name:     "URL With Trailing Slash",
			input:    "https://api.github.com/users/torvalds/events/",
			expected: "users_torvalds_events.json",
		},
		{
			name:     "Missing HTTPS Prefix",
			input:    "api.github.com/users/torvalds/events",
			expected: "api.github.com_users_torvalds_events.json",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			actual := URLToCacheKey(testCase.input)
			if actual != testCase.expected {
				t.Errorf("Expected %q, got %q", testCase.expected, actual)
			}
		})
	}
}

func TestLoadCache(t *testing.T) {
	tempDir := t.TempDir()
	cacheKey := "test_load.json"
	filePath := filepath.Join(tempDir, cacheKey)

	rawJSON := []byte(`{
		"etag": "\"load-etag\"",
		"activity": [{"type": "PushEvent"}]
	}`)

	err := os.WriteFile(filePath, rawJSON, 0o600)
	if err != nil {
		t.Fatalf("Error when writing to test path: %v", err)
	}

	envelope, err := LoadCache(tempDir, cacheKey)
	if err != nil {
		t.Fatalf("LoadCache failed: %v", err)
	}

	if envelope.ETag != "\"load-etag\"" {
		t.Errorf("Expected ETag %q, got %q", "\"load-etag\"", envelope.ETag)
	}

	if len(envelope.Activity) != 1 || envelope.Activity[0].Type != "PushEvent" {
		t.Errorf("Expected Activity Type %q, got %q", "PushEvent", envelope.Activity[0].Type)
	}
}

func TestSaveCache(t *testing.T) {
	tempDir := t.TempDir()
	cacheKey := "test_save.json"

	envelope := ActivityEnvelope{
		ETag: "save-etag",
	}

	err := SaveCache(tempDir, envelope, cacheKey)
	if err != nil {
		t.Fatalf("SaveCache failed: %v", err)
	}

	filePath := filepath.Join(tempDir, cacheKey)

	file, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf("Error reading cached file.")
	}

	fileString := string(file)
	if !strings.Contains(fileString, `"save-etag"`) {
		t.Errorf("ETag was not correctly saved to cached file.")
	}
}
