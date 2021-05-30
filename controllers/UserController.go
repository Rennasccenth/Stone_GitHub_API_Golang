package controllers

import (
	handler "Stone_GitHub_API_Golang/handler"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func GetUserMostStarredRepo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userParam := params["user"]

	mostStarredRepo := handler.GetUserMostStarredRepository(userParam)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Date", time.Now().String())

	marshalResp, _ := json.Marshal(mostStarredRepo)
	_, err := w.Write(marshalResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}
}

func GetUserMostCommentedOpenedIssue(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["user"]
	repository := params["repository"]

	mostCommentedOpenedIssue := handler.GetMostCommentedIssue(user, repository)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Date", time.Now().String())

	marshalResp, _ := json.Marshal(mostCommentedOpenedIssue)

	_, err := w.Write(marshalResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}
}

func GetUserOpenedPullRequests(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["user"]
	repository := params["repository"]

	nonInteractedPullRequests := handler.GetNonInteractedPullRequests(user, repository)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Date", time.Now().String())

	marshalResp, _ := json.Marshal(nonInteractedPullRequests)

	_, err := w.Write(marshalResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}
}
