package handler

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/rennasccenth/Stone_GitHub_API_Golang/pgk"
	"time"
)

var cacheLayer = startCacheLayer()

// startCacheLayer inicialize a Cache layer
func startCacheLayer() *cache.Cache {
	c := cache.New(5*time.Minute, 10*time.Minute)
	return c
}

// GetUserMostStarredRepository tries to get the most starred pgk.Repository
// on cache layer. If not found, fetches and updates the cache.
func GetUserMostStarredRepository(userLogin string) (mostStarredRepo pgk.Repository, isCached bool) {
	cacheKey := fmt.Sprintf("{Action_1}{%s}", userLogin)
	isCached = false

	cachedResponse, found := cacheLayer.Get(cacheKey)
	if found {
		mostStarredRepo, _ = cachedResponse.(pgk.Repository)
		return mostStarredRepo, true
	}
	mostStarredRepo = pgk.GetUserMostStarredRepository(userLogin)
	cacheLayer.Set(cacheKey, mostStarredRepo, cache.DefaultExpiration)

	return mostStarredRepo, false
}

// GetMostCommentedIssue tries to get the most commented pgk.Issue
// on cache layer. If not found , fetches and updates the cache
func GetMostCommentedIssue(user string, repository string) (mostCommentedOpenedIssue pgk.Issue, isCached bool) {
	cacheKey := fmt.Sprintf("{Action_2}{%s}{%s}", user, repository)
	isCached = false

	cachedResponse, found := cacheLayer.Get(cacheKey)
	if found {
		mostCommentedOpenedIssue = cachedResponse.(pgk.Issue)
		return mostCommentedOpenedIssue, true
	}
	mostCommentedOpenedIssue = pgk.GetMostCommentedIssue(user, repository)
	cacheLayer.Set(cacheKey, mostCommentedOpenedIssue, cache.DefaultExpiration)

	return mostCommentedOpenedIssue, false
}

// GetNonInteractedPullRequests tries to get all non interacted pgk.PullRequest
// of a pgk.Repository on cache layer. If not found , fetches and updates the cache.
func GetNonInteractedPullRequests(user string, repository string) (nonInteractedPullRequests []pgk.PullRequest, isCached bool) {
	cacheKey := fmt.Sprintf("{Action_3}{%s}{%s}", user, repository)
	isCached = false

	cachedResponse, found := cacheLayer.Get(cacheKey)
	if found {
		nonInteractedPullRequests, _ = cachedResponse.([]pgk.PullRequest)
		return nonInteractedPullRequests, true
	}
	nonInteractedPullRequests = pgk.GetNonInteractedPullRequests(user, repository)
	cacheLayer.Set(cacheKey, nonInteractedPullRequests, cache.DefaultExpiration)

	return nonInteractedPullRequests, false
}
