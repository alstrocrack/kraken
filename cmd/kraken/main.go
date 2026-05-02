package main

import (
	"fmt"
	"kraken/config/logging"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logging.LoggingInit("../../log.txt")
		fmt.Fprintf(w, "Hello, World!")
	})

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server on :8080")
	server := &http.Server{Addr: ":8080"}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}