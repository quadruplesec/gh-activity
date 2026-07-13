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
		log.Println("Usage: ./github-user-activity <username>")
		return
	}

	var username string = os.Args[1]
	var request string = fmt.Sprintf("https://api.github.com/users/%s/events", username)

	// TODO : Check if err automatically handles non-existent users. If not, handle that

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
	fmt.Printf("%v", jsonResponse)

	for i := 0; i < len(jsonResponse); i++ {

	}
}
