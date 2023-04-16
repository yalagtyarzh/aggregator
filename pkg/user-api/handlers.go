package user_api

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/yalagtyarzh/aggregator/pkg/http/helpers"
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"github.com/yalagtyarzh/aggregator/pkg/models"
	"net/http"
	"strconv"
	"time"
)

// query params
const (
	productId = "pid"
	after     = "after"
	limit     = "limit"
	year      = "year"
	genre     = "genre"
)

// errors
var (
	errInvalidProductID = errors.New("invalid product id in request")
	errNoProductID      = errors.New("no product id in request")
	errInvalidUserID    = errors.New("invalid user id")
	errNoPermissions    = errors.New("no permission for to do request")
	errInvalidScore     = errors.New("invalid score")
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

func (h *Handlers) ReviewsCreate(w http.ResponseWriter, r *http.Request) {
	helpers.CallHandler(h.reviewsCreate, w, r, h.log)
}

func (h *Handlers) reviewsCreate(w http.ResponseWriter, r *http.Request) *helpers.AppError {
	actorID, ok := r.Context().Value("userID").(uuid.UUID)
	if !ok {
		return helpers.NewError(http.StatusInternalServerError, errInvalidUserID, "internal server error", false)
	}

	var req models.ReviewCreate
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&req); err != nil {
		return helpers.NewError(http.StatusBadRequest, err, "invalid request body", false)
	}

	if req.Score > 100 || req.Score < 0 {
		return helpers.NewError(http.StatusBadRequest, errInvalidScore, "invalid score", true)
	}

	err := h.logic.CreateReview(req, actorID)
	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, err, "internal server error", false)
	}

	w.WriteHeader(http.StatusOK)

	return nil
}

func (h *Handlers) ReviewsUpdate(w http.ResponseWriter, r *http.Request) {
	helpers.CallHandler(h.reviewsUpdate, w, r, h.log)
}

func (h *Handlers) reviewsUpdate(w http.ResponseWriter, r *http.Request) *helpers.AppError {
	actorID, ok := r.Context().Value("userID").(uuid.UUID)
	if !ok {
		return helpers.NewError(http.StatusInternalServerError, errInvalidUserID, "internal server error", false)
	}

	var req models.ReviewUpdate
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&req); err != nil {
		return helpers.NewError(http.StatusBadRequest, err, "invalid request body", false)
	}

	if req.Score > 100 || req.Score < 0 {
		return helpers.NewError(http.StatusBadRequest, errInvalidScore, "invalid score", true)
	}

	err := h.logic.UpdateReview(req, actorID)
	if err == errNoPermissions {
		return helpers.NewError(http.StatusForbidden, err, "no permissions", true)
	}

	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, err, "internal server error", false)
	}

	w.WriteHeader(http.StatusOK)

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
	after, err := strconv.Atoi(r.URL.Query().Get(after))
	if err != nil || after < 0 {
		after = 0
	}

	limit, err := strconv.Atoi(r.URL.Query().Get(limit))
	if err != nil || limit > 100 {
		limit = 100
	}

	year, _ := strconv.Atoi(year)

	now := time.Now().Year()
	if year > now {
		year = now
	}

	genre := r.URL.Query().Get(genre)

	resp, err := h.logic.GetProducts(after, limit, year, genre)
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

//func (h *Handlers) UserReviewsGet(w http.ResponseWriter, r *http.Request) {
//	helpers.CallHandler(h.userReviewsGet, w, r, h.log)
//}
//
//func (h *Handlers) userReviewsGet(w http.ResponseWriter, r *http.Request) *helpers.AppError {
//	return nil
//}
