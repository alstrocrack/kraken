package main

import (
	"kraken/config/logging"
	server "kraken/internal"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	logging.LoggingInit("../../log.txt")

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalln(err)
	}

	server.StartServer()
}
