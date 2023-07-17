package help

import (
	"io"
	"net/http"
)

const help = `
	GET /help:
		Desc: Display help,
	
	POST /register:
		Desc: Register as a user
		Request Header: "Content-Type: application/json"
		Request Body:
			email: {string},
			password: {string},
			name: {string}
`

func Help(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, help)
}
