package main

import (
	"Stone_GitHub_API_Golang/cmd"
	"Stone_GitHub_API_Golang/controllers"
	"Stone_GitHub_API_Golang/env"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/users/{user}/popular-repository",
		controllers.GetUserMostStarredRepo)

	http.HandleFunc("/users/{user}/repositories/{repository}/popular-issue",
		controllers.GetUserMostCommentedOpenedIssue)

	http.HandleFunc("/users/{user}/repositories/{repository}/open-pull-requests",
		controllers.GetUserOpenedPullRequests)

	serveOnDefaultPort()
	cmd.GenerateAuthenticationHeader()
}

// serveOnDefaultPort serves the application on the default port stored
// on .env as APPLICATION_PORT
func serveOnDefaultPort() {
	applicationPort := env.Get("APPLICATION_PORT")
	addr := fmt.Sprintf(":%s", applicationPort)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("Falha ao servir a aplicação na porta %s.", applicationPort)
		os.Exit(1)
	}
}
