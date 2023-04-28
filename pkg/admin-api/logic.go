package admin_api

import (
	"github.com/google/uuid"
	"github.com/yalagtyarzh/aggregator/pkg/models"
)

type IAdminAPILogic interface {
	CreateProduct(userID uuid.UUID, req models.ProductCreate) error
	UpdateProduct(userID uuid.UUID, req models.ProductUpdate) error
}

const CRUDProduct = "CRUDProduct"
