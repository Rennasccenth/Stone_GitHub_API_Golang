package handler

import (
	"github.com/dchest/uniuri"
	"reflect"
	"testing"
)

func TestCacheOnGetUserMostStarredRepository(t *testing.T) {
	mockRandomUsername := uniuri.NewLen(10)

	_, isCached := GetUserMostStarredRepository(mockRandomUsername)

	_, isCached = GetUserMostStarredRepository(mockRandomUsername)

	if reflect.DeepEqual(isCached, true) {
		t.Log("CacheOnGetUserMostStarredRepository PASSED")
	} else {
		t.Errorf("CacheOnGetUserMostStarredRepository FAILED, expected %v but got %v",
			true, isCached)
	}
}

func TestCacheGetMostCommentedIssue(t *testing.T) {
	mockRandomUsername := uniuri.NewLen(10)
	mockRandomRepository := uniuri.NewLen(10)

	_, isCached := GetMostCommentedIssue(mockRandomUsername, mockRandomRepository)
	_, isCached = GetMostCommentedIssue(mockRandomUsername, mockRandomRepository)

	if reflect.DeepEqual(isCached, true) {
		t.Log("TestCacheGetMostCommentedIssue PASSED")
	} else {
		t.Errorf("TestCacheGetMostCommentedIssue FAILED, expected %v but got %v",
			true, isCached)
	}
}

func TestCacheGetNonInteractedPullRequests(t *testing.T) {
	mockRandomUsername := uniuri.NewLen(10)
	mockRandomRepository := uniuri.NewLen(10)

	_, isCached := GetNonInteractedPullRequests(mockRandomUsername, mockRandomRepository)
	_, isCached = GetNonInteractedPullRequests(mockRandomUsername, mockRandomRepository)

	if reflect.DeepEqual(isCached, true) {
		t.Log("TestCacheGetNonInteractedPullRequests PASSED")
	} else {
		t.Errorf("TestCacheGetNonInteractedPullRequests FAILED, expected %v but got %v",
			true, isCached)
	}
}
