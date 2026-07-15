package github

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ActivityEnvelope struct {
	Activity  ActivityFeed  `json:"activity"`
	ETag      string        `json:"etag"`
	FetchedAt time.Time     `json:"fetched_at"`
	TTL       time.Duration `json:"ttl"`
}

func URLToCacheKey(urlStr string) string {
    clean := strings.TrimPrefix(urlStr, "https://api.github.com/")
    clean = strings.TrimPrefix(clean, "/")
    clean = strings.ReplaceAll(clean, "/", "_")
    return clean + ".json"
}

func SaveCache(activity ActivityEnvelope, cacheKey string) error {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return err
	}
	appCacheDir := filepath.Join(cacheDir, "gh-activity")

	err = os.MkdirAll(appCacheDir, 0o700) // 0o700 -> rwx------
	if err != nil {
		return err
	}

	filePath := filepath.Join(appCacheDir, cacheKey+".json")

	data, err := json.Marshal(activity)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0o600) // 0o600 -> rw-------
}

func IsStale(envelope ActivityEnvelope) bool {
	return time.Since(envelope.FetchedAt) > envelope.TTL
}

func LoadCache(cacheKey string) (ActivityEnvelope, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return ActivityEnvelope{}, err
	}
	appCacheDir := filepath.Join(cacheDir, "gh-activity")
	filePath := filepath.Join(appCacheDir, cacheKey+".json")

	data, err := os.ReadFile(filePath)
	if err != nil {
		return ActivityEnvelope{}, err
	}

	var envelope ActivityEnvelope
	err = json.Unmarshal(data, &envelope)
	if err != nil {
		return ActivityEnvelope{}, err
	}

	return envelope, nil
}
