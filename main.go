package main

import (
	"Stone_GitHub_API_Golang/controllers"
	"Stone_GitHub_API_Golang/env"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	routerHander := mux.NewRouter()

	routerHander.HandleFunc("/users/{user}/popular-repository",
		controllers.GetUserMostStarredRepo).Methods(http.MethodGet)

	routerHander.HandleFunc("/users/{user}/repositories/{repository}/popular-issue",
		controllers.GetUserMostCommentedOpenedIssue).Methods(http.MethodGet)

	routerHander.HandleFunc("/users/{user}/repositories/{repository}/open-pull-requests",
		controllers.GetUserOpenedPullRequests).Methods(http.MethodGet)

	serveOnDefaultPort(routerHander)
}

// serveOnDefaultPort serves the handler on the port
// stored on .env as APPLICATION_PORT
func serveOnDefaultPort(handler http.Handler) {
	applicationPort := env.Get("APPLICATION_PORT")
	if applicationPort == "" {
		applicationPort = "8080"
	}
	addr := fmt.Sprintf(":%s", applicationPort)
	log.Printf("Subindo API na porta %s.", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
