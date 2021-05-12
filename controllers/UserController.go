package controllers

import (
	"fmt"
	"net/http"
)

func GetUserMostStarredRepo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "delete called"}`))
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func GetUserMostCommentedOpenedIssue(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Printf("Método GET")
	}

	w.WriteHeader(200)
}
func GetUserOpenedPullRequests(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Printf("Método GET")
	}

	w.WriteHeader(200)
}
