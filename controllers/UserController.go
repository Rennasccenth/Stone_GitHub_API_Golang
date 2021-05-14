package controllers

import (
	"Stone_GitHub_API_Golang/cmd"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
)

func GetUserMostStarredRepo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["user"]
	// Testar se retorna nil ou vazio
	if user == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Componentizar isso mais tarde ( passar o endpoint pelo .env )
	formatedEndpoint := fmt.Sprintf("/users/%s/popular-repository", user)
	adress := cmd.GitHubBaseURL + formatedEndpoint

	buildedUrl, _ := url.Parse(adress)

	req := http.Request{
		Method: http.MethodGet,
		URL:    buildedUrl,
		Header: http.Header{
			"chave": {"Valor"},
		},
		ContentLength:    0,
		TransferEncoding: nil,
		RequestURI:       "",
	}

}

func GetUserMostCommentedOpenedIssue(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["user"]
	repository := params["repository"]
}

func GetUserOpenedPullRequests(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["user"]
	repository := params["repository"]
	writable := user + repository
	_, err := w.Write([]byte(writable))
	if err != nil {
		log.Fatal(err)
	}
}
