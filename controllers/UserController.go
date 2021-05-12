package controllers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetUserMostStarredRepo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["user"]
	_, err := w.Write([]byte(user))
	if err != nil {
		log.Fatal(err)
	}
}

func GetUserMostCommentedOpenedIssue(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["user"]
	repository := params["repository"]
	writable := user + repository
	_, err := w.Write([]byte(writable))
	if err != nil {
		log.Fatal(err)
	}
}

func GetUserOpenedPullRequests(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["user"]
	repository := params["repository"]
	writable := user + repository
	_, err := w.Write([]byte(writable))
	if err != nil {
		log.Fatal(err)
	}
}
