package githubapi

import (
	"encoding/json"
)

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
