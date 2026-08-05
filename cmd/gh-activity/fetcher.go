package main

import (
	"encoding/json"
	"fmt"
	"log"
	"io"
	"net/http"
	"slices"
	"sync"

	"github.com/quadruplesec/gh-activity/github"
)

var apiBaseURL = "https://api.github.com"

func FetchUsers(usernames []string, filtersSlice []string, client *http.Client, out io.Writer) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, username := range usernames {
		wg.Add(1)

		go func(uname string) {
			defer wg.Done()

			request := fmt.Sprintf("%s/users/%s/events", apiBaseURL, uname)

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

			mu.Lock()

			fmt.Fprintf(out, "Recent GitHub Activity of %s:\n", uname)
			if len(filteredEvents) > 0 {
				for _, line := range filteredEvents.Format() {
					fmt.Fprintln(out, line)
				}
			} else {
				fmt.Fprintf(out, "No activity was found user %s.\n", uname)
			}
			
			fmt.Fprintln(out) // Blank Line between users

			mu.Unlock()
		}(username)
	}

	wg.Wait()
}