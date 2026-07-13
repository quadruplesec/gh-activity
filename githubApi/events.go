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
	Ref           string `json:"ref"`
	Ref_type      string `json:"ref_type"`
	Full_ref      string `json:"full_ref"`
	Master_branch string `json:"master_branch"`
	Description   string `json:"description"`
	Pusher_type   string `json:"pusher_type"`
}

type DeleteEventPayload struct {
	Ref         string `json:"ref"`
	Ref_type    string `json:"ref_type"`
	Full_ref    string `json:"full_ref"`
	Pusher_type string `json:"pusher_type"`
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
	Id                     int         `json:"id"`
	Node_id                string      `json:"node_id"`
	Name                   string      `json:"name"`
	Full_name              string      `json:"full_name"`
	Owner                  User        `json:"owner"`
	Private                bool        `json:"private"`
	Html_url               string      `json:"html_url"`
	Description            string      `json:"description"`
	Fork                   bool        `json:"fork"`
	Url                    string      `json:"url"`
	Archive_url            string      `json:"archive_url"`
	Assignees_url          string      `json:"assignees_url"`
	Blobs_url              string      `json:"blobs_url"`
	Branches_url           string      `json:"branches_url"`
	Collaborators_url      string      `json:"collaborators_url"`
	Comments_url           string      `json:"comments_url"`
	Commits_url            string      `json:"commits_url"`
	Compare_url            string      `json:"compare_url"`
	Contents_url           string      `json:"contents_url"`
	Contributors_url       string      `json:"contributors_url"`
	Deployments_url        string      `json:"deployments_url"`
	Downloads_url          string      `json:"downloads_url"`
	Events_url             string      `json:"events_url"`
	Forks_url              string      `json:"forks_url"`
	Git_commits_url        string      `json:"git_commits_url"`
	Git_refs_url           string      `json:"git_refs_url"`
	Git_tags_url           string      `json:"git_tags_url"`
	Git_url                string      `json:"git_url"`
	Issue_comment_url      string      `json:"issue_comment_url"`
	Issue_events_url       string      `json:"issue_events_url"`
	Issues_url             string      `json:"issues_url"`
	Keys_url               string      `json:"keys_url"`
	Labels_url             string      `json:"labels_url"`
	Languages_url          string      `json:"languages_url"`
	Merges_url             string      `json:"merges_url"`
	Milestones_url         string      `json:"milestones_url"`
	Notifications_url      string      `json:"notifications_url"`
	Pulls_url              string      `json:"pulls_url"`
	Releases_url           string      `json:"releases_url"`
	Ssh_url                string      `json:"ssh_url"`
	Stargazers_url         string      `json:"stargazers_url"`
	Statuses_url           string      `json:"statuses_url"`
	Tags_url               string      `json:"tags_url"`
	Teams_url              string      `json:"teams_url"`
	Trees_url              string      `json:"trees_url"`
	Clone_url              string      `json:"clone_url"`
	Mirror_url             string      `json:"mirror_url"`
	Hooks_url              string      `json:"hooks_url"`
	Svn_url                string      `json:"svn_url"`
	Homepage               string      `json:"homepage"`
	Language               string      `json:"language"`
	Forks_count            int         `json:"forks_count"`
	Forks                  int         `json:"forks"`
	Stargazers_count       int         `json:"stargazers_count"`
	Watchers_count         int         `json:"watchers_count"`
	Watchers               int         `json:"watchers"`
	Size                   int         `json:"size"`
	Default_branch         string      `json:"default_branch"`
	Open_issues_count      int         `json:"open_issues_count"`
	Open_issues            int         `json:"open_issues"`
	Is_template            bool        `json:"is_template"`
	License                License     `json:"license"`
	Topics                 []string    `json:"topics"`
	Has_issues             bool        `json:"has_issues"`
	Has_projects           bool        `json:"has_projects"`
	Has_wiki               bool        `json:"has_wiki"`
	Has_pages              bool        `json:"has_pages"`
	Has_downloads          bool        `json:"has_downloads"`
	Has_discussions        bool        `json:"has_discussions"`
	Archived               bool        `json:"archived"`
	Disabled               bool        `json:"disabled"`
	Visibility             string      `json:"visibility"`
	Pushed_at              string      `json:"pushed_at"`
	Created_at             string      `json:"created_at"`
	Updated_at             string      `json:"updated_at"`
	Permissions            Permissions `json:"permissions"`
	Allow_rebase_merge     bool        `json:"allow_rebase_merge"`
	Temp_clone_token       string      `json:"temp_clone_token"`
	Allow_squash_merge     bool        `json:"allow_squash_merge"`
	Allow_auto_merge       bool        `json:"allow_auto_merge"`
	Delete_branch_on_merge bool        `json:"delete_branch_on_merge"`
	Allow_merge_commit     bool        `json:"allow_merge_commit"`
	Subscribers_count      int         `json:"subscribers_count"`
	Network_count          int         `json:"network_count"`
	Organization           User        `json:"organization"`
	Template_repository    *Repository `json:"template_repository"`
	Source                 *Repository `json:"source"`
}

type Permissions struct {
	Pull  bool `json:"pull"`
	Push  bool `json:"push"`
	Admin bool `json:"admin"`
}

type License struct {
	Key     string `json:"key"`
	Name    string `json:"name"`
	Spdx_id string `json:"spdx_id"`
	Url     string `json:"url"`
	Node_id string `json:"node_id"`
}

type GollumEventPayload struct {
	Pages []Page `json:"pages"`
}

type Page struct {
	Page_name string `json:"page_name"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	Action    string `json:"action"`
	Sha       string `json:"sha"`
	Html_url  string `json:"html_url"`
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
