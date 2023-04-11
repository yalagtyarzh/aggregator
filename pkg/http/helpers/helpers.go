package helpers

import (
	"encoding/json"
	"github.com/yalagtyarzh/aggregator/pkg/http/middleware"
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"net/http"
	"strings"
)

type AppError struct {
	Err *Err `json:"error"`
}

type Err struct {
	Code            int    `json:"code"`
	Message         string `json:"message"`
	IsBusinessError bool   `json:"isBusinessError"`
	Error           string `json:"error"`
}

func NewError(err error, errMsg string, isBusinessError bool) *AppError {
	if err, ok := err.(*AppError); ok {
		return err
	}

	var errStr string
	if err != nil {
		errStr = err.Error()
	}

	return &AppError{
		Err: &Err{
			Message:         strings.ReplaceAll(errMsg, `"`, `\"`),
			IsBusinessError: isBusinessError,
			Error:           errStr,
		},
	}
}

func (e *AppError) Error() string {
	return e.Err.Message + ": " + e.Err.Error
}

type SuperHandler func(w http.ResponseWriter, r *http.Request) *AppError

func CallHandler(h SuperHandler, w http.ResponseWriter, r *http.Request, log logger.ILogger) {
	if err := h(w, r); err != nil {
		log.Errorf("X-Request-Id: %s, %s %s: %s", r.Header.Get(middleware.RequestIdHeader), r.URL.Path, r.Method, err.Error())
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(err.Err.Code)
		b, _ := json.Marshal(err)
		_, _ = w.Write(b)
	}
}
