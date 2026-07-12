package githubapi

import (
	"encoding/json"
)

type GithubEvent struct {
	Id        int
	Type     string
	Actor     Actor
	Repo      Repo
	Payload   json.RawMessage
	Public    bool
	CreatedAt string
}

type Actor struct {
	Id            int
	Login         string
	Display_login string
	Gravatar_id   string
	Url           string
	Avatar_url    string
}

type Repo struct {
	Id   int
	Name string
	Url  string
}

type CommitCommentEventPayload struct {
	Action string
	Comment CommitComment
}

type CommitComment struct {
	Html_url string
	Url string
	Id int
	Node_id string
	Body string
	Path string
	Position int
	Line int
	Commit_Id string
	Author_Association string
	User User
	Created_at string
	Updated_at string
}

type User struct {
	Login string
	Id int
	Node_id string
	Avatar_url string
	Gravatar_id string
	Url string
	Html_url string
	Followers_url string
	Following_url string
	Gists_url string
	Starred_url string
	Subscriptions_url string
	Organizations_url string
	Repos_url string
	Events_url string
	Received_events_url string
	Type string
	Site_admin bool
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
