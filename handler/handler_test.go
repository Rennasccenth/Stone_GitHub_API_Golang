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
