package main

import (
	"fmt"
	"kraken/app/db"
	"kraken/app/server"
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
