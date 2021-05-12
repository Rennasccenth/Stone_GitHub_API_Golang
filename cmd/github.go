package cmd

import (
	"Stone_GitHub_API_Golang/env"
	"fmt"
)

// GenerateAuthenticationHeader generates the header used to authorize our requests using the personal access token stored at .env
func GenerateAuthenticationHeader() (key string, value string) {
	personalToken := env.Get("GITHUB_PERSONAL_ACCESS_TOKEN")

	header := fmt.Sprintf("token %s", personalToken)

	key = "Authentication"
	return key, header
}
