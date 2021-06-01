package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rennasccenth/Stone_GitHub_API_Golang/handler"
	"log"
	"net/http"
	"time"
)

func GetUserMostStarredRepo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userParam := params["user"]

	mostStarredRepo, _ := handler.GetUserMostStarredRepository(userParam)

	setDefaultHeader(w)

	setResponse(mostStarredRepo, w)
}

func GetUserMostCommentedOpenedIssue(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["user"]
	repository := params["repository"]

	mostCommentedOpenedIssue, _ := handler.GetMostCommentedIssue(user, repository)

	setDefaultHeader(w)

	setResponse(mostCommentedOpenedIssue, w)
}

func GetUserOpenedPullRequests(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["user"]
	repository := params["repository"]

	nonInteractedPullRequests, _ := handler.GetNonInteractedPullRequests(user, repository)

	setDefaultHeader(w)

	setResponse(nonInteractedPullRequests, w)

}

// setDefaultHeader set default Header of this Controller
func setDefaultHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Date", time.Now().String())
}

// setResponse set the commom response of this Controller
func setResponse(data interface{}, w http.ResponseWriter) {
	marshalResp, _ := json.Marshal(data)

	_, err := w.Write(marshalResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}
}
