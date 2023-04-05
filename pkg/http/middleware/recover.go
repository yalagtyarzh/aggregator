package middleware

import (
	"fmt"
	"github.com/yalagtyarzh/aggregator/pkg/http/helpers"
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
				case string:
					err = helpers.NewError(fmt.Errorf(`panic: %s, stacktrace: %s`, t, stacktrace), "", false)
				case error:
					err = helpers.NewError(fmt.Errorf(`panic: %v, stacktrace: %s`, t, stacktrace), "", false)
				default:
					err = helpers.NewError(fmt.Errorf(`unknown panic`), "", false)
				}

				m.appLogger.Error(err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()

		h.ServeHTTP(w, r)
	})
}
