package responsecodes

import (
	"encoding/json"
	"net/http"
)

type errMsg struct {
	Message string `json:"message"`
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

// Write400 Status Bad Request
func Write400(err string, writer http.ResponseWriter) {
	msg := errMsg{err}
	b, _ := json.Marshal(msg)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	_, _ = writer.Write(b)
}

// Write401 Unauthorized
func Write401(err string, writer http.ResponseWriter) {
	msg := errMsg{err}
	b, _ := json.Marshal(msg)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)
	_, _ = writer.Write(b)
}

// Write403 Status Forbidden
func Write403(err string, writer http.ResponseWriter) {
	msg := errMsg{err}
	b, _ := json.Marshal(msg)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusForbidden)
	_, _ = writer.Write(b)
}

func Write404(err string, writer http.ResponseWriter) {
	msg := errMsg{err}
	b, _ := json.Marshal(msg)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNotFound)
	_, _ = writer.Write(b)
}

// Write500 Status Internal Server Error
func Write500(err string, writer http.ResponseWriter) {
	msg := errMsg{err}
	b, _ := json.Marshal(msg)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)
	_, _ = writer.Write(b)
}
