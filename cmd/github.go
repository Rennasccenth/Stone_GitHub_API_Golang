package cmd

import "awesomeProject/env"


// GenerateAuthenticationHeader using the personal access token on .env
func GenerateAuthenticationHeader()  {
	personalToken := env.Get("GITHUB_PERSONAL_ACCESS_TOKEN")
	header :=
	{
	Authorization: personalToken
	}
	return header
}
