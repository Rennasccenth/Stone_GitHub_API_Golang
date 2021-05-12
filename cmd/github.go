package cmd

import "Stone_GitHub_API_Golang/env"


// GenerateAuthenticationHeader using the personal access token at .env
func GenerateAuthenticationHeader()  {
	personalToken := env.Get("GITHUB_PERSONAL_ACCESS_TOKEN")
	header :=
	{
	Authorization: personalToken
	}
	return header
}
