package github_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/quadruplesec/gh-activity/github"
)

func TestMiddleware_CacheMiss(t *testing.T) {
	tempDir := t.TempDir()

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("ETag", `"fake-etag"`)
		w.Header().Set("Cache-Control", "public, max-age=60")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(`[{
			"type": "PushEvent",
			"repo": {
				"name": "test/repo"
			}
		}]`))
	}))

	defer mockServer.Close()

	client := &http.Client{
		Transport: github.NewCachingMiddleware(http.DefaultTransport, tempDir),
	}

	req, err := http.NewRequest("GET", mockServer.URL+"/users/testuser/events", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Network error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	if resp.Header.Get("ETag") != `"fake-etag"` {
		t.Errorf("Expected ETag %q, got %q", `"fake-etag"`, resp.Header.Get("ETag"))
	}
}

func TestMiddleware_FreshCache(t *testing.T) {
	tempDir := t.TempDir()

	// Pre-populate the cache so its fresh
	cacheKey := "users_testuser_events.json"
	envelope := github.ActivityEnvelope{
		ETag:      `"fresh-disk-etag"`,
		FetchedAt: time.Now(),
		TTL:       time.Hour,
		Activity: github.ActivityFeed{
			{
				Type: "IssuesEvent",
			},
		},
	}

	err := github.SaveCache(tempDir, envelope, cacheKey)
	if err != nil {
		t.Fatalf("SaveCache failed: %v", err)
	}

	// To ensure that the server is never actually called (because the cache is fresh), this server throws
	// a fatal error when called.
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatalf("Network hit, middleware should not have made a request to the server when cache was fresh.")
	}))
	defer mockServer.Close()

	client := &http.Client{
		Transport: github.NewCachingMiddleware(http.DefaultTransport, tempDir),
	}

	req, err := http.NewRequest("GET", mockServer.URL+"/users/testuser/events", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Request error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	if resp.Header.Get("ETag") != `"fresh-disk-etag"` {
		t.Errorf("Expected Etag %q, got %q", `"fresh-disk-etag"`, resp.Header.Get("ETag"))
	}
}

func TestMiddleware_StaleCache_NotModified(t *testing.T) {
	tempDir := t.TempDir()

	// Prepopulate a stale cache with data that still has a valid ETag
	cacheKey := "users_testuser_events.json"
	envelope := github.ActivityEnvelope{
		ETag: `"stale-but-valid-etag"`,
		FetchedAt: time.Now().Add(-24 * time.Hour), // Set fetch time to yesterday
		TTL: time.Hour,
		Activity: github.ActivityFeed{
			{
				Type: "PushEvent",
			},
		},
	}

	err := github.SaveCache(tempDir, envelope, cacheKey)
	if err != nil {
		t.Fatalf("SaveCache failed: %v", err)
	}

	// Create a mock server that expects the ETag
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the middleware sent the stored ETag
		if r.Header.Get("If-None-Match") != `"stale-but-valid-etag"` {
			t.Errorf("Expected If-None-Match header %q, got %q", `"stale-but-valid-etag"`, r.Header.Get("If-None-Match"))
		}

		// Respond with 304 Not Modified
		w.WriteHeader(http.StatusNotModified)
	}))

	client := &http.Client{
		Transport: github.NewCachingMiddleware(http.DefaultTransport, tempDir),
	}

	req, err := http.NewRequest("GET", mockServer.URL+"/users/testuser/events", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Request error: %v", err)
	}
	defer resp.Body.Close()

	// Even though the server replies with 304, the middleware should return a 200
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	if resp.Header.Get("ETag") != `"stale-but-valid-etag"` {
		t.Errorf("Expected ETag %q, got %q", `"stale-but-valid-etag"`, resp.Header.Get("ETag"))
	}
}

func TestMiddleware_StaleCache_NewData(t *testing.T) {
	tempDir := t.TempDir()

	// Prepopulate a stale cache with outdated data
	cacheKey := "users_testuser_events.json"
	envelope := github.ActivityEnvelope{
		ETag: `"old-etag"`,
		FetchedAt: time.Now().Add(-24 * time.Hour),
		TTL: time.Hour,
		Activity: github.ActivityFeed{
			{
				Type: "PushEvent",
			},
		},
	}

	err := github.SaveCache(tempDir, envelope, cacheKey)
	if err != nil {
		t.Fatalf("SaveCache failed: %v", err)
	}

	// Create mock server that returns fresh data
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify that the middleware sent the old ETag
		if r.Header.Get("If-None-Match") != `"old-etag"` {
			t.Errorf("Expected If-None-Match %q, got %q", `"old-etag"`, r.Header.Get("If-None-Match"))
		}

		// Return 200 with a new ETag and new data
		w.Header().Set("ETag", `"new-etag"`)
		w.Header().Set("Cache-Control", "public, max-age=60")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(`[{
			"type": "PullRequestEvent",
			"repo": {
				"name": "test/repo"
			}
		}]`))
	}))
	defer mockServer.Close()

	client := &http.Client{
		Transport: github.NewCachingMiddleware(http.DefaultTransport, tempDir),
	}

	req, err := http.NewRequest("GET", mockServer.URL+"/users/testuser/events", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Request error: %v", err)
	}
	defer resp.Body.Close()

	// Check if the status code is 200
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Check if the middleware successfully got the new ETag back
	if resp.Header.Get("ETag") != `"new-etag"` {
		t.Errorf("Expected ETag %q, got %q", `"new-etag"`, resp.Header.Get("ETag"))
	}
}