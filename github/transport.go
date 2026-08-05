package github

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type CachingMiddleware struct {
	Base     http.RoundTripper
	CacheDir string
}

func NewCachingMiddleware(base http.RoundTripper, cacheDir string) *CachingMiddleware {
	if base == nil {
		base = http.DefaultTransport
	}
	return &CachingMiddleware{
		Base:     base,
		CacheDir: cacheDir,
	}
}

func parseMaxAge(cacheControl string) time.Duration {
	for _, part := range strings.Split(cacheControl, ",") {
		part = strings.TrimSpace(part)
		if strings.Contains(part, "max-age=") {
			ttl := strings.TrimPrefix(part, "max-age=")
			if seconds, err := strconv.Atoi(ttl); err == nil {
				return time.Duration(seconds) * time.Second
			}
		}
	}
	return 0
}

func createCachedResponse(req *http.Request, envelope ActivityEnvelope) (*http.Response, error) {
	bodyBytes, err := json.Marshal(envelope.Activity)
	if err != nil {
		return nil, err
	}

	fakeResponse := http.Response{
		Status:        "200 OK",
		StatusCode:    http.StatusOK,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          io.NopCloser(bytes.NewReader(bodyBytes)),
		ContentLength: int64(len(bodyBytes)),
		Request:       req,
		Header:        make(http.Header),
	}

	fakeResponse.Header.Set("Content-Type", "application/json; charset=utf-8")
	fakeResponse.Header.Set("ETag", envelope.ETag)

	return &fakeResponse, nil
}

func (c *CachingMiddleware) RoundTrip(req *http.Request) (*http.Response, error) {
	cacheKey := URLToCacheKey(req.URL.Path)

	envelope, err := LoadCache(c.CacheDir, cacheKey)

	// Cache Miss
	if err != nil {
		response, err := c.Base.RoundTrip(req)
		fetchTime := time.Now()
		if err != nil {
			return nil, err
		}

		if response.StatusCode != http.StatusOK {
			return response, nil
		}

		// If we read the HTTP Response as a normal stream, we would return an empty, closed body to the calling function.
		// Instead, we have to read the HTTP Response into memory, close the original stream,
		// reconstruct the response body and only then return to caller
		bodyBytes, err := io.ReadAll(response.Body)
		response.Body.Close()
		if err != nil {
			return nil, err
		}

		var jsonResponse ActivityFeed
		if err := json.Unmarshal(bodyBytes, &jsonResponse); err != nil {
			return nil, err
		}

		envelope := ActivityEnvelope{
			Activity:  jsonResponse,
			ETag:      response.Header.Get("ETag"),
			FetchedAt: fetchTime,
			TTL:       parseMaxAge(response.Header.Get("Cache-Control")),
		}

		SaveCache(c.CacheDir, envelope, cacheKey)

		// We use NopCloser so that, when the caller function attempts to close to body
		// nothing will happen.
		response.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		return response, nil
	}

	// Cache is still fresh
	if !IsStale(envelope) {
		// Create fake HTTP response and send it to the caller
		return createCachedResponse(req, envelope)
	}

	// Stale Cache
	// In this step we have to clone the request and set the "If-None-Match" field of the header
	// to use the cached ETag. If the server responds with "304 Not Modified" we refresh the cache
	// and pass the cached response. If the server responds with "200 OK", we update the cache, and pass
	// the new response.
	clone := req.Clone(req.Context())
	clone.Header.Set("If-None-Match", envelope.ETag)

	response, err := c.Base.RoundTrip(clone)
	fetchTime := time.Now()
	if err != nil {
		return nil, err
	}
	// Cache is still valid
	if response.StatusCode == http.StatusNotModified {
		envelope.FetchedAt = time.Now()
		SaveCache(c.CacheDir, envelope, cacheKey)
		return createCachedResponse(req, envelope)
	}
	// Cache Invalidation
	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		response.Body.Close()
		if err != nil {
			return nil, err
		}

		var jsonResponse ActivityFeed
		if err := json.Unmarshal(bodyBytes, &jsonResponse); err != nil {
			return nil, err
		}

		envelope := ActivityEnvelope{
			Activity:  jsonResponse,
			ETag:      response.Header.Get("ETag"),
			FetchedAt: fetchTime,
			TTL:       parseMaxAge(response.Header.Get("Cache-Control")),
		}

		SaveCache(c.CacheDir, envelope, cacheKey)

		response.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		return response, nil
	}

	return response, nil
}