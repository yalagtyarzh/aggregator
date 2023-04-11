package user_api

import (
	"errors"
	"github.com/yalagtyarzh/aggregator/pkg/http/helpers"
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"net/http"
)

// query params
const (
	productId = "pid"
)

// errors
const (
	errNoProductID = errors.New("")
)

type Handlers struct {
	logic IUserAPILogic
	log   logger.ILogger
}

func NewUserAPIHandlers(u IUserAPILogic, l logger.ILogger) *Handlers {
	return &Handlers{u, l}
}

func (h *Handlers) ReviewsGet(w http.ResponseWriter, r *http.Request) {
	helpers.CallHandler(h.reviewsGet, w, r, h.log)
}

func (h *Handlers) reviewsGet(w http.ResponseWriter, r *http.Request) *helpers.AppError {
	productId := r.URL.Query().Get(productId)
	if productId == "" {
		return helpers.NewError()
	}
}
