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
		fmt.Fprintf(os.Stderr, "Usage: gh-activity [flags] <username>\n\nFlags:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if len(flag.Args()) != 1 {
		log.Fatalf("Error: Exactly one valid username must be provided after all flags.\nUsage: gh-activity [flags] <username>\nRun 'gh-activity -h' for more help.")
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

	username := flag.Args()[0]
	request := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	resp, err := http.Get(request)
	if err != nil {
		log.Fatalf("Network error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("GitHub API returned status: %s", resp.Status)
	}

	var jsonResponse []github.Event
	if err := json.NewDecoder(resp.Body).Decode(&jsonResponse); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

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
}
