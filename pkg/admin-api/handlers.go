package admin_api

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/yalagtyarzh/aggregator/pkg/errors"
	"github.com/yalagtyarzh/aggregator/pkg/http/helpers"
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"github.com/yalagtyarzh/aggregator/pkg/models"
	"net/http"
)

type Handlers struct {
	logic     IAdminAPILogic
	log       logger.ILogger
	validator *validator.Validate
}

func NewAdminAPIHandlers(u IAdminAPILogic, l logger.ILogger) *Handlers {
	return &Handlers{u, l, validator.New()}
}

func (h *Handlers) ProductCreate(w http.ResponseWriter, r *http.Request) {
	helpers.CallHandler(h.productCreate, w, r, h.log)
}

func (h *Handlers) productCreate(w http.ResponseWriter, r *http.Request) *helpers.AppError {
	token, ok := r.Context().Value("token").(models.TokenPayload)
	if !ok {
		return helpers.NewError(http.StatusUnauthorized, errors.ErrInvalidUserID, "invalid user", false)
	}

	var req models.ProductCreate
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&req); err != nil {
		return helpers.NewError(http.StatusBadRequest, err, "invalid request body", false)
	}

	err := h.validator.Struct(req)
	if err != nil {
		return helpers.NewError(http.StatusBadRequest, err, "invalid request body", false)
	}

	err = h.logic.CreateProduct(token.UserID, req)
	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, err, "internal server error", false)
	}
	w.WriteHeader(http.StatusOK)

	return nil
}

func (h *Handlers) ProductUpdate(w http.ResponseWriter, r *http.Request) {
	helpers.CallHandler(h.productUpdate, w, r, h.log)
}

func (h *Handlers) productUpdate(w http.ResponseWriter, r *http.Request) *helpers.AppError {
	token, ok := r.Context().Value("token").(models.TokenPayload)
	if !ok {
		return helpers.NewError(http.StatusUnauthorized, errors.ErrInvalidUserID, "invalid user", false)
	}

	var req models.ProductUpdate
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&req); err != nil {
		return helpers.NewError(http.StatusBadRequest, err, "invalid request body", false)
	}

	err := h.validator.Struct(req)
	if err != nil {
		return helpers.NewError(http.StatusBadRequest, err, "invalid request body", false)
	}

	err = h.logic.UpdateProduct(token.UserID, req)
	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, err, "internal server error", false)
	}
	w.WriteHeader(http.StatusOK)

	return nil
}
