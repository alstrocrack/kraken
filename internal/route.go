package internal

import (
	"kraken/internal/help"
	"kraken/internal/posts"
	"kraken/internal/users"
	"net/http"
)

const apiPath = "/api/v1"

func StartServer() {
	// internal
	http.HandleFunc("/users/register", users.Register)
	http.HandleFunc("/post", posts.Post)

	// api
	http.HandleFunc(apiPath+"/help", help.Help)
	http.HandleFunc(apiPath+"/users/register", users.Register)
	http.HandleFunc(apiPath+"/post/post", posts.Post)

	http.ListenAndServe("localhost:8080", nil)
}
