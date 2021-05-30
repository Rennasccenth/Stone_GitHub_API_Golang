package controllers

import (
	"Stone_GitHub_API_Golang/cmd"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	"log"
	"net/http"
	"time"
)

var cacheLayer = startCacheLayer()

func startCacheLayer() *cache.Cache {
	c := cache.New(5*time.Minute, 10*time.Minute)
	return c
}

func GetUserMostStarredRepo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userParam := params["user"]
	cacheKey := fmt.Sprintf("{Action_1}{%s}", userParam)
	var mostStarredRepo cmd.Repository

	cachedResponse, found := cacheLayer.Get(cacheKey)
	if found {
		mostStarredRepo, _ = cachedResponse.(cmd.Repository)
	} else {
		mostStarredRepo = cmd.GetMostStarredRepository(userParam)
		cacheLayer.Set(cacheKey, mostStarredRepo, cache.DefaultExpiration)
	}

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
	cacheKey := fmt.Sprintf("{Action_2}{%s}{%s}", user, repository)

	var mostCommentedOpenedIssue cmd.Issue

	cachedResponse, found := cacheLayer.Get(cacheKey)
	if found {
		mostCommentedOpenedIssue = cachedResponse.(cmd.Issue)
	} else {
		mostCommentedOpenedIssue = cmd.GetMostCommentedIssues(user, repository)
		cacheLayer.Set(cacheKey, mostCommentedOpenedIssue, cache.DefaultExpiration)
	}

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
	cacheKey := fmt.Sprintf("{Action_3}{%s}{%s}", user, repository)
	var nonInteractedPullRequests []cmd.PullRequest

	cachedResponse, found := cacheLayer.Get(cacheKey)
	if found {
		nonInteractedPullRequests, _ = cachedResponse.([]cmd.PullRequest)
	} else {
		nonInteractedPullRequests = cmd.GetNonInteractedPullRequests(user, repository)
		cacheLayer.Set(cacheKey, nonInteractedPullRequests, cache.DefaultExpiration)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Date", time.Now().String())

	marshalResp, _ := json.Marshal(nonInteractedPullRequests)

	_, err := w.Write(marshalResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}
}
