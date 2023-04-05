package middleware

import (
	"github.com/yalagtyarzh/aggregator/pkg/http/responsecodes"
	"net/http"
)

func (m *Middleware) AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("Some-Thing")
		if userID == "" {
			responsecodes.Write401("unauthorized", w)
			return
		}
		h.ServeHTTP(w, r)
	})
}
