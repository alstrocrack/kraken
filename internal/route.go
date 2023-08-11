package internal

import (
	"kraken/internal/help"
	"kraken/internal/posts"
	"kraken/internal/users"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/help", help.Help)
	http.HandleFunc("/register", users.Register)
	http.HandleFunc("/post", posts.Post)

	http.ListenAndServe("localhost:8080", nil)
}
