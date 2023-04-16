package admin_api

import (
	"github.com/google/uuid"
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"github.com/yalagtyarzh/aggregator/pkg/models"
	"github.com/yalagtyarzh/aggregator/pkg/repo"
)

type AdminAPILogic struct {
	log  logger.ILogger
	repo *repo.UserAPIRepository
}

const CRUDProduct = "CRUDProduct"

func (l *AdminAPILogic) UpdateProduct(userID uuid.UUID, req models.ProductUpdate) error {
	perms, err := l.repo.DB.GetPermissionsByRole(userID)
	if err != nil {
		return err
	}

	var flag bool
	for _, v := range perms {
		if v.Permission == CRUDProduct {
			flag = true
			break
		}
	}

	if !flag {
		return errNoPermissions
	}

	if req.Delete {
		return l.repo.DB.DeleteProduct(req.ID)
	}

	return l.repo.DB.UpdateProduct(req)
}

func NewUserAPILogic(repositoryPool *repo.UserAPIRepository, log logger.ILogger) IAdminAPILogic {
	return &AdminAPILogic{
		log:  log,
		repo: repositoryPool,
	}
}

func (l *AdminAPILogic) CreateProduct(userID uuid.UUID, req models.ProductCreate) error {
	perms, err := l.repo.DB.GetPermissionsByRole(userID)
	if err != nil {
		return err
	}

	var flag bool
	for _, v := range perms {
		if v.Permission == CRUDProduct {
			flag = true
			break
		}
	}

	if !flag {
		return errNoPermissions
	}

	return l.repo.DB.InsertProduct(req)
}
