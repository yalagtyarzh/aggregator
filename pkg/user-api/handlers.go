package user_api

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/yalagtyarzh/aggregator/pkg/http/helpers"
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"github.com/yalagtyarzh/aggregator/pkg/models"
	"net/http"
	"net/mail"
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
	errInvalidProductID   = errors.New("invalid product id in request")
	errNoProductID        = errors.New("no product id in request")
	errInvalidUserID      = errors.New("invalid user id")
	errNoPermissions      = errors.New("no permission for to do request")
	errUserAlreadyCreated = errors.New("user is already created")
	errInvalidScore       = errors.New("invalid score")
	errInvalidRole        = errors.New("invalid role")
	errInvalidPassword    = errors.New("invalid password")
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

func (h *Handlers) Registration(w http.ResponseWriter, r *http.Request) {
	helpers.CallHandler(h.registration, w, r, h.log)
}

func (h *Handlers) registration(w http.ResponseWriter, r *http.Request) *helpers.AppError {
	var req models.CreateUser
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&req); err != nil {
		return helpers.NewError(http.StatusBadRequest, err, "invalid request body", false)
	}

	if len(req.FirstName) == 0 {
		return helpers.NewError(http.StatusBadRequest, nil, "no firstname", false)
	}

	if len(req.LastName) == 0 {
		return helpers.NewError(http.StatusBadRequest, nil, "no lastname", false)
	}

	if len(req.UserName) == 0 {
		return helpers.NewError(http.StatusBadRequest, nil, "no username", false)
	}

	if len(req.Password) <= 3 {
		return helpers.NewError(http.StatusBadRequest, nil, "too short password", false)
	}

	if len(req.Password) > 32 {
		return helpers.NewError(http.StatusBadRequest, nil, "too long password", false)
	}

	e, err := mail.ParseAddress(req.Email)
	if err != nil {
		return helpers.NewError(http.StatusBadRequest, err, "invalid mail", false)
	}

	req.Email = e.Address

	resp, err := h.logic.CreateUser(req)
	if err == errInvalidRole {
		return helpers.NewError(http.StatusBadRequest, err, "invalid role", false)
	}

	if err == errUserAlreadyCreated {
		return helpers.NewError(http.StatusBadRequest, err, "user already created", true)
	}

	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, err, "internal server error", false)
	}

	b, err := json.Marshal(resp)
	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, err, "internal server error", false)
	}

	c := http.Cookie{
		Name:     "refreshToken",
		Value:    resp.RefreshToken,
		MaxAge:   int((30 * 24 * time.Hour).Milliseconds()),
		HttpOnly: true,
	}

	http.SetCookie(w, &c)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(b)

	return nil
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	helpers.CallHandler(h.login, w, r, h.log)
}

func (h *Handlers) login(w http.ResponseWriter, r *http.Request) *helpers.AppError {
	username := r.FormValue("username")
	password := r.FormValue("password")
	if len(username) == 0 {
		return helpers.NewError(http.StatusBadRequest, nil, "no username", false)
	}

	if len(password) <= 3 {
		return helpers.NewError(http.StatusBadRequest, nil, "too short password", false)
	}

	if len(password) > 32 {
		return helpers.NewError(http.StatusBadRequest, nil, "too long password", false)
	}

	resp, err := h.logic.Login(username, password)
	if err == errInvalidPassword {
		return helpers.NewError(http.StatusBadRequest, nil, "invalid password", false)
	}

	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, err, "internal server error", false)
	}

	b, err := json.Marshal(resp)

	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, err, "internal server error", false)
	}

	c := http.Cookie{
		Name:     "refreshToken",
		Value:    resp.RefreshToken,
		MaxAge:   int((30 * 24 * time.Hour).Milliseconds()),
		HttpOnly: true,
	}

	http.SetCookie(w, &c)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(b)

	return nil
}

func (h *Handlers) Logout(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) Refresh(w http.ResponseWriter, r *http.Request) {

}

//func (h *Handlers) UserReviewsGet(w http.ResponseWriter, r *http.Request) {
//	helpers.CallHandler(h.userReviewsGet, w, r, h.log)
//}
//
//func (h *Handlers) userReviewsGet(w http.ResponseWriter, r *http.Request) *helpers.AppError {
//	return nil
//}
