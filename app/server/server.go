package server

import (
	"encoding/json"
	"fmt"
	"io"
	"kraken/app/db"
	"kraken/app/help"
	"net/http"
)

type RegisterItem struct {
	Email    string
	Password string
	Name     string
}

func register(w http.ResponseWriter, r *http.Request) {
	item := &RegisterItem{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &item)
	if err != nil {
		fmt.Println(err)
	}
	con, err := db.OpenDB()
	if err != nil {
		fmt.Println(err)
	}
	query := "INSERT INTO user_accounts (name, email, password_hash, created_at, updated_at) VALUES (?, ?, ?, Now(), Now());"
	con.Exec(query, item.Name, item.Email, item.Password)
}

func StartServer() {
	http.HandleFunc("/help", help.Help)
	http.HandleFunc("/register", register)

	http.ListenAndServe("localhost:8080", nil)
	// http.ListenAndServe(":8080", nil)
}
