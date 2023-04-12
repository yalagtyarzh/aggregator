package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (m *Middleware) RecoverMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			r := recover()
			if r != nil {
				stacktrace := string(debug.Stack())
				switch t := r.(type) {
				case string, error:
					err = fmt.Errorf(`panic: %s, stacktrace: %s`, t, stacktrace)
				default:
					err = fmt.Errorf(`unknown panic`)
				}

				m.appLogger.Error(err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()

		h.ServeHTTP(w, r)
	})
}
