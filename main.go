package main

import (
	"kraken/app/server"
	"kraken/logging"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	logging.LoggingInit("log.txt")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}

	server.StartServer()
}
