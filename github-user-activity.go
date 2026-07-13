package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/quadruplesec/github-user-activity/githubApi"
)

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

	fmt.Printf("Recent Github Activity of %s:\n", username)
	for i := 0; i < len(jsonResponse); i++ {
		formatter, err := jsonResponse[i].UnmarshalEventPayload()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(formatter.FormatActivity(
			jsonResponse[i].Actor,
			jsonResponse[i].Repo,
		))
	}
}
