package main

import (
	"fmt"
	"kraken/db"
	"kraken/server"
	"log"
)

func main() {
	db, err := db.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)

	server.StartServer()
}
