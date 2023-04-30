package admin_api

import (
	"github.com/google/uuid"
	"github.com/yalagtyarzh/aggregator/pkg/errors"
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"github.com/yalagtyarzh/aggregator/pkg/models"
	"github.com/yalagtyarzh/aggregator/pkg/repo"
)

type AdminAPILogic struct {
	log  logger.ILogger
	repo *repo.AdminAPIRepository
}

func NewAdminAPILogic(repositoryPool *repo.AdminAPIRepository, log logger.ILogger) IAdminAPILogic {
	return &AdminAPILogic{
		log:  log,
		repo: repositoryPool,
	}
}

func (l *AdminAPILogic) UpdateProduct(userID uuid.UUID, req models.ProductUpdate) error {
	u, err := l.repo.DB.GetUserByID(userID)
	if err != nil {
		return err
	}

	if u == nil {
		return errors.ErrNoUser
	}

	if u.Role != "Admin" {
		return errors.ErrNoPermissions
	}

	if req.Delete {
		return l.repo.DB.DeleteProduct(req.ID)
	}

	return l.repo.DB.UpdateProduct(req)
}

func (l *AdminAPILogic) CreateProduct(userID uuid.UUID, req models.ProductCreate) error {
	u, err := l.repo.DB.GetUserByID(userID)
	if err != nil {
		return err
	}

	if u == nil {
		return errors.ErrNoUser
	}

	if u.Role != "Admin" {
		return errors.ErrNoPermissions
	}

	return l.repo.DB.InsertProduct(req)
}
