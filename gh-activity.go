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

	pushCount := 0
	pushRepo := ""

	fmt.Printf("%s", github.Colorize(github.Bold, fmt.Sprintf("Recent GitHub Activity of %s:\n", username)))
	for _, event := range jsonResponse {
		if len(filtersSlice) > 0 && !slices.Contains(filtersSlice, event.Type) {
			if pushCount > 0 {
				printPushes(pushCount, pushRepo)
				pushRepo = ""
				pushCount = 0
			}
			continue
		}

		if event.Type == "PushEvent" {
			if pushRepo == "" || pushRepo == event.Repo.Name {
				pushRepo = event.Repo.Name
				pushCount++
				continue
			}

			printPushes(pushCount, pushRepo)
			pushRepo = event.Repo.Name
			pushCount = 1
			continue
		}

		if pushCount > 0 {
			printPushes(pushCount, pushRepo)
			pushRepo = ""
			pushCount = 0
		}

		formatter, err := event.UnmarshalEventPayload()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(formatter.FormatActivity(event.Actor, event.Repo))
	}
}
