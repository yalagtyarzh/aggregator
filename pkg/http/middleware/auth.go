package middleware

import (
	"context"
	"github.com/google/uuid"
	"github.com/yalagtyarzh/aggregator/pkg/http/responsecodes"
	"net/http"
)

func (m *Middleware) AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := uuid.Parse(r.Header.Get("X-User-Id"))
		if err != nil {
			responsecodes.Write401("unauthorized", w)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), "userID", userID))
		h.ServeHTTP(w, r)
	})
}
