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

	p, err := l.repo.DB.GetProduct(req.ID)
	if err != nil {
		return err
	}

	if p == nil {
		return errors.ErrNoProduct
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

func (l *AdminAPILogic) PromoteRole(token models.TokenPayload, role string, userId uuid.UUID) error {
	actor, err := l.repo.DB.GetUserByID(token.UserID)
	if err != nil {
		return err
	}

	if actor == nil || actor.Role != "Admin" {
		return errors.ErrNoPermissions
	}

	user, err := l.repo.DB.GetUserByID(userId)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.ErrNoUser
	}

	err = l.repo.DB.UpdateUserRole(userId, role)
	if err == repo.ErrForeignKeyViolation {
		return errors.ErrInvalidRole
	}

	return err
}

func (l *AdminAPILogic) GetUsers(token models.TokenPayload) ([]models.User, error) {
	actor, err := l.repo.DB.GetUserByID(token.UserID)
	if err != nil {
		return nil, err
	}

	if actor == nil || actor.Role != "Admin" {
		return nil, err
	}

	resp, err := l.repo.DB.GetUsers()
	if err != nil {
		return nil, err
	}

	if resp == nil {
		resp = make([]models.User, 0)
	}

	return resp, nil
}
