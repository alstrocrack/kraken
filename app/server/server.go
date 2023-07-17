package server

import (
	"encoding/json"
	"fmt"
	"kraken/app/db"
	"kraken/app/help"
	"net/http"
)

type RegisterItem struct {
	Email    string
	Password string
	Name     string
}

func register(params []byte) {
	items := &RegisterItem{}
	err := json.Unmarshal([]byte(params), items)
	if err != nil {
		fmt.Println(err)
	}
	con, err := db.OpenDB()
	if err != nil {
		fmt.Println(err)
	}
	query := "INSERT INTO user_accounts (name, email, password_hash, created_at, updated_at) VALUES (?, ?, ?, Now(), Now());"
	con.Exec(query, items.Name, items.Email, items.Password)
}

func StartServer() {
	http.HandleFunc("/help", help.Help)
	http.ListenAndServe(":8080", nil)

	http.HandleFunc("/register", register(params))
}
