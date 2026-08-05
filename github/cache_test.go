package github

import (
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
