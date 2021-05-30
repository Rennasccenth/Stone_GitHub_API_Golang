package handler

import (
	"Stone_GitHub_API_Golang/cmd"
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

var cacheLayer = startCacheLayer()

// startCacheLayer inicialize a Cache layer
func startCacheLayer() *cache.Cache {
	c := cache.New(5*time.Minute, 10*time.Minute)
	return c
}

// GetMostStarredRepository tries to get the most starred cmd.Repository
// on cache layer. If not found, fetches and updates the cache.
func GetMostStarredRepository(userLogin string) cmd.Repository {
	cacheKey := fmt.Sprintf("{Action_1}{%s}", userLogin)
	var mostStarredRepo cmd.Repository

	cachedResponse, found := cacheLayer.Get(cacheKey)
	if found {
		mostStarredRepo, _ = cachedResponse.(cmd.Repository)
		return mostStarredRepo
	}
	mostStarredRepo = cmd.GetUserMostStarredRepository(userLogin)
	cacheLayer.Set(cacheKey, mostStarredRepo, cache.DefaultExpiration)

	return mostStarredRepo
}

// GetMostCommentedIssue tries to get the most commented cmd.Issue
// on cache layer. If not found , fetches and updates the cache
func GetMostCommentedIssue(user string, repository string) cmd.Issue {
	cacheKey := fmt.Sprintf("{Action_2}{%s}{%s}", user, repository)
	var mostCommentedOpenedIssue cmd.Issue

	cachedResponse, found := cacheLayer.Get(cacheKey)
	if found {
		mostCommentedOpenedIssue = cachedResponse.(cmd.Issue)
		return mostCommentedOpenedIssue
	}
	mostCommentedOpenedIssue = cmd.GetMostCommentedIssue(user, repository)
	cacheLayer.Set(cacheKey, mostCommentedOpenedIssue, cache.DefaultExpiration)

	return mostCommentedOpenedIssue
}

// GetNonInteractedPullRequests tries to get all non interacted cmd.PullRequest
// of a cmd.Repository on cache layer. If not found , fetches and updates the cache.
func GetNonInteractedPullRequests(user string, repository string) []cmd.PullRequest {
	cacheKey := fmt.Sprintf("{Action_3}{%s}{%s}", user, repository)
	var nonInteractedPullRequests []cmd.PullRequest

	cachedResponse, found := cacheLayer.Get(cacheKey)
	if found {
		nonInteractedPullRequests, _ = cachedResponse.([]cmd.PullRequest)
		return nonInteractedPullRequests
	}
	nonInteractedPullRequests = cmd.GetNonInteractedPullRequests(user, repository)
	cacheLayer.Set(cacheKey, nonInteractedPullRequests, cache.DefaultExpiration)

	return nonInteractedPullRequests
}
