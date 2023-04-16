package admin_api

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/yalagtyarzh/aggregator/pkg/http/helpers"
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"github.com/yalagtyarzh/aggregator/pkg/models"
	"net/http"
)

type Handlers struct {
	logic IAdminAPILogic
	log   logger.ILogger
}

var (
	errInvalidProductID = errors.New("invalid product id in request")
	errNoProductID      = errors.New("no product id in request")
	errInvalidUserID    = errors.New("invalid user id")
	errNoPermissions    = errors.New("no permission for to do request")
)

func NewAdminAPIHandlers(u IAdminAPILogic, l logger.ILogger) *Handlers {
	return &Handlers{u, l}
}

func (h *Handlers) ProductCreate(w http.ResponseWriter, r *http.Request) {
	helpers.CallHandler(h.productCreate, w, r, h.log)
}

func (h *Handlers) productCreate(w http.ResponseWriter, r *http.Request) *helpers.AppError {
	actorID, ok := r.Context().Value("userID").(uuid.UUID)
	if !ok {
		return helpers.NewError(http.StatusInternalServerError, errInvalidUserID, "internal server error", false)
	}

	var req models.ProductCreate
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&req); err != nil {
		return helpers.NewError(http.StatusBadRequest, err, "invalid request body", false)
	}

	err := h.logic.CreateProduct(actorID, req)
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
	actorID, ok := r.Context().Value("userID").(uuid.UUID)
	if !ok {
		return helpers.NewError(http.StatusInternalServerError, errInvalidUserID, "internal server error", false)
	}

	var req models.ProductUpdate
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&req); err != nil {
		return helpers.NewError(http.StatusBadRequest, err, "invalid request body", false)
	}

	err := h.logic.UpdateProduct(actorID, req)
	if err != nil {
		return helpers.NewError(http.StatusInternalServerError, err, "internal server error", false)
	}
	w.WriteHeader(http.StatusOK)

	return nil
}
