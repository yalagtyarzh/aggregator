package middleware

import (
	"context"
	"github.com/yalagtyarzh/aggregator/pkg/http/responsecodes"
	"net/http"
	"strings"
)

func (m *Middleware) AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			responsecodes.Write401("unauthorized", w)
			return
		}

		splitted := strings.Split(header, " ")
		if len(splitted) != 2 {
			responsecodes.Write401("unauthorized", w)
			return
		}

		payload, err := m.appJWT.ValidateAccessToken(splitted[1])
		if err != nil {
			responsecodes.Write401("unauthorized", w)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), "token", payload))
		h.ServeHTTP(w, r)
	})
}
