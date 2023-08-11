package server

import (
	"encoding/json"
	"io"
	"kraken/config/db"
	"kraken/internal/help"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type RegisterItem struct {
	Email    string
	Password string
	Name     string
}

type Post struct {
	Content string
}

func register(w http.ResponseWriter, r *http.Request) {
	item := &RegisterItem{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "ERROR: Cannot read request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &item)
	if err != nil {
		log.Println(err)
		http.Error(w, "ERROR: Cannot parse json", http.StatusInternalServerError)
		return
	}

	con, err := db.OpenDB()
	if err != nil {
		log.Println(err)
		http.Error(w, "ERROR: Cannot connect DB", http.StatusInternalServerError)
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(item.Password), 10)
	if err != nil {
		log.Println(err)
		http.Error(w, "ERROR: Cannot generate hash from password", http.StatusInternalServerError)
		return
	}

	query := "INSERT INTO user_accounts (name, email, password_hash, created_at, updated_at) VALUES (?, ?, ?, Now(), Now());"
	_, err = con.Exec(query, item.Name, item.Email, string(passwordHash))
	if err != nil {
		log.Println(err)
		http.Error(w, "ERROR: Cannot insert new record", http.StatusInternalServerError)
		return
	}

	log.Println("SUCCESS: Register")
	w.WriteHeader(http.StatusOK)
}

func post(w http.ResponseWriter, r *http.Request) {
	post := &Post{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "ERROR: Cannot read request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &post)
	if err != nil {
		log.Println(err)
		http.Error(w, "ERROR: Cannot read request body", http.StatusBadRequest)
		return
	}

	con, err := db.OpenDB()
	if err != nil {
		log.Println(err)
		http.Error(w, "ERROR: Cannot connect DB", http.StatusInternalServerError)
		return
	}

	query := "INSERT INTO posts (user_account_id, content, created_at, updated_at) VALUES (?, ?, Now(), Now());"
	_, err = con.Exec(query, 1, post.Content)
	if err != nil {
		log.Println(err)
		http.Error(w, "ERROR: Cannot insert new record", http.StatusInternalServerError)
		return
	}
}

func StartServer() {
	http.HandleFunc("/help", help.Help)
	http.HandleFunc("/register", register)
	http.HandleFunc("/post", post)

	http.ListenAndServe("localhost:8080", nil)
}
