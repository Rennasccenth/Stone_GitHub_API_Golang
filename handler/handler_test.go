package handler

import (
	"github.com/dchest/uniuri"
	"reflect"
	"testing"
)

func TestCacheOnGetUserMostStarredRepository(t *testing.T) {
	mockRandomUsername := uniuri.NewLen(10)

	_, cacheKey := GetUserMostStarredRepository(mockRandomUsername)

	_, found := cacheLayer.Get(cacheKey)

	if reflect.DeepEqual(true, found) {
		t.Log("CacheOnGetUserMostStarredRepository PASSED")
	} else {
		t.Errorf("CacheOnGetUserMostStarredRepository FAILED, expected %v but got %v",
			true, found)
	}
}

func TestCacheGetMostCommentedIssue(t *testing.T) {
	mockRandomUsername := uniuri.NewLen(10)
	mockRandomRepository := uniuri.NewLen(10)

	_, cacheKey := GetMostCommentedIssue(mockRandomUsername, mockRandomRepository)

	_, found := cacheLayer.Get(cacheKey)

	if reflect.DeepEqual(true, found) {
		t.Log("TestCacheGetMostCommentedIssue PASSED")
	} else {
		t.Errorf("TestCacheGetMostCommentedIssue FAILED, expected %v but got %v",
			true, found)
	}
}

func TestCacheGetNonInteractedPullRequests(t *testing.T) {
	mockRandomUsername := uniuri.NewLen(10)
	mockRandomRepository := uniuri.NewLen(10)

	_, cacheKey := GetNonInteractedPullRequests(mockRandomUsername, mockRandomRepository)

	_, found := cacheLayer.Get(cacheKey)

	if reflect.DeepEqual(found, true) {
		t.Log("TestCacheGetNonInteractedPullRequests PASSED")
	} else {
		t.Errorf("TestCacheGetNonInteractedPullRequests FAILED, expected %v but got %v",
			true, found)
	}
}
