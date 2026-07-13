package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"log"
	"net/http"

	"github.com/quadruplesec/github-user-activity/githubApi"
)

func main() {
	if len(os.Args) != 2 {
		log.Println("Usage: ./github-user-activity <username>")	}

	username := os.Args[1]
	request := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	resp, err := http.Get(request)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var jsonResponse []githubapi.GithubEvent
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		log.Printf("Error: %s", err)
	}

	fmt.Printf("Recent Github Activity of %s:\n", username)
	for i := 0; i < len(jsonResponse); i++ {
		formatter, err := jsonResponse[i].UnmarshalEventPayload()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(formatter.FormatActivity(
			jsonResponse[i].Actor,
			jsonResponse[i].Repo,
		))
	}
}
