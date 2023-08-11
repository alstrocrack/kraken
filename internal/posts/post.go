package posts

import (
	"encoding/json"
	"io"
	"kraken/config/db"
	"log"
	"net/http"
)

type PostContent struct {
	Content string
}

func Post(w http.ResponseWriter, r *http.Request) {
	post := &PostContent{}

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
