package user_api

import (
	"encoding/json"
	"errors"
	"github.com/yalagtyarzh/aggregator/pkg/http/helpers"
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"net/http"
	"strconv"
)

// query params
const (
	productId = "pid"
	after     = "after"
	limit     = "limit"
)

// errors
var (
	errInvalidProductID = errors.New("invalid product id in request")
	errNoProductID      = errors.New("no product id in request")
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
		return helpers.NewError(http.StatusBadRequest, errNoProductID, "no product id in request", true)
	}

	pid, err := strconv.Atoi(productId)
	if err != nil || pid < 1 {
		return helpers.NewError(http.StatusBadRequest, errInvalidProductID, "invalid product id in request", true)
	}

	resp, err := h.logic.GetReviews(pid)
	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, err, "internal server error", false)
	}

	b, err := json.Marshal(resp)
	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, err, "internal server error", false)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(b)

	return nil
}

func (h *Handlers) ProductsGet(w http.ResponseWriter, r *http.Request) {
	helpers.CallHandler(h.productsGet, w, r, h.log)
}

func (h *Handlers) productsGet(w http.ResponseWriter, r *http.Request) *helpers.AppError {
	productId := r.URL.Query().Get(productId)
	if productId == "" {
		return helpers.NewError(http.StatusBadRequest, errNoProductID, "no product id in request", true)
	}

	pid, err := strconv.Atoi(productId)
	if err != nil || pid < 1 {
		return helpers.NewError(http.StatusBadRequest, errInvalidProductID, "invalid product id in request", true)
	}

	resp, err := h.logic.GetProduct(pid)
	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, err, "internal server error", false)
	}

	b, err := json.Marshal(resp)
	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, err, "internal server error", false)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(b)

	return nil
}

func (h *Handlers) ProductsScoreGet(w http.ResponseWriter, r *http.Request) {
	helpers.CallHandler(h.productsScoreGet, w, r, h.log)
}

func (h *Handlers) productsScoreGet(w http.ResponseWriter, r *http.Request) *helpers.AppError {
	productId := r.URL.Query().Get(productId)
	if productId == "" {
		return helpers.NewError(http.StatusBadRequest, errNoProductID, "no product id in request", true)
	}

	pid, err := strconv.Atoi(productId)
	if err != nil || pid < 1 {
		return helpers.NewError(http.StatusBadRequest, errInvalidProductID, "invalid product id in request", true)
	}

	resp, err := h.logic.GetProductScore(pid)
	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, err, "internal server error", false)
	}

	b, err := json.Marshal(resp)
	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, err, "internal server error", false)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(b)

	return nil
}

func (h *Handlers) ProductsGetMany(w http.ResponseWriter, r *http.Request) {
	helpers.CallHandler(h.productsGetMany, w, r, h.log)
}

func (h *Handlers) productsGetMany(w http.ResponseWriter, r *http.Request) *helpers.AppError {

	return nil
}

//func (h *Handlers) UserReviewsGet(w http.ResponseWriter, r *http.Request) {
//	helpers.CallHandler(h.userReviewsGet, w, r, h.log)
//}
//
//func (h *Handlers) userReviewsGet(w http.ResponseWriter, r *http.Request) *helpers.AppError {
//	return nil
//}
