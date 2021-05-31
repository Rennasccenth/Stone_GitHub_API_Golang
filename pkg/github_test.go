package pkg

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
			expectedResult, evaluatedResult)
	}
}

func TestGetMostCommentedIssue(t *testing.T) {
	userNameTest := "Rennasccenth"
	repositoryNameTest := "Stone_GitHub_API_Golang"
	expectedResult := "Issue com mais comentários"

	result := GetMostCommentedIssue(userNameTest, repositoryNameTest)
	evaluatedResult := result.Title

	if reflect.DeepEqual(evaluatedResult, expectedResult) {
		t.Log("GetGetMostCommentedIssue PASSED")
	} else {
		t.Errorf("GetGetMostCommentedIssue FAILED, expected %v but got %v",
			expectedResult, evaluatedResult)
	}
}

func TestGetNonInteractedPullRequests(t *testing.T) {
	userNameTest := "Rennasccenth"
	repositoryNameTest := "Stone_GitHub_API_Golang"

	result := GetNonInteractedPullRequests(userNameTest, repositoryNameTest)
	expectedResult := 0
	alreadyInteractedCounter := 0

	for _, pr := range result {
		if pr.Comments != 0 {
			alreadyInteractedCounter++
		}
	}

	if reflect.DeepEqual(alreadyInteractedCounter, 0) {
		t.Log("GetNonInteractedPullRequests PASSED")
	} else {
		t.Errorf("GetNonInteractedPullRequests FAILED, expected %d but got %d repositories interacted.",
			expectedResult, alreadyInteractedCounter)
	}
}

func TestFindMostStarredRepository(t *testing.T) {
	var mockedRepositories []Repository

	repo1 := Repository{
		FullName:        "Repo com menos estrelas.",
		StargazersCount: 1,
	}
	repo2 := Repository{
		FullName:        "Repo com um pouco mais de estrelas.",
		StargazersCount: 500,
	}
	repo3 := Repository{
		FullName:        "Repo vencedor com maior número de estrelas.",
		StargazersCount: 987,
	}

	mockedRepositories = append(mockedRepositories, repo1, repo2, repo3)

	mostStarredRepository := findMostStarredRepository(mockedRepositories)
	result := mostStarredRepository.StargazersCount
	expectedResult := repo3.StargazersCount

	if reflect.DeepEqual(result, expectedResult) {
		t.Log("findMostStarredRepository PASSED")
	} else {
		t.Errorf("findMostStarredRepository FAILED, expected %d stars but got %d instead.",
			expectedResult, result)
	}

}

func TestFindMostCommentedOpenedIssue(t *testing.T) {
	var mockedIssues []Issue

	issue1 := Issue{
		Title:    "Menos comentada",
		State:    "open",
		Comments: 4,
	}
	issue2 := Issue{
		Title:    "Quantidade mediana de comentários",
		State:    "open",
		Comments: 80,
	}
	issue3 := Issue{
		Title:    "Mais comentada com status aberto (famoso chat do UOL)",
		State:    "open",
		Comments: 900,
	}
	issue4 := Issue{
		Title:    "Quantidade alta de comentários mas está fechada",
		State:    "closed",
		Comments: 870,
	}
	issue5 := Issue{
		Title:    "Quantidade Altíssima de comentários porém bloqueada",
		State:    "closed",
		Comments: 8001,
	}

	mockedIssues = append(mockedIssues, issue1, issue2, issue3, issue4, issue5)

	winnerIssue := findMostCommentedOpenedIssue(mockedIssues)
	result := winnerIssue.Comments
	expectedResult := issue3.Comments

	if reflect.DeepEqual(result, expectedResult) {
		t.Log("findMostCommentedOpenedIssue PASSED")
	} else {
		t.Errorf("findMostCommentedOpenedIssue FAILED, expected %d comments but got %d instead.",
			expectedResult, result)
	}
}

func TestFilterNonInteractedPullRequests(t *testing.T) {
	var mockedPullRequests []PullRequest

	pr1 := PullRequest{
		Title:    "Pull Request Avulso",
		State:    "open",
		Comments: 10,
	}
	pr2 := PullRequest{
		Title:    "Pull Request com status fechado",
		State:    "closed",
		Comments: 100,
	}
	pr3 := PullRequest{
		Title:    "Pull Request com status fechado e sem interação",
		State:    "closed",
		Comments: 0,
	}
	pr4 := PullRequest{
		Title:    "Pull Request sem interação numero um",
		State:    "open",
		Comments: 0,
	}
	pr5 := PullRequest{
		Title:    "Pull Request sem interação numero dois",
		State:    "open",
		Comments: 0,
	}
	pr6 := PullRequest{
		Title:    "Pull Request sem interação numero três",
		State:    "open",
		Comments: 0,
	}

	mockedPullRequests = append(mockedPullRequests, pr1, pr2, pr3, pr4, pr5, pr6)

	filteredPullRequests := filterNonInteractedPullRequests(mockedPullRequests)

	interactedPullRequests := 0

	for _, pr := range filteredPullRequests {
		if pr.Comments > 0 || pr.State != "open" {
			interactedPullRequests++
		}
	}

	result := interactedPullRequests
	expectedResult := 0

	if reflect.DeepEqual(result, expectedResult) {
		t.Log("filterNonInteractedPullRequests PASSED")
	} else {
		t.Errorf("filterNonInteractedPullRequests FAILED, expected %d interacted Pull Requests but got %d instead.",
			expectedResult, result)
	}

}
