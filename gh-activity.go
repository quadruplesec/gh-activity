package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"
	"sync"

	"github.com/quadruplesec/gh-activity/github"
)

// If count is greater than 1, outputs with plural "commits"; if count is 1, outputs with singular "commit"; if 0, do nothing
func printPushes(count int, repo string) {
	if count == 0 {
		return
	}
	word := "commits"
	if count == 1 {
		word = "commit"
	}
	fmt.Printf(" - Pushed %d %s to %s\n", count, word, repo)
}

func main() {
	validEvents := []string{
		"CommitCommentEvent",
		"CreateEvent",
		"DeleteEvent",
		"DiscussionEvent",
		"ForkEvent",
		"GollumEvent",
		"IssueCommentEvent",
		"IssuesEvent",
		"MemberEvent",
		"PublicEvent",
		"PullRequestEvent",
		"PullRequestReviewEvent",
		"PullRequestReviewCommentEvent",
		"PushEvent",
		"ReleaseEvent",
		"WatchEvent",
	}

	filtersUsage := fmt.Sprintf("Space-separated list of event types to filter by.\nValid values are:\n - %s",
        strings.Join(validEvents, "\n - "))
    filters := flag.String("filters", "", filtersUsage)

    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage: gh-activity [flags] <username1> [username2] ...\n\nFlags:\n")
        flag.PrintDefaults()
    }
    flag.Parse()

	if len(flag.Args()) < 1 {
        log.Fatalf("Error: At least one valid username must be provided after all flags.\nUsage: gh-activity [flags] <username1> [username2] ...\nRun 'gh-activity -h' for more help.")
    }

	var filtersSlice []string
	if *filters != "" {
		filtersSlice = strings.Split(*filters, " ")
	}
	// Check if all filter flags are valid
	for _, filter := range filtersSlice {
		if !slices.Contains(validEvents, filter) {
			log.Fatalf("Invalid filter flag '%s'. Run 'gh-activity -h' for more help.", filter)
		}
	}

	usernames := flag.Args()

	// HTTP client with custom middleware to handle caching
	appCacheDir, err := github.GetAppCacheDir()
	if err != nil {
		log.Fatalf("Could not locate the system cache directory: %v", err)
	}

	client := &http.Client{
		Transport: github.NewCachingMiddleware(http.DefaultTransport, appCacheDir),
	}

	var wg sync.WaitGroup

	for _, username := range usernames {
		wg.Add(1)

		go func(uname string) {
			defer wg.Done()

			request := fmt.Sprintf("https://api.github.com/users/%s/events", uname)

			resp, err := client.Get(request)
			if err != nil {
				log.Printf("Network error for %s: %v\n", uname, err)
				return
			}

			if resp.StatusCode != http.StatusOK {
				log.Printf("GitHub API returned status %s for %s\n", resp.Status, uname)
				resp.Body.Close()
				return
			}

			var jsonResponse []github.Event
			if err := json.NewDecoder(resp.Body).Decode(&jsonResponse); err != nil {
				log.Printf("Error parsing JSON for %s: %v\n", uname, err)
				resp.Body.Close()
				return
			}
			resp.Body.Close()

			var filteredEvents github.ActivityFeed
			for _, event := range jsonResponse {
				// If specific events were set with the filters flag, remove them
				if len(filtersSlice) > 0 && !slices.Contains(filtersSlice, event.Type) {
					continue
				}
				filteredEvents = append(filteredEvents, event)
			}

			fmt.Printf("Recent GitHub Activity of %s:\n", username)
			if len(filteredEvents) > 0 {
				for _, line := range filteredEvents.Format() {
					fmt.Println(line)
				}
				return
			}
			fmt.Println("No activity was found.")
		}(username)
	}
}
