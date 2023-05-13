package main

import (
	"fmt"
	"kraken/db"
	"log"
)

func main() {
	db, err := db.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)
}
