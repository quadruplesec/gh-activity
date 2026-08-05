package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"

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

	FetchUsers(usernames, filtersSlice, client, os.Stdout)
}
