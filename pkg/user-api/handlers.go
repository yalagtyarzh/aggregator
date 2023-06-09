package user_api

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/yalagtyarzh/aggregator/pkg/errors"
	"github.com/yalagtyarzh/aggregator/pkg/http/helpers"
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"github.com/yalagtyarzh/aggregator/pkg/models"
	"github.com/yalagtyarzh/aggregator/pkg/repo"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// query params
const (
	productId    = "pid"
	after        = "after"
	limit        = "limit"
	year         = "year"
	genre        = "genre"
	refreshToken = "refreshToken"
)

type Handlers struct {
	logic     IUserAPILogic
	log       logger.ILogger
	validator *validator.Validate
}

func NewUserAPIHandlers(u IUserAPILogic, l logger.ILogger) *Handlers {
	return &Handlers{u, l, validator.New()}
}

func (h *Handlers) ReviewsGet(w http.ResponseWriter, r *http.Request) {
	helpers.CallHandler(h.reviewsGet, w, r, h.log)
}

func (h *Handlers) reviewsGet(w http.ResponseWriter, r *http.Request) *helpers.AppError {
	productId := r.URL.Query().Get(productId)
	if productId == "" {
		return helpers.NewError(http.StatusBadRequest, errors.ErrNoProductID, "no product id in request", false)
	}

	pid, err := strconv.Atoi(productId)
	if err != nil || pid < 1 {
		return helpers.NewError(http.StatusBadRequest, errors.ErrInvalidProductID, "invalid product id in request", false)
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
	token, ok := r.Context().Value("token").(models.TokenPayload)
	if !ok {
		return helpers.NewError(http.StatusUnauthorized, errors.ErrInvalidUserID, "invalid user", false)
	}

	var req models.ReviewCreate
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&req); err != nil {
		return helpers.NewError(http.StatusBadRequest, err, "invalid request body", false)
	}

	err := h.validator.Struct(req)
	if err != nil {
		return helpers.NewError(http.StatusBadRequest, err, "invalid request body", false)
	}

	err = h.logic.CreateReview(req, token.UserID)
	if err == errors.ErrNoUser {
		return helpers.NewError(http.StatusBadRequest, err, "no user", true)
	}

	if err == errors.ErrTooManyReviews {
		return helpers.NewError(http.StatusBadRequest, err, "too many reviews for one user", true)
	}

	if err == repo.ErrForeignKeyViolation {
		return helpers.NewError(http.StatusBadRequest, err, "product not found", true)
	}

	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, err, "internal server error", false)
	}

	b, err := json.Marshal(models.StdResp{Message: "Ok"})
	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, err, "internal server error", false)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(b)

	return nil
}

func (h *Handlers) ReviewsUpdate(w http.ResponseWriter, r *http.Request) {
	helpers.CallHandler(h.reviewsUpdate, w, r, h.log)
}

func (h *Handlers) reviewsUpdate(w http.ResponseWriter, r *http.Request) *helpers.AppError {
	token, ok := r.Context().Value("token").(models.TokenPayload)
	if !ok {
		return helpers.NewError(http.StatusUnauthorized, errors.ErrInvalidUserID, "invalid user", false)
	}

	var req models.ReviewUpdate
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&req); err != nil {
		return helpers.NewError(http.StatusBadRequest, err, "invalid request body", false)
	}

	if !req.Delete {
		err := h.validator.Struct(req)
		if err != nil {
			return helpers.NewError(http.StatusBadRequest, err, "invalid request body", false)
		}
	}

	err := h.logic.UpdateReview(req, token.UserID)
	if err == errors.ErrNoPermissions {
		return helpers.NewError(http.StatusForbidden, err, "no permissions", true)
	}

	if err == errors.ErrNoUser {
		return helpers.NewError(http.StatusBadRequest, err, "no user", true)
	}

	if err == errors.ErrNoReview {
		return helpers.NewError(http.StatusBadRequest, err, "review not found", true)
	}

	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, err, "internal server error", false)
	}

	b, err := json.Marshal(models.StdResp{Message: "Ok"})
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
		return helpers.NewError(http.StatusBadRequest, errors.ErrNoProductID, "no product id in request", false)
	}

	pid, err := strconv.Atoi(productId)
	if err != nil || pid < 1 {
		return helpers.NewError(http.StatusBadRequest, errors.ErrInvalidProductID, "invalid product id in request", false)
	}

	resp, err := h.logic.GetProduct(pid)
	if err == errors.ErrNoProduct {
		return helpers.NewError(http.StatusNotFound, err, "product not found", true)
	}

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

	year, err := strconv.Atoi(r.URL.Query().Get(year))
	if err != nil {
		year = 0
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

	req.Email, req.Password, req.UserName, req.FirstName, req.LastName = strings.TrimSpace(req.Email),
		strings.TrimSpace(req.Password), strings.TrimSpace(req.UserName), strings.TrimSpace(req.FirstName), strings.TrimSpace(req.LastName)
	err := h.validator.Struct(req)
	if err != nil {
		return helpers.NewError(http.StatusBadRequest, err, "invalid request body", false)
	}

	resp, err := h.logic.CreateUser(req)
	if err == errors.ErrInvalidRole {
		return helpers.NewError(http.StatusBadRequest, err, "invalid role", true)
	}

	if err == errors.ErrUserAlreadyCreated {
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
		Name:     refreshToken,
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
	var req models.LoginRequest
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&req); err != nil {
		return helpers.NewError(http.StatusBadRequest, err, "invalid request body", false)
	}

	err := h.validator.Struct(req)
	if err != nil {
		return helpers.NewError(http.StatusBadRequest, err, "invalid request body", false)
	}

	resp, err := h.logic.Login(req.Username, req.Password)
	if err == errors.ErrInvalidPassword {
		return helpers.NewError(http.StatusBadRequest, err, "invalid password", true)
	}

	if err == errors.ErrNoUser {
		return helpers.NewError(http.StatusBadRequest, err, "user not found", true)
	}

	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, err, "internal server error", false)
	}

	b, err := json.Marshal(resp)

	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, err, "internal server error", false)
	}

	c := http.Cookie{
		Name:     refreshToken,
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
	helpers.CallHandler(h.logout, w, r, h.log)
}

func (h *Handlers) logout(w http.ResponseWriter, r *http.Request) *helpers.AppError {
	rToken, err := r.Cookie(refreshToken)
	if err != nil {
		return helpers.NewError(http.StatusBadRequest, err, "already logout", false)
	}

	err = h.logic.Logout(rToken.Value)
	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, err, "internal server error", false)
	}

	c := http.Cookie{
		Name:     refreshToken,
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
	}

	b, err := json.Marshal(models.StdResp{Message: "Ok"})
	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, err, "internal server error", false)
	}
	http.SetCookie(w, &c)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(b)

	return nil
}

func (h *Handlers) Refresh(w http.ResponseWriter, r *http.Request) {
	helpers.CallHandler(h.refresh, w, r, h.log)
}

func (h *Handlers) refresh(w http.ResponseWriter, r *http.Request) *helpers.AppError {
	rToken, err := r.Cookie(refreshToken)
	if err != nil || rToken.Value == "" {
		return helpers.NewError(http.StatusUnauthorized, err, "no token", false)
	}

	resp, err := h.logic.Refresh(rToken.Value)
	if err == repo.ErrInvalidTokenSignature {
		return helpers.NewError(http.StatusUnauthorized, err, "invalid token signature", true)
	}

	if err == repo.ErrValidate {
		return helpers.NewError(http.StatusBadRequest, err, "error validating token", true)
	}

	if err == repo.ErrInvalidToken {
		return helpers.NewError(http.StatusBadRequest, err, "invalid Token", true)
	}

	if err == errors.ErrNoToken {
		return helpers.NewError(http.StatusUnauthorized, err, "token not found", true)
	}

	if err == errors.ErrNoUser {
		return helpers.NewError(http.StatusUnauthorized, err, "user not found", true)
	}

	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, err, "internal server error", false)
	}

	b, err := json.Marshal(resp)

	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, err, "internal server error", false)
	}

	c := http.Cookie{
		Name:     refreshToken,
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

func (h *Handlers) GenresGet(w http.ResponseWriter, r *http.Request) {
	helpers.CallHandler(h.genresGet, w, r, h.log)
}

func (h *Handlers) genresGet(w http.ResponseWriter, r *http.Request) *helpers.AppError {
	g, err := h.logic.ListGenres()
	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, err, "internal server error", false)
	}

	b, err := json.Marshal(g)
	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, err, "internal server error", false)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(b)

	return nil
}

func (h *Handlers) GraphqlList(w http.ResponseWriter, r *http.Request) {
	helpers.CallHandler(h.graphqlList, w, r, h.log)
}

func (h *Handlers) graphqlList(w http.ResponseWriter, r *http.Request) *helpers.AppError {
	q, _ := io.ReadAll(r.Body)

	resp, err := h.logic.GraphqlList(string(q))
	if err == errors.ErrInvalidGraphql {
		return helpers.NewError(http.StatusBadRequest, err, "invalid request", true)
	}

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
