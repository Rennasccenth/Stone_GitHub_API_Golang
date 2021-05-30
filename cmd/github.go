package cmd

import (
	"Stone_GitHub_API_Golang/env"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type User struct {
	Login                   string      `json:"login"`
	Id                      int         `json:"id"`
	NodeId                  string      `json:"node_id"`
	AvatarUrl               string      `json:"avatar_url"`
	GravatarId              string      `json:"gravatar_id"`
	Url                     string      `json:"url"`
	HtmlUrl                 string      `json:"html_url"`
	FollowersUrl            string      `json:"followers_url"`
	FollowingUrl            string      `json:"following_url"`
	GistsUrl                string      `json:"gists_url"`
	StarredUrl              string      `json:"starred_url"`
	SubscriptionsUrl        string      `json:"subscriptions_url"`
	OrganizationsUrl        string      `json:"organizations_url"`
	ReposUrl                string      `json:"repos_url"`
	EventsUrl               string      `json:"events_url"`
	ReceivedEventsUrl       string      `json:"received_events_url"`
	Type                    string      `json:"type"`
	SiteAdmin               bool        `json:"site_admin"`
	Name                    string      `json:"name"`
	Company                 interface{} `json:"company"`
	Blog                    string      `json:"blog"`
	Location                interface{} `json:"location"`
	Email                   interface{} `json:"email"`
	Hireable                interface{} `json:"hireable"`
	Bio                     string      `json:"bio"`
	TwitterUsername         interface{} `json:"twitter_username"`
	PublicRepos             int         `json:"public_repos"`
	PublicGists             int         `json:"public_gists"`
	Followers               int         `json:"followers"`
	Following               int         `json:"following"`
	CreatedAt               time.Time   `json:"created_at"`
	UpdatedAt               time.Time   `json:"updated_at"`
	PrivateGists            int         `json:"private_gists"`
	TotalPrivateRepos       int         `json:"total_private_repos"`
	OwnedPrivateRepos       int         `json:"owned_private_repos"`
	DiskUsage               int         `json:"disk_usage"`
	Collaborators           int         `json:"collaborators"`
	TwoFactorAuthentication bool        `json:"two_factor_authentication"`
	Plan                    struct {
		Name          string `json:"name"`
		Space         int    `json:"space"`
		Collaborators int    `json:"collaborators"`
		PrivateRepos  int    `json:"private_repos"`
	} `json:"plan"`
}

type Repository struct {
	Id       int    `json:"id"`
	NodeId   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
	Owner    struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
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
	} `json:"owner"`
	HtmlUrl          string      `json:"html_url"`
	Description      *string     `json:"description"`
	Fork             bool        `json:"fork"`
	Url              string      `json:"url"`
	ForksUrl         string      `json:"forks_url"`
	KeysUrl          string      `json:"keys_url"`
	CollaboratorsUrl string      `json:"collaborators_url"`
	TeamsUrl         string      `json:"teams_url"`
	HooksUrl         string      `json:"hooks_url"`
	IssueEventsUrl   string      `json:"issue_events_url"`
	EventsUrl        string      `json:"events_url"`
	AssigneesUrl     string      `json:"assignees_url"`
	BranchesUrl      string      `json:"branches_url"`
	TagsUrl          string      `json:"tags_url"`
	BlobsUrl         string      `json:"blobs_url"`
	GitTagsUrl       string      `json:"git_tags_url"`
	GitRefsUrl       string      `json:"git_refs_url"`
	TreesUrl         string      `json:"trees_url"`
	StatusesUrl      string      `json:"statuses_url"`
	LanguagesUrl     string      `json:"languages_url"`
	StargazersUrl    string      `json:"stargazers_url"`
	ContributorsUrl  string      `json:"contributors_url"`
	SubscribersUrl   string      `json:"subscribers_url"`
	SubscriptionUrl  string      `json:"subscription_url"`
	CommitsUrl       string      `json:"commits_url"`
	GitCommitsUrl    string      `json:"git_commits_url"`
	CommentsUrl      string      `json:"comments_url"`
	IssueCommentUrl  string      `json:"issue_comment_url"`
	ContentsUrl      string      `json:"contents_url"`
	CompareUrl       string      `json:"compare_url"`
	MergesUrl        string      `json:"merges_url"`
	ArchiveUrl       string      `json:"archive_url"`
	DownloadsUrl     string      `json:"downloads_url"`
	IssuesUrl        string      `json:"issues_url"`
	PullsUrl         string      `json:"pulls_url"`
	MilestonesUrl    string      `json:"milestones_url"`
	NotificationsUrl string      `json:"notifications_url"`
	LabelsUrl        string      `json:"labels_url"`
	ReleasesUrl      string      `json:"releases_url"`
	DeploymentsUrl   string      `json:"deployments_url"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	PushedAt         time.Time   `json:"pushed_at"`
	GitUrl           string      `json:"git_url"`
	SshUrl           string      `json:"ssh_url"`
	CloneUrl         string      `json:"clone_url"`
	SvnUrl           string      `json:"svn_url"`
	Homepage         interface{} `json:"homepage"`
	Size             int         `json:"size"`
	StargazersCount  int         `json:"stargazers_count"`
	WatchersCount    int         `json:"watchers_count"`
	Language         string      `json:"language"`
	HasIssues        bool        `json:"has_issues"`
	HasProjects      bool        `json:"has_projects"`
	HasDownloads     bool        `json:"has_downloads"`
	HasWiki          bool        `json:"has_wiki"`
	HasPages         bool        `json:"has_pages"`
	ForksCount       int         `json:"forks_count"`
	MirrorUrl        interface{} `json:"mirror_url"`
	Archived         bool        `json:"archived"`
	Disabled         bool        `json:"disabled"`
	OpenIssuesCount  int         `json:"open_issues_count"`
	License          interface{} `json:"license"`
	Forks            int         `json:"forks"`
	OpenIssues       int         `json:"open_issues"`
	Watchers         int         `json:"watchers"`
	DefaultBranch    string      `json:"default_branch"`
	Permissions      struct {
		Admin bool `json:"admin"`
		Push  bool `json:"push"`
		Pull  bool `json:"pull"`
	} `json:"permissions"`
	url   string
	stars int
}

type Issue struct {
	Url           string `json:"url"`
	RepositoryUrl string `json:"repository_url"`
	LabelsUrl     string `json:"labels_url"`
	CommentsUrl   string `json:"comments_url"`
	EventsUrl     string `json:"events_url"`
	HtmlUrl       string `json:"html_url"`
	Id            int    `json:"id"`
	NodeId        string `json:"node_id"`
	Number        int    `json:"number"`
	Title         string `json:"title"`
	User          struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
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
	} `json:"user"`
	Labels                []interface{} `json:"labels"`
	State                 string        `json:"state"`
	Locked                bool          `json:"locked"`
	Assignee              interface{}   `json:"assignee"`
	Assignees             []interface{} `json:"assignees"`
	Milestone             interface{}   `json:"milestone"`
	Comments              int           `json:"comments"`
	CreatedAt             time.Time     `json:"created_at"`
	UpdatedAt             time.Time     `json:"updated_at"`
	ClosedAt              interface{}   `json:"closed_at"`
	AuthorAssociation     string        `json:"author_association"`
	ActiveLockReason      interface{}   `json:"active_lock_reason"`
	Body                  string        `json:"body"`
	PerformedViaGithubApp interface{}   `json:"performed_via_github_app"`
}

type PullRequest struct {
	Url      string `json:"url"`
	Id       int    `json:"id"`
	NodeId   string `json:"node_id"`
	HtmlUrl  string `json:"html_url"`
	DiffUrl  string `json:"diff_url"`
	PatchUrl string `json:"patch_url"`
	IssueUrl string `json:"issue_url"`
	Number   int    `json:"number"`
	State    string `json:"state"`
	Locked   bool   `json:"locked"`
	Title    string `json:"title"`
	User     struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
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
	} `json:"user"`
	Body               string        `json:"body"`
	CreatedAt          time.Time     `json:"created_at"`
	UpdatedAt          time.Time     `json:"updated_at"`
	ClosedAt           interface{}   `json:"closed_at"`
	MergedAt           interface{}   `json:"merged_at"`
	MergeCommitSha     interface{}   `json:"merge_commit_sha"`
	Assignee           interface{}   `json:"assignee"`
	Assignees          []interface{} `json:"assignees"`
	RequestedReviewers []interface{} `json:"requested_reviewers"`
	RequestedTeams     []interface{} `json:"requested_teams"`
	Labels             []interface{} `json:"labels"`
	Milestone          interface{}   `json:"milestone"`
	Draft              bool          `json:"draft"`
	CommitsUrl         string        `json:"commits_url"`
	ReviewCommentsUrl  string        `json:"review_comments_url"`
	ReviewCommentUrl   string        `json:"review_comment_url"`
	CommentsUrl        string        `json:"comments_url"`
	StatusesUrl        string        `json:"statuses_url"`
	Head               struct {
		Label string `json:"label"`
		Ref   string `json:"ref"`
		Sha   string `json:"sha"`
		User  struct {
			Login             string `json:"login"`
			Id                int    `json:"id"`
			NodeId            string `json:"node_id"`
			AvatarUrl         string `json:"avatar_url"`
			GravatarId        string `json:"gravatar_id"`
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
		} `json:"user"`
		Repo struct {
			Id       int    `json:"id"`
			NodeId   string `json:"node_id"`
			Name     string `json:"name"`
			FullName string `json:"full_name"`
			Private  bool   `json:"private"`
			Owner    struct {
				Login             string `json:"login"`
				Id                int    `json:"id"`
				NodeId            string `json:"node_id"`
				AvatarUrl         string `json:"avatar_url"`
				GravatarId        string `json:"gravatar_id"`
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
			} `json:"owner"`
			HtmlUrl          string      `json:"html_url"`
			Description      string      `json:"description"`
			Fork             bool        `json:"fork"`
			Url              string      `json:"url"`
			ForksUrl         string      `json:"forks_url"`
			KeysUrl          string      `json:"keys_url"`
			CollaboratorsUrl string      `json:"collaborators_url"`
			TeamsUrl         string      `json:"teams_url"`
			HooksUrl         string      `json:"hooks_url"`
			IssueEventsUrl   string      `json:"issue_events_url"`
			EventsUrl        string      `json:"events_url"`
			AssigneesUrl     string      `json:"assignees_url"`
			BranchesUrl      string      `json:"branches_url"`
			TagsUrl          string      `json:"tags_url"`
			BlobsUrl         string      `json:"blobs_url"`
			GitTagsUrl       string      `json:"git_tags_url"`
			GitRefsUrl       string      `json:"git_refs_url"`
			TreesUrl         string      `json:"trees_url"`
			StatusesUrl      string      `json:"statuses_url"`
			LanguagesUrl     string      `json:"languages_url"`
			StargazersUrl    string      `json:"stargazers_url"`
			ContributorsUrl  string      `json:"contributors_url"`
			SubscribersUrl   string      `json:"subscribers_url"`
			SubscriptionUrl  string      `json:"subscription_url"`
			CommitsUrl       string      `json:"commits_url"`
			GitCommitsUrl    string      `json:"git_commits_url"`
			CommentsUrl      string      `json:"comments_url"`
			IssueCommentUrl  string      `json:"issue_comment_url"`
			ContentsUrl      string      `json:"contents_url"`
			CompareUrl       string      `json:"compare_url"`
			MergesUrl        string      `json:"merges_url"`
			ArchiveUrl       string      `json:"archive_url"`
			DownloadsUrl     string      `json:"downloads_url"`
			IssuesUrl        string      `json:"issues_url"`
			PullsUrl         string      `json:"pulls_url"`
			MilestonesUrl    string      `json:"milestones_url"`
			NotificationsUrl string      `json:"notifications_url"`
			LabelsUrl        string      `json:"labels_url"`
			ReleasesUrl      string      `json:"releases_url"`
			DeploymentsUrl   string      `json:"deployments_url"`
			CreatedAt        time.Time   `json:"created_at"`
			UpdatedAt        time.Time   `json:"updated_at"`
			PushedAt         time.Time   `json:"pushed_at"`
			GitUrl           string      `json:"git_url"`
			SshUrl           string      `json:"ssh_url"`
			CloneUrl         string      `json:"clone_url"`
			SvnUrl           string      `json:"svn_url"`
			Homepage         interface{} `json:"homepage"`
			Size             int         `json:"size"`
			StargazersCount  int         `json:"stargazers_count"`
			WatchersCount    int         `json:"watchers_count"`
			Language         string      `json:"language"`
			HasIssues        bool        `json:"has_issues"`
			HasProjects      bool        `json:"has_projects"`
			HasDownloads     bool        `json:"has_downloads"`
			HasWiki          bool        `json:"has_wiki"`
			HasPages         bool        `json:"has_pages"`
			ForksCount       int         `json:"forks_count"`
			MirrorUrl        interface{} `json:"mirror_url"`
			Archived         bool        `json:"archived"`
			Disabled         bool        `json:"disabled"`
			OpenIssuesCount  int         `json:"open_issues_count"`
			License          interface{} `json:"license"`
			Forks            int         `json:"forks"`
			OpenIssues       int         `json:"open_issues"`
			Watchers         int         `json:"watchers"`
			DefaultBranch    string      `json:"default_branch"`
		} `json:"repo"`
	} `json:"head"`
	Base struct {
		Label string `json:"label"`
		Ref   string `json:"ref"`
		Sha   string `json:"sha"`
		User  struct {
			Login             string `json:"login"`
			Id                int    `json:"id"`
			NodeId            string `json:"node_id"`
			AvatarUrl         string `json:"avatar_url"`
			GravatarId        string `json:"gravatar_id"`
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
		} `json:"user"`
		Repo struct {
			Id       int    `json:"id"`
			NodeId   string `json:"node_id"`
			Name     string `json:"name"`
			FullName string `json:"full_name"`
			Private  bool   `json:"private"`
			Owner    struct {
				Login             string `json:"login"`
				Id                int    `json:"id"`
				NodeId            string `json:"node_id"`
				AvatarUrl         string `json:"avatar_url"`
				GravatarId        string `json:"gravatar_id"`
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
			} `json:"owner"`
			HtmlUrl          string      `json:"html_url"`
			Description      string      `json:"description"`
			Fork             bool        `json:"fork"`
			Url              string      `json:"url"`
			ForksUrl         string      `json:"forks_url"`
			KeysUrl          string      `json:"keys_url"`
			CollaboratorsUrl string      `json:"collaborators_url"`
			TeamsUrl         string      `json:"teams_url"`
			HooksUrl         string      `json:"hooks_url"`
			IssueEventsUrl   string      `json:"issue_events_url"`
			EventsUrl        string      `json:"events_url"`
			AssigneesUrl     string      `json:"assignees_url"`
			BranchesUrl      string      `json:"branches_url"`
			TagsUrl          string      `json:"tags_url"`
			BlobsUrl         string      `json:"blobs_url"`
			GitTagsUrl       string      `json:"git_tags_url"`
			GitRefsUrl       string      `json:"git_refs_url"`
			TreesUrl         string      `json:"trees_url"`
			StatusesUrl      string      `json:"statuses_url"`
			LanguagesUrl     string      `json:"languages_url"`
			StargazersUrl    string      `json:"stargazers_url"`
			ContributorsUrl  string      `json:"contributors_url"`
			SubscribersUrl   string      `json:"subscribers_url"`
			SubscriptionUrl  string      `json:"subscription_url"`
			CommitsUrl       string      `json:"commits_url"`
			GitCommitsUrl    string      `json:"git_commits_url"`
			CommentsUrl      string      `json:"comments_url"`
			IssueCommentUrl  string      `json:"issue_comment_url"`
			ContentsUrl      string      `json:"contents_url"`
			CompareUrl       string      `json:"compare_url"`
			MergesUrl        string      `json:"merges_url"`
			ArchiveUrl       string      `json:"archive_url"`
			DownloadsUrl     string      `json:"downloads_url"`
			IssuesUrl        string      `json:"issues_url"`
			PullsUrl         string      `json:"pulls_url"`
			MilestonesUrl    string      `json:"milestones_url"`
			NotificationsUrl string      `json:"notifications_url"`
			LabelsUrl        string      `json:"labels_url"`
			ReleasesUrl      string      `json:"releases_url"`
			DeploymentsUrl   string      `json:"deployments_url"`
			CreatedAt        time.Time   `json:"created_at"`
			UpdatedAt        time.Time   `json:"updated_at"`
			PushedAt         time.Time   `json:"pushed_at"`
			GitUrl           string      `json:"git_url"`
			SshUrl           string      `json:"ssh_url"`
			CloneUrl         string      `json:"clone_url"`
			SvnUrl           string      `json:"svn_url"`
			Homepage         interface{} `json:"homepage"`
			Size             int         `json:"size"`
			StargazersCount  int         `json:"stargazers_count"`
			WatchersCount    int         `json:"watchers_count"`
			Language         string      `json:"language"`
			HasIssues        bool        `json:"has_issues"`
			HasProjects      bool        `json:"has_projects"`
			HasDownloads     bool        `json:"has_downloads"`
			HasWiki          bool        `json:"has_wiki"`
			HasPages         bool        `json:"has_pages"`
			ForksCount       int         `json:"forks_count"`
			MirrorUrl        interface{} `json:"mirror_url"`
			Archived         bool        `json:"archived"`
			Disabled         bool        `json:"disabled"`
			OpenIssuesCount  int         `json:"open_issues_count"`
			License          interface{} `json:"license"`
			Forks            int         `json:"forks"`
			OpenIssues       int         `json:"open_issues"`
			Watchers         int         `json:"watchers"`
			DefaultBranch    string      `json:"default_branch"`
		} `json:"repo"`
	} `json:"base"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Html struct {
			Href string `json:"href"`
		} `json:"html"`
		Issue struct {
			Href string `json:"href"`
		} `json:"issue"`
		Comments struct {
			Href string `json:"href"`
		} `json:"comments"`
		ReviewComments struct {
			Href string `json:"href"`
		} `json:"review_comments"`
		ReviewComment struct {
			Href string `json:"href"`
		} `json:"review_comment"`
		Commits struct {
			Href string `json:"href"`
		} `json:"commits"`
		Statuses struct {
			Href string `json:"href"`
		} `json:"statuses"`
	} `json:"_links"`
	AuthorAssociation   string      `json:"author_association"`
	AutoMerge           interface{} `json:"auto_merge"`
	ActiveLockReason    interface{} `json:"active_lock_reason"`
	Merged              bool        `json:"merged"`
	Mergeable           bool        `json:"mergeable"`
	Rebaseable          bool        `json:"rebaseable"`
	MergeableState      string      `json:"mergeable_state"`
	MergedBy            interface{} `json:"merged_by"`
	Comments            int         `json:"comments"`
	ReviewComments      int         `json:"review_comments"`
	MaintainerCanModify bool        `json:"maintainer_can_modify"`
	Commits             int         `json:"commits"`
	Additions           int         `json:"additions"`
	Deletions           int         `json:"deletions"`
	ChangedFiles        int         `json:"changed_files"`
}

var gitHubBaseUrl = env.Get("GITHUB_BASE_URL")
var gitHubCommonHeader = generateCommonHeader()

func GetMostStarredRepository(userLogin string) Repository {
	repositories, httpStatus := getUserRepositories(userLogin)

	var mostStarredRepo Repository
	if httpStatus == http.StatusOK {
		mostStarredRepo = findMostStarredRepository(repositories)
	}

	return mostStarredRepo
}

func GetMostCommentedIssues(userLogin string, repositoryName string) Issue {
	repositoryIssues := getRepositoryIssues(userLogin, repositoryName)

	mostCommentedIssue := findMostCommentedOpenedIssue(repositoryIssues)

	return mostCommentedIssue
}

func GetNonInteractedPullRequests(userLogin string, repositoryName string) []PullRequest {
	pullRequestsOnRepo := getPullRequests(userLogin, repositoryName)

	nonInteractedPullRequests := filterNonInteractedPullRequests(pullRequestsOnRepo)

	return nonInteractedPullRequests
}

// getUserRepositories get all user repositories
func getUserRepositories(userName string) ([]Repository, int) {
	repositoriesEndpoint := fmt.Sprintf("/users/%s/repos", userName)

	url := gitHubBaseUrl + repositoriesEndpoint

	client := http.Client{}

	preparedRequest, _ := http.NewRequest(http.MethodGet, url, nil)
	preparedRequest.Header = gitHubCommonHeader

	repositoriesResponse, err := client.Do(preparedRequest)
	if err != nil {
		log.Printf("Error na requisição de repositórios.")
		return nil, http.StatusInternalServerError
	}

	if repositoriesResponse.StatusCode == http.StatusForbidden {
		log.Printf("Limite máximo de requests atingidos a GitHub API.")
		return nil, http.StatusInternalServerError
	}

	repositories := makeRepositoriesFromBody(repositoriesResponse.Body)

	return repositories, repositoriesResponse.StatusCode
}

// getPullRequests get all pull requests on a repository
func getPullRequests(userLogin string, repositoryName string) []PullRequest {
	pullRequestsEndpoint :=
		fmt.Sprintf("/repos/%s/%s/pulls", userLogin, repositoryName)
	url := gitHubBaseUrl + pullRequestsEndpoint

	preparedRequest, _ := http.NewRequest(http.MethodGet, url, nil)
	preparedRequest.Header = gitHubCommonHeader

	client := http.Client{}

	pullRequestResponse, err := client.Do(preparedRequest)
	if err != nil {
		return nil
	}

	pullRequests := makePullRequestsFromBody(pullRequestResponse.Body)

	return pullRequests
}

// getRepositoryIssues gets a array of Issue based on a userLogin and
// a repository
func getRepositoryIssues(userLogin string, repositoryName string) []Issue {
	issuesRepositoryEndpoint :=
		fmt.Sprintf("/repos/%s/%s/issues", userLogin, repositoryName)

	url := gitHubBaseUrl + issuesRepositoryEndpoint

	issuesRequest, _ := http.NewRequest(http.MethodGet, url, nil)
	issuesRequest.Header = gitHubCommonHeader

	client := http.Client{}

	issuesResponse, err := client.Do(issuesRequest)
	if err != nil {
		log.Printf("Error na requisição das issues do %s.", repositoryName)
		return nil
	}

	issues := generateIssuesFromBody(issuesResponse.Body)

	return issues
}

// getGitHubUser get a User instance from GitHub API
func getGitHubUser(user string) User {
	formatedEndpoint := fmt.Sprintf("/users/%s", user)
	buildedURL := gitHubBaseUrl + formatedEndpoint

	userRequest, _ := http.NewRequest(http.MethodGet, buildedURL, nil)
	userRequest.Header = gitHubCommonHeader

	client := http.Client{}
	userResponse, err := client.Do(userRequest)
	if err != nil {
		log.Print(err)
	}

	var gitHubUser = generateUserFromBody(userResponse.Body)

	return gitHubUser
}

// GenerateAuthenticationHeader generates the header kv used to authorize
// our requests using the personal access token stored at .env
func generateAuthenticationHeader() (key string, value string) {
	personalToken := env.Get("GITHUB_PERSONAL_ACCESS_TOKEN")

	header := fmt.Sprintf("token %s", personalToken)

	key = "Authentication"
	return key, header
}

// findMostStarredRepository finds the most starred
// repository over a []Repository and returns it
func findMostStarredRepository(repos []Repository) Repository {
	var mostStarred Repository
	starThreshold := 0
	for _, repo := range repos {
		if repo.StargazersCount >= starThreshold {
			mostStarred = repo
			starThreshold = repo.StargazersCount
		}
	}
	return mostStarred
}

// filterNonInteractedPullRequests get only non interacted PullRequest
// in a array of PullRequest
func filterNonInteractedPullRequests(pullRequests []PullRequest) []PullRequest {
	// Nota: Por padrão da API os PR's retornados já estão com status OPEN ;)

	filledPullRequests := fillComments(pullRequests)

	log.Printf("Comments prenchidos")
	for _, request := range filledPullRequests {
		log.Print(request.Comments)
	}

	filteredPullRequests := removeInteractedPullRequests(filledPullRequests)

	return filteredPullRequests
}

// fillComments fill an new array of PullRequest with comments field
func fillComments(pullRequests []PullRequest) []PullRequest {
	var buffer []PullRequest

	for _, pullRequest := range pullRequests {
		buffer = append(buffer, fetchPullRequestComments(&pullRequest))
	}

	return buffer
}

// fetchPullRequestComments get a PullRequest with comments field
func fetchPullRequestComments(pullRequest *PullRequest) PullRequest {
	preparedRequest, _ := http.NewRequest(http.MethodGet, pullRequest.Url, nil)
	preparedRequest.Header = gitHubCommonHeader

	client := http.Client{}

	pullRequestResponse, _ := client.Do(preparedRequest)

	bodyBytes, _ := ioutil.ReadAll(pullRequestResponse.Body)
	buffer := PullRequest{}
	_ = json.Unmarshal(bodyBytes, &buffer)

	return buffer
}

// removeInteractedPullRequests retuns a new PullRequest array with only
// non interacted PullRequests
func removeInteractedPullRequests(prs []PullRequest) []PullRequest {
	var buffer []PullRequest

	for _, pr := range prs {
		if pr.Comments == 0 && pr.State == "open" {
			buffer = append(buffer, pr)
		}
	}

	return buffer
}

// findMostCommentedOpenedIssue find the most commented and opened
// Issue
func findMostCommentedOpenedIssue(issues []Issue) Issue {
	var mostCommentedOpenedIssue Issue
	commentThreshold := 0
	for _, issue := range issues {
		if issue.State == "open" &&
			issue.Comments >= commentThreshold {
			mostCommentedOpenedIssue = issue
			commentThreshold = issue.Comments
		}
	}
	return mostCommentedOpenedIssue
}

// generateCommonHeader generates a common http.Header used to
// make a http.Request to GitHub API
func generateCommonHeader() http.Header {
	hostname, _ := os.Hostname()
	header := http.Header{}
	header.Add(generateAuthenticationHeader())
	header.Add("Accept", "application/vnd.github.v3+json")
	header.Add("Connection", "keep-alive")
	header.Add("Host", hostname)

	return header
}

// generateUserFromBody creates a User from a
// io.Reader generally provided by a http.Request
func generateUserFromBody(requestBody io.Reader) User {
	user := User{}
	bodyBytes, _ := ioutil.ReadAll(requestBody)
	err := json.Unmarshal(bodyBytes, &user)
	if err != nil {
		log.Print(err)
	}
	return user
}

// generateIssuesFromBody creates a array of
// Issue from a io.Reader generally provided
// by a http.Request
func generateIssuesFromBody(requestBody io.Reader) []Issue {
	bodyBytes, _ := ioutil.ReadAll(requestBody)
	issues := make([]Issue, 0)
	err := json.Unmarshal(bodyBytes, &issues)
	if err != nil {
		log.Print(err)
	}
	return issues
}

// makeRepositoriesFromBody creates a array of
// Repository from a io.Reader generally provided
// by a http.Request
func makeRepositoriesFromBody(requestBody io.Reader) []Repository {
	bodyBytes, _ := ioutil.ReadAll(requestBody)
	repos := make([]Repository, 0)
	err := json.Unmarshal(bodyBytes, &repos)
	if err != nil {
		log.Print(err)
	}
	return repos
}

// makePullRequestsFromBody creates a array of
// PullRequest from a io.Reader generally provided
// by a http.Request
func makePullRequestsFromBody(requestBody io.Reader) []PullRequest {
	bodyBytes, _ := ioutil.ReadAll(requestBody)
	pullRequests := make([]PullRequest, 0)
	err := json.Unmarshal(bodyBytes, &pullRequests)
	if err != nil {
		log.Print(err)
	}
	return pullRequests
}
