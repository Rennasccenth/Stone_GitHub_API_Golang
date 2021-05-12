package controllers

import (
	"fmt"
	"net/http"
)

func GetUserMostStarredRepos(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Printf("Método GET")
	}

	w.WriteHeader(200)
}
