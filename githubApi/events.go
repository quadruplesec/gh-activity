package githubapi

import (
	"encoding/json"
	"fmt"
)

type GithubEvent struct {
	Id        int             `json:"id"`
	Type      string          `json:"type"`
	Actor     User            `json:"actor"`
	Repo      Repo            `json:"repo"`
	Payload   json.RawMessage `json:"payload"`
	Public    bool            `json:"public"`
	CreatedAt string          `json:"created_at"`
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
	HtmlUrl           string `json:"html_url"`
	Url               string `json:"url"`
	Id                int    `json:"id"`
	NodeId            string `json:"node_id"`
	Body              string `json:"body"`
	Path              string `json:"path"`
	Position          int    `json:"position"`
	Line              int    `json:"line"`
	CommitId          string `json:"commit_id"`
	User              User   `json:"user"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	AuthorAssociation string `json:"author_association"`
}

type User struct {
	Login             string `json:"login"`
	Id                int    `json:"id"`
	NodeId            string `json:"node_id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_url"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type CreateEventPayload struct {
	Ref          string `json:"ref"`
	RefType      string `json:"ref_type"`
	FullRef      string `json:"full_ref"`
	MasterBranch string `json:"master_branch"`
	Description  string `json:"description"`
	PusherType   string `json:"pusher_type"`
}

type DeleteEventPayload struct {
	Ref        string `json:"ref"`
	RefType    string `json:"ref_type"`
	FullRef    string `json:"full_ref"`
	PusherType string `json:"pusher_type"`
}

type DiscussionEventPayload struct {
	Action     string     `json:"action"`
	Discussion Discussion `json:"discussion"` // TODO : Figure how this one works...
}

type Discussion struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
}

type ForkEventPayload struct {
	Action string     `json:"action"`
	Forkee Repository `json:"forkee"`
}

type Repository struct {
	Id                  int         `json:"id"`
	NodeId              string      `json:"node_id"`
	Name                string      `json:"name"`
	FullName            string      `json:"full_name"`
	Owner               User        `json:"owner"`
	Private             bool        `json:"private"`
	HtmlUrl             string      `json:"html_url"`
	Description         string      `json:"description"`
	Fork                bool        `json:"fork"`
	Url                 string      `json:"url"`
	ArchiveUrl          string      `json:"archive_url"`
	AssigneesUrl        string      `json:"assignees_url"`
	BlobsUrl            string      `json:"blobs_url"`
	BranchesUrl         string      `json:"branches_url"`
	CollaboratorsUrl    string      `json:"collaborators_url"`
	CommentsUrl         string      `json:"comments_url"`
	CommitsUrl          string      `json:"commits_url"`
	CompareUrl          string      `json:"compare_url"`
	ContentsUrl         string      `json:"contents_url"`
	ContributorsUrl     string      `json:"contributors_url"`
	DeploymentsUrl      string      `json:"deployments_url"`
	DownloadsUrl        string      `json:"downloads_url"`
	EventsUrl           string      `json:"events_url"`
	ForksUrl            string      `json:"forks_url"`
	GitCommitsUrl       string      `json:"git_commits_url"`
	GitRefsUrl          string      `json:"git_refs_url"`
	GitTagsUrl          string      `json:"git_tags_url"`
	GitUrl              string      `json:"git_url"`
	IssueCommentUrl     string      `json:"issue_comment_url"`
	IssueEventsUrl      string      `json:"issue_events_url"`
	IssuesUrl           string      `json:"issues_url"`
	KeysUrl             string      `json:"keys_url"`
	LabelsUrl           string      `json:"labels_url"`
	LanguagesUrl        string      `json:"languages_url"`
	MergesUrl           string      `json:"merges_url"`
	MilestonesUrl       string      `json:"milestones_url"`
	NotificationsUrl    string      `json:"notifications_url"`
	PullsUrl            string      `json:"pulls_url"`
	ReleasesUrl         string      `json:"releases_url"`
	SshUrl              string      `json:"ssh_url"`
	StargazersUrl       string      `json:"stargazers_url"`
	StatusesUrl         string      `json:"statuses_url"`
	TagsUrl             string      `json:"tags_url"`
	TeamsUrl            string      `json:"teams_url"`
	TreesUrl            string      `json:"trees_url"`
	CloneUrl            string      `json:"clone_url"`
	MirrorUrl           string      `json:"mirror_url"`
	HooksUrl            string      `json:"hooks_url"`
	SvnUrl              string      `json:"svn_url"`
	Homepage            string      `json:"homepage"`
	Language            string      `json:"language"`
	ForksCount          int         `json:"forks_count"`
	Forks               int         `json:"forks"`
	StargazersCount     int         `json:"stargazers_count"`
	WatchersCount       int         `json:"watchers_count"`
	Watchers            int         `json:"watchers"`
	Size                int         `json:"size"`
	DefaultBranch       string      `json:"default_branch"`
	OpenIssuesCount     int         `json:"open_issues_count"`
	OpenIssues          int         `json:"open_issues"`
	IsTemplate          bool        `json:"is_template"`
	License             License     `json:"license"`
	Topics              []string    `json:"topics"`
	HasIssues           bool        `json:"has_issues"`
	HasProjects         bool        `json:"has_projects"`
	HasWiki             bool        `json:"has_wiki"`
	HasPages            bool        `json:"has_pages"`
	HasDownloads        bool        `json:"has_downloads"`
	HasDiscussions      bool        `json:"has_discussions"`
	Archived            bool        `json:"archived"`
	Disabled            bool        `json:"disabled"`
	Visibility          string      `json:"visibility"`
	PushedAt            string      `json:"pushed_at"`
	CreatedAt           string      `json:"created_at"`
	UpdatedAt           string      `json:"updated_at"`
	Permissions         Permissions `json:"permissions"`
	AllowRebaseMerge    bool        `json:"allow_rebase_merge"`
	TempCloneToken      string      `json:"temp_clone_token"`
	AllowSquashMerge    bool        `json:"allow_squash_merge"`
	AllowAutoMerge      bool        `json:"allow_auto_merge"`
	DeleteBranchOnMerge bool        `json:"delete_branch_on_merge"`
	AllowMergeCommit    bool        `json:"allow_merge_commit"`
	SubscribersCount    int         `json:"subscribers_count"`
	NetworkCount        int         `json:"network_count"`
	Organization        *User       `json:"organization"`
	TemplateRepository  *Repository `json:"template_repository"`
	Source              *Repository `json:"source"`
}

type Permissions struct {
	Pull  bool `json:"pull"`
	Push  bool `json:"push"`
	Admin bool `json:"admin"`
}

type License struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	SpdxId string `json:"spdx_id"`
	Url    string `json:"url"`
	NodeId string `json:"node_id"`
}

type GollumEventPayload struct {
	Pages []Page `json:"pages"`
}

type Page struct {
	PageName string `json:"page_name"`
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	Action   string `json:"action"`
	Sha      string `json:"sha"`
	HtmlUrl  string `json:"html_url"`
}

type IssueCommentEventPayload struct {
	Action  string  `json:"action"`
	Issue   Issue   `json:"issue"`
	Comment Comment `json:"comment"`
}

type Issue struct {
	Id                int                 `json:"id"`
	NodeId            string              `json:"node_id"`
	Url               string              `json:"url"`
	RepositoryUrl     string              `json:"repository_url"`
	LabelsUrl         string              `json:"labels_url"`
	CommentsUrl       string              `json:"comments_url"`
	EventsUrl         string              `json:"events_url"`
	HtmlUrl           string              `json:"html_url"`
	Number            int                 `json:"number"`
	State             string              `json:"state"`
	StateReason       string              `json:"state_reason"`
	Title             string              `json:"title"`
	Body              string              `json:"body"`
	User              User                `json:"user"`
	Labels            []Label             `json:"labels"`
	Assignees         []User              `json:"assignees"`
	Milestone         *Milestone          `json:"milestone"`
	Locked            bool                `json:"locked"`
	ActiveLockReason  string              `json:"active_lock_reason"`
	Comments          int                 `json:"comments"`
	PullRequest       *PullRequestSummary `json:"pull_request"`
	ClosedAt          string              `json:"closed_at"`
	CreatedAt         string              `json:"created_at"`
	UpdatedAt         string              `json:"updated_at"`
	ClosedBy          *User               `json:"closed_by"`
	AuthorAssociation string              `json:"author_association"`
}

type Label struct {
	Id          int    `json:"id"`
	NodeId      string `json:"node_id"`
	Url         string `json:"url"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Default     bool   `json:"default"`
}

type Milestone struct {
	Url          string `json:"url"`
	HtmlUrl      string `json:"html_url"`
	LabelsUrl    string `json:"labels_url"`
	Id           int    `json:"id"`
	NodeId       string `json:"node_id"`
	Number       int    `json:"number"`
	State        string `json:"state"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Creator      User   `json:"creator"`
	OpenIssues   int    `json:"open_issues"`
	ClosedIssues int    `json:"closed_issues"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	ClosedAt     string `json:"closed_at"`
	DueOn        string `json:"due_on"`
}

type PullRequestSummary struct {
	Id       int    `json:"id"`
	NodeId   string `json:"node_id"`
	MergedAt string `json:"merged_at"`
	Url      string `json:"url"`
	HtmlUrl  string `json:"html_url"`
	DiffUrl  string `json:"diff_url"`
	PatchUrl string `json:"patch_url"`
}

type Comment struct {
	Body string `json:"body"`
}

type IssuesEventPayload struct {
	Action string `json:"action"`
	Issue  Issue  `json:"issue"`
}

type MemberEventPayload struct {
	Action string `json:"action"`
	Member User   `json:"member"`
}

type PublicEventPayload struct {
	// Empty Payload
}

type PullRequestEventPayload struct {
	Action      string      `json:"action"`
	Number      int         `json:"number"`
	PullRequest PullRequest `json:"pull_request"`
}

type PullRequest struct {
	Number int             `json:"number"`
	Title  string          `json:"title"`
	User   User            `json:"user"`
	Body   string          `json:"body"`
	Head   PullRequestHead `json:"head"`
}

type PullRequestHead struct {
	Repo Repo `json:"repo"`
}

type PullRequestReviewEventPayload struct {
	Action      string      `json:"action"`
	PullRequest PullRequest `json:"pull_request"`
}

type PullRequestReviewCommentEventPayload struct {
	Action      string      `json:"action"`
	PullRequest PullRequest `json:"pull_request"`
}

type PushEventPayload struct {
	RepositoryId int    `json:"repository_id"`
	PushId       int    `json:"push_id"`
	Ref          string `json:"ref"`
	Size         int    `json:"size"`
}

type ReleaseEventPayload struct {
	Action  string  `json:"action"`
	Release Release `json:"release"`
}

type Release struct {
	TagName string `json:"tag_name"`
	Name    string `json:"name"`
}

type WatchEventPayload struct {
	Action string `json:"action"`
}

type ActivityFormatter interface {
	FormatActivity(User, Repo) string
}

func (payload CommitCommentEventPayload) FormatActivity(user User, repo Repo) string {
	return fmt.Sprintf(" - Commented \"%s\" on commit %s in %s", payload.Comment.Body, payload.Comment.CommitId[:7], repo.Name)
}

func (payload CreateEventPayload) FormatActivity(user User, repo Repo) string {
	if payload.RefType == "repository" {
		return fmt.Sprintf(" - Created repository %s", repo.Name)
	}

	return fmt.Sprintf(" - Created %s '%s' in %s", payload.RefType, payload.Ref, repo.Name)
}

func (payload DeleteEventPayload) FormatActivity(user User, repo Repo) string {
	return fmt.Sprintf(" - Deleted %s '%s' in %s", payload.RefType, payload.Ref, repo.Name)
}

func (payload DiscussionEventPayload) FormatActivity(user User, repo Repo) string {
	return fmt.Sprintf(" - %s discussion #%d in %s: \"%s\"", payload.Action, payload.Discussion.Number, repo.Name, payload.Discussion.Title)
}

func (payload ForkEventPayload) FormatActivity(user User, repo Repo) string {
	return fmt.Sprintf(" - Forked %s to %s", repo.Name, payload.Forkee.FullName)
}

func (payload GollumEventPayload) FormatActivity(user User, repo Repo) string {
	if len(payload.Pages) == 0 {
		return fmt.Sprintf(" - Updated wiki in %s", repo.Name)
	}

	if len(payload.Pages) == 1 {
		return fmt.Sprintf(" - %s wiki page '%s' in %s", payload.Pages[0].Action, payload.Pages[0].Title, repo.Name)
	}

	return fmt.Sprintf(" - Updated %d wiki pages in %s", len(payload.Pages), repo.Name)
}

func (payload IssueCommentEventPayload) FormatActivity(user User, repo Repo) string {
	target := "issue"
	if payload.Issue.PullRequest != nil {
		target = "pull request"
	}

	return fmt.Sprintf(" - %s comment on %s #%d in %s", payload.Action, target, payload.Issue.Number, repo.Name)
}

func (payload IssuesEventPayload) FormatActivity(user User, repo Repo) string {
	return fmt.Sprintf(" - %s issue #%d in %s: \"%s\"", payload.Action, payload.Issue.Number, repo.Name, payload.Issue.Title)
}

func (payload MemberEventPayload) FormatActivity(user User, repo Repo) string {
	return fmt.Sprintf(" - %s %s as a collaborator to %s", payload.Action, payload.Member.Login, repo.Name)
}

func (payload PublicEventPayload) FormatActivity(user User, repo Repo) string {
	return fmt.Sprintf(" - Made %s public", repo.Name)
}

func (payload PullRequestEventPayload) FormatActivity(user User, repo Repo) string {
	return fmt.Sprintf(" - %s pull request #%d in %s: \"%s\"", payload.Action, payload.Number, repo.Name, payload.PullRequest.Title)
}

func (payload PullRequestReviewEventPayload) FormatActivity(user User, repo Repo) string {
	return fmt.Sprintf(" - %s a review on pull request #%d in %s", payload.Action, payload.PullRequest.Number, repo.Name)
}

func (payload PullRequestReviewCommentEventPayload) FormatActivity(user User, repo Repo) string {
	return fmt.Sprintf(" - %s a review comment on pull request #%d in %s", payload.Action, payload.PullRequest.Number, repo.Name)
}

func (payload PushEventPayload) FormatActivity(user User, repo Repo) string {
	return fmt.Sprintf(" - Pushed %d commits to %s", payload.Size, repo.Name)
}

func (payload ReleaseEventPayload) FormatActivity(user User, repo Repo) string {
	return fmt.Sprintf(" - %s release %s in %s", payload.Action, payload.Release.TagName, repo.Name)
}

func (payload WatchEventPayload) FormatActivity(user User, repo Repo) string {
	return fmt.Sprintf(" - Starred %s", repo.Name)
}

type UnmarshalEventPayloadError struct {
	eventType string
	message   string
}

func (e UnmarshalEventPayloadError) Error() string {
	return fmt.Sprintf("%s: %s", e.message, e.eventType)
}

func (event *GithubEvent) UnmarshalEventPayload() (ActivityFormatter, error) {
	switch event.Type {
	case "CommitCommentEvent":
		var payload CommitCommentEventPayload
		err := json.Unmarshal(event.Payload, &payload)
		if err != nil {
			return nil, UnmarshalEventPayloadError{eventType: event.Type, message: "An error occured when unmarshalling a payload of type"}
		}
		return payload, nil
	case "CreateEvent":
		var payload CreateEventPayload
		err := json.Unmarshal(event.Payload, &payload)
		if err != nil {
			return nil, UnmarshalEventPayloadError{eventType: event.Type, message: "An error occured when unmarshalling a payload of type"}
		}
		return payload, nil
	case "DeleteEvent":
		var payload DeleteEventPayload
		err := json.Unmarshal(event.Payload, &payload)
		if err != nil {
			return nil, UnmarshalEventPayloadError{eventType: event.Type, message: "An error occured when unmarshalling a payload of type"}
		}
		return payload, nil
	case "DiscussionEvent":
		var payload DiscussionEventPayload
		err := json.Unmarshal(event.Payload, &payload)
		if err != nil {
			return nil, UnmarshalEventPayloadError{eventType: event.Type, message: "An error occured when unmarshalling a payload of type"}
		}
		return payload, nil
	case "ForkEvent":
		var payload ForkEventPayload
		err := json.Unmarshal(event.Payload, &payload)
		if err != nil {
			return nil, UnmarshalEventPayloadError{eventType: event.Type, message: "An error occured when unmarshalling a payload of type"}
		}
		return payload, nil
	case "GollumEvent":
		var payload GollumEventPayload
		err := json.Unmarshal(event.Payload, &payload)
		if err != nil {
			return nil, UnmarshalEventPayloadError{eventType: event.Type, message: "An error occured when unmarshalling a payload of type"}
		}
		return payload, nil
	case "IssueCommentEvent":
		var payload IssueCommentEventPayload
		err := json.Unmarshal(event.Payload, &payload)
		if err != nil {
			return nil, UnmarshalEventPayloadError{eventType: event.Type, message: "An error occured when unmarshalling a payload of type"}
		}
		return payload, nil
	case "IssuesEvent":
		var payload IssuesEventPayload
		err := json.Unmarshal(event.Payload, &payload)
		if err != nil {
			return nil, UnmarshalEventPayloadError{eventType: event.Type, message: "An error occured when unmarshalling a payload of type"}
		}
		return payload, nil
	case "MemberEvent":
		var payload MemberEventPayload
		err := json.Unmarshal(event.Payload, &payload)
		if err != nil {
			return nil, UnmarshalEventPayloadError{eventType: event.Type, message: "An error occured when unmarshalling a payload of type"}
		}
		return payload, nil
	case "PublicEvent":
		var payload PublicEventPayload
		err := json.Unmarshal(event.Payload, &payload)
		if err != nil {
			return nil, UnmarshalEventPayloadError{eventType: event.Type, message: "An error occured when unmarshalling a payload of type"}
		}
		return payload, nil
	case "PullRequestEvent":
		var payload PullRequestEventPayload
		err := json.Unmarshal(event.Payload, &payload)
		if err != nil {
			return nil, UnmarshalEventPayloadError{eventType: event.Type, message: "An error occured when unmarshalling a payload of type"}
		}
		return payload, nil
	case "PullRequestReviewEvent":
		var payload PullRequestReviewEventPayload
		err := json.Unmarshal(event.Payload, &payload)
		if err != nil {
			return nil, UnmarshalEventPayloadError{eventType: event.Type, message: "An error occured when unmarshalling a payload of type"}
		}
		return payload, nil
	case "PullRequestReviewCommentEvent":
		var payload PullRequestReviewCommentEventPayload
		err := json.Unmarshal(event.Payload, &payload)
		if err != nil {
			return nil, UnmarshalEventPayloadError{eventType: event.Type, message: "An error occured when unmarshalling a payload of type"}
		}
		return payload, nil
	case "PushEvent":
		var payload PushEventPayload
		err := json.Unmarshal(event.Payload, &payload)
		if err != nil {
			return nil, UnmarshalEventPayloadError{eventType: event.Type, message: "An error occured when unmarshalling a payload of type"}
		}
		return payload, nil
	case "ReleaseEvent":
		var payload ReleaseEventPayload
		err := json.Unmarshal(event.Payload, &payload)
		if err != nil {
			return nil, UnmarshalEventPayloadError{eventType: event.Type, message: "An error occured when unmarshalling a payload of type"}
		}
		return payload, nil
	case "WatchEvent":
		var payload WatchEventPayload
		err := json.Unmarshal(event.Payload, &payload)
		if err != nil {
			return nil, UnmarshalEventPayloadError{eventType: event.Type, message: "An error occured when unmarshalling a payload of type"}
		}
		return payload, nil
	default:
		return nil, UnmarshalEventPayloadError{eventType: event.Type, message: "Unknown Github Event Type"}
	}
}
