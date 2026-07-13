package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/quadruplesec/github-user-activity/githubApi"
)

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
	if len(os.Args) != 2 {
		log.Fatal("Usage: ./github-user-activity <username>")
	}

	username := os.Args[1]
	request := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	resp, err := http.Get(request)
	if err != nil {
		log.Fatalf("Network error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("GitHub API returned status: %s", resp.Status)
	}

	var jsonResponse []githubapi.GithubEvent
	if err := json.NewDecoder(resp.Body).Decode(&jsonResponse); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	pushCount := 0
	pushRepo := ""

	fmt.Printf("Recent Github Activity of %s:\n", username)
	for _, event := range jsonResponse {
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
