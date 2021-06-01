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

// GetUserMostStarredRepository tries to get the most starred cmd.Repository
// on cache layer. If not found, fetches and updates the cache.
func GetUserMostStarredRepository(userLogin string) (mostStarredRepo cmd.Repository, cacheKey string) {
	cacheKey = fmt.Sprintf("{Action_1}{%s}", userLogin)

	cachedResponse, found := cacheLayer.Get(cacheKey)
	if found {
		mostStarredRepo, _ = cachedResponse.(cmd.Repository)
		return mostStarredRepo, cacheKey
	}
	mostStarredRepo = cmd.GetUserMostStarredRepository(userLogin)
	cacheLayer.Set(cacheKey, mostStarredRepo, cache.DefaultExpiration)

	return mostStarredRepo, cacheKey
}

// GetMostCommentedIssue tries to get the most commented cmd.Issue
// on cache layer. If not found , fetches and updates the cache
func GetMostCommentedIssue(user string, repository string) (mostCommentedOpenedIssue cmd.Issue, cacheKey string) {
	cacheKey = fmt.Sprintf("{Action_2}{%s}{%s}", user, repository)

	cachedResponse, found := cacheLayer.Get(cacheKey)
	if found {
		mostCommentedOpenedIssue = cachedResponse.(cmd.Issue)
		return mostCommentedOpenedIssue, cacheKey
	}
	mostCommentedOpenedIssue = cmd.GetMostCommentedIssue(user, repository)
	cacheLayer.Set(cacheKey, mostCommentedOpenedIssue, cache.DefaultExpiration)

	return mostCommentedOpenedIssue, cacheKey
}

// GetNonInteractedPullRequests tries to get all non interacted cmd.PullRequest
// of a cmd.Repository on cache layer. If not found , fetches and updates the cache.
func GetNonInteractedPullRequests(user string, repository string) (nonInteractedPullRequests []cmd.PullRequest, cacheKey string) {
	cacheKey = fmt.Sprintf("{Action_3}{%s}{%s}", user, repository)

	cachedResponse, found := cacheLayer.Get(cacheKey)
	if found {
		nonInteractedPullRequests, _ = cachedResponse.([]cmd.PullRequest)
		return nonInteractedPullRequests, cacheKey
	}
	nonInteractedPullRequests = cmd.GetNonInteractedPullRequests(user, repository)
	cacheLayer.Set(cacheKey, nonInteractedPullRequests, cache.DefaultExpiration)

	return nonInteractedPullRequests, cacheKey
}
