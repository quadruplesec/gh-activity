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
