package main

import (
	"encoding/json"

	"fmt"

	"io"

	"os"

	"log"

	"net/http"
)

// TODO : Organize these structs into modules

type GithubEvent struct {
	id        int
	type_     string
	actor     Actor
	repo      Repo
	payload   json.RawMessage
	public    bool
	createdAt string
}

type Actor struct {
	id            int
	login         string
	display_login string
	gravatar_id   string
	url           string
	avatar_url    string
}

type Repo struct {
	id   int
	name string
	url  string
}

type CommitCommentEventPayload struct {
	
}

type CreateEventPayload struct {

}

type DeleteEventPayload struct {

}

type DiscussionEventPayload struct {

}

type ForkEventPayload struct {

}

type GollumEventPayload struct {

}

type IssueCommentEventPayload struct {

}

type IssuesEventPayload struct {

}

type MemberEventPayload struct {

}

type PublicEventPayload struct {

}

type PullRequestEventPayload struct {

}

type PullRequestReviewEventPayload struct {

}

type PullRequestReviewCommentEventPayload struct {

}

type PushEventPayload struct {

}

type ReleaseEventPayload struct {

}

type WatchEventPayload struct {

}

func main() {
	if len(os.Args) != 2 {
		log.Println("Usage: ./github-user-activity <username>")
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

	fmt.Println(body)
}
