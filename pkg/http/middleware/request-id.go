package middleware

import (
	"github.com/google/uuid"
	"net/http"
)

const RequestIdHeader = "X-Request-Id"

func (m *Middleware) ReqIDMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := r.Header.Get(RequestIdHeader)

		if _, err := uuid.Parse(requestId); err != nil {
			UUID, _ := uuid.NewRandom()
			w.Header().Set(RequestIdHeader, UUID.String())
		} else {
			w.Header().Set(RequestIdHeader, requestId)
		}

		h.ServeHTTP(w, r)
	})
}
