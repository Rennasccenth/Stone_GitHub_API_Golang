package main

import (
	"Stone_GitHub_API_Golang/cmd"
	"Stone_GitHub_API_Golang/controllers"
	"Stone_GitHub_API_Golang/env"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/users/", controllers.GetUserMostStarredRepos)
	http.HandleFunc("/users", controllers.GetUserMostStarredRepos)
	http.HandleFunc("/user", controllers.GetUserMostStarredRepos)

	serveOnDefaultPort()
	cmd.GenerateAuthenticationHeader()
}

// serveOnDefaultPort serve the application on the default port stored on .env as APPLICATION_PORT
func serveOnDefaultPort() {
	applicationPort := env.Get("APPLICATION_PORT")
	addr := fmt.Sprintf(":%s", applicationPort)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("Não foi possível subir o servidor na porta %s", applicationPort)
	}
}
