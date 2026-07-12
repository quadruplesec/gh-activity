package githubapi

import (
	"encoding/json"
)

type GithubEvent struct {
	Id        int             `json:"id"`
	Type      string          `json:"type"`
	Actor     Actor           `json:"actor"`
	Repo      Repo            `json:"repo"`
	Payload   json.RawMessage `json:"payload"`
	Public    bool            `json:"public"`
	CreatedAt string          `json:"created_at"`
}

type Actor struct {
	Id            int    `json:"id"`
	Login         string `json:"login"`
	Display_login string `json:"display_login"`
	Gravatar_id   string `json:"gravatar_id"`
	Url           string `json:"url"`
	Avatar_url    string `json:"avatar_url"`
}

type Repo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type CommitCommentEventPayload struct {
	Action  string        `json:"action"`
	Comment CommitComment `json:"comment"`
}

type CommitComment struct {
	Html_url           string `json:"html_url"`
	Url                string `json:"url"`
	Id                 int    `json:"id"`
	Node_id            string `json:"node_id"`
	Body               string `json:"body"`
	Path               string `json:"path"`
	Position           int    `json:"position"`
	Line               int    `json:"line"`
	Commit_Id          string `json:"commit_id"`
	User               User   `json:"user"`
	Created_at         string `json:"created_at"`
	Updated_at         string `json:"updated_at"`
	Author_Association string `json:"autor_association"`
}

type User struct {
	Login               string `json:"login"`
	Id                  int    `json:"id"`
	Node_id             string `json:"node_id"`
	Avatar_url          string `json:"avatar_url"`
	Gravatar_id         string `json:"gravatar_url"`
	Url                 string `json:"url"`
	Html_url            string `json:"html_url"`
	Followers_url       string `json:"followers_url"`
	Following_url       string `json:"following_url"`
	Gists_url           string `json:"gists_url"`
	Starred_url         string `json:"starred_url"`
	Subscriptions_url   string `json:"subscriptions_url"`
	Organizations_url   string `json:"organizations_url"`
	Repos_url           string `json:"repos_url"`
	Events_url          string `json:"events_url"`
	Received_events_url string `json:"received_events_url"`
	Type                string `json:"type"`
	Site_admin          bool   `json:"site_admin"`
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
