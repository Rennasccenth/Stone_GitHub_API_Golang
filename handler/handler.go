package handler

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/rennasccenth/Stone_GitHub_API_Golang/cmd"
	"time"
)

var cacheLayer = startCacheLayer()

// startCacheLayer inicialize a Cache layer
func startCacheLayer() *cache.Cache {
	c := cache.New(5*time.Minute, 10*time.Minute)
	return c
}

// GetUserMostStarredRepository tries to get the most starred cmd.Repository
// on cache layer. If not found, fetches and updates the cache.
func GetUserMostStarredRepository(userLogin string) (mostStarredRepo cmd.Repository, isCached bool) {
	cacheKey := fmt.Sprintf("{Action_1}{%s}", userLogin)
	isCached = false

	cachedResponse, found := cacheLayer.Get(cacheKey)
	if found {
		mostStarredRepo, _ = cachedResponse.(cmd.Repository)
		return mostStarredRepo, true
	}
	mostStarredRepo = cmd.GetUserMostStarredRepository(userLogin)
	cacheLayer.Set(cacheKey, mostStarredRepo, cache.DefaultExpiration)

	return mostStarredRepo, false
}

// GetMostCommentedIssue tries to get the most commented cmd.Issue
// on cache layer. If not found , fetches and updates the cache
func GetMostCommentedIssue(user string, repository string) (mostCommentedOpenedIssue cmd.Issue, isCached bool) {
	cacheKey := fmt.Sprintf("{Action_2}{%s}{%s}", user, repository)
	isCached = false

	cachedResponse, found := cacheLayer.Get(cacheKey)
	if found {
		mostCommentedOpenedIssue = cachedResponse.(cmd.Issue)
		return mostCommentedOpenedIssue, true
	}
	mostCommentedOpenedIssue = cmd.GetMostCommentedIssue(user, repository)
	cacheLayer.Set(cacheKey, mostCommentedOpenedIssue, cache.DefaultExpiration)

	return mostCommentedOpenedIssue, false
}

// GetNonInteractedPullRequests tries to get all non interacted cmd.PullRequest
// of a cmd.Repository on cache layer. If not found , fetches and updates the cache.
func GetNonInteractedPullRequests(user string, repository string) (nonInteractedPullRequests []cmd.PullRequest, isCached bool) {
	cacheKey := fmt.Sprintf("{Action_3}{%s}{%s}", user, repository)
	isCached = false

	cachedResponse, found := cacheLayer.Get(cacheKey)
	if found {
		nonInteractedPullRequests, _ = cachedResponse.([]cmd.PullRequest)
		return nonInteractedPullRequests, true
	}
	nonInteractedPullRequests = cmd.GetNonInteractedPullRequests(user, repository)
	cacheLayer.Set(cacheKey, nonInteractedPullRequests, cache.DefaultExpiration)

	return nonInteractedPullRequests, false
}
