package server

import (
	"kraken/help"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/help", help.Help)
	http.ListenAndServe(":8080", nil)
}
