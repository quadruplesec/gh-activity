package github

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type ActivityEnvelope struct {
	Activity  ActivityFeed  `json:"activity"`
	ETag      string        `json:"etag"`
	FetchedAt time.Time     `json:"fetched_at"`
	TTL       time.Duration `json:"ttl"`
}

func SaveCache(activity ActivityEnvelope, username string) error {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return err
	}
	appCacheDir := filepath.Join(cacheDir, "gh-activity")

	err = os.MkdirAll(appCacheDir, 0o700) // 0o700 -> rwx------
	if err != nil {
		return err
	}

	filePath := filepath.Join(appCacheDir, username+".json")

	data, err := json.Marshal(activity)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0o600) // 0o600 -> rw-------
}
