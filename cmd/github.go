package cmd

import (
	"Stone_GitHub_API_Golang/env"
	"fmt"
	"net/http"
	"time"
)

var GitHubClient = NewDefaultGitHubClient()
var GitHubBaseURL = env.Get("GIT_HUB_BASE_URL")

var AuthHeader
// GenerateAuthenticationHeader generates the header used to authorize
// our requests using the personal access token stored at .env
func GenerateAuthenticationHeader() (key string, value []string) {
	personalToken := env.Get("GITHUB_PERSONAL_ACCESS_TOKEN")
	header := []strings.
	headerValue := fmt.Sprintf("token %s", personalToken)

	key = "Authentication"
	return key, append([]string, header)
}

// NewDefaultGitHubClient creates a personalized http.Client to handle
// our http requests to GitHubAPI
func NewDefaultGitHubClient() http.Client {
	gitHubClient := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       time.Duration(10000),
	}
	return gitHubClient
}