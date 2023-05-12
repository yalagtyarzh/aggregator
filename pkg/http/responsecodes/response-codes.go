package responsecodes

import (
	"encoding/json"
	"net/http"
)

type errMsg struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

// Write200 OK
func Write200(body []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if len(body) > 0 {
		_, _ = w.Write(body)
	}
}

// Write204 Status No Content
func Write204(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

// Write401 Unauthorized
func Write401(err string, writer http.ResponseWriter) {
	msg := errMsg{Error: struct {
		Message string `json:"message"`
	}(struct{ Message string }{Message: err})}
	b, _ := json.Marshal(msg)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)
	_, _ = writer.Write(b)
}
