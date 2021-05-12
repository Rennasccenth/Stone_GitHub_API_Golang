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
	router := mux.NewRouter()

	router.HandleFunc("/users/{user}/popular-repository",
		controllers.GetUserMostStarredRepo).Methods(http.MethodGet)

	router.HandleFunc("/users/{user}/repositories/{repository}/popular-issue",
		controllers.GetUserMostCommentedOpenedIssue).Methods(http.MethodGet)

	router.HandleFunc("/users/{user}/repositories/{repository}/open-pull-requests",
		controllers.GetUserOpenedPullRequests).Methods(http.MethodGet)

	serveOnDefaultPort(router)
}

// serveOnDefaultPort serves the handler on the port
// stored on .env as APPLICATION_PORT
func serveOnDefaultPort(handler http.Handler) {
	applicationPort := env.Get("APPLICATION_PORT")
	addr := fmt.Sprintf(":%s", applicationPort)
	log.Printf("Subindo API na porta %s.", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
