package githubapi

import (
	"encoding/json"
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
	// TODO
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
	HasUssues           bool        `json:"has_issues"`
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
	Id            int     `json:"id"`
	NodeId        string  `json:"node_id"`
	Url           string  `json:"url"`
	RepositoryUrl string  `json:"repository_url"`
	LabelsUrl     string  `json:"labels_url"`
	CommentsUrl   string  `json:"comments_url"`
	EventsUrl     string  `json:"events_url"`
	HtmlUrl       string  `json:"html_url"`
	Number        int     `json:"number"`
	State         string  `json:"state"`
	StateReason   string  `json:"state_reason"`
	Title         string  `json:"title"`
	Body          string  `json:"body"`
	User          User    `json:"user"`
	Labels        []Label `json:"labels"`
}

type Label struct {
}

type Comment struct {
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
