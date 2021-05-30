package handler

import (
	"reflect"
	"testing"
)

func TestGetUserMostStarredRepository(t *testing.T) {
	userNameTest := "Rennasccenth"
	expectedResult := "Rennasccenth/Stone_GitHub_API_Golang"

	result := GetUserMostStarredRepository(userNameTest)
	evaluatedResult := result.FullName

	if reflect.DeepEqual(evaluatedResult, expectedResult) {
		t.Log("GetMostStarredRepository PASSED")
	} else {
		t.Errorf("GetMostStarredRepository FAILED, expected %v but got %v",
			evaluatedResult, expectedResult)
	}
}
