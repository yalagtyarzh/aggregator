package user_api

import (
	"github.com/google/uuid"
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"github.com/yalagtyarzh/aggregator/pkg/models"
	"github.com/yalagtyarzh/aggregator/pkg/repo"
)

type UserAPILogic struct {
	log  logger.ILogger
	repo *repo.UserAPIRepository
}

func NewUserAPILogic(repositoryPool *repo.UserAPIRepository, log logger.ILogger) IUserAPILogic {
	return &UserAPILogic{
		log:  log,
		repo: repositoryPool,
	}
}

func (l *UserAPILogic) GetReviews(productId int) ([]models.Review, error) {
	r, err := l.repo.DB.GetReviewsByProductID(productId)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (l *UserAPILogic) GetProduct(productId int) (models.Product, error) {
	p, err := l.repo.DB.GetProductById(productId)
	if err != nil {
		return models.Product{}, nil
	}

	return *p, nil
}

func (l *UserAPILogic) GetProductScore(productId int) (models.Score, error) {
	s, err := l.repo.DB.GetProductScoreById(productId)
	if err != nil {
		return models.Score{}, nil
	}

	return *s, nil
}

func (l *UserAPILogic) GetProducts(after, limit, year int, genre string) ([]models.Product, error) {
	s, err := l.repo.DB.GetProducts(after, limit, year, genre)
	if err != nil {
		return []models.Product{}, err
	}

	return s, nil
}

func (l *UserAPILogic) CreateReview(rc models.ReviewCreate, userID uuid.UUID) error {
	return l.repo.DB.InsertReview(rc, userID)
}

func (l *UserAPILogic) UpdateReview(rc models.ReviewUpdate, userID uuid.UUID) error {
	r, err := l.repo.DB.GetReviewByID(rc.ID)
	if err != nil {
		return err
	}

	if r.User.ID.String() != userID.String() {
		perms, err := l.repo.DB.GetPermissionsByRole(userID)
		if err != nil {
			return err
		}

		var flag bool
		for _, v := range perms {
			if v.Permission == "editForeignReviews" {
				flag = true
				break
			}
		}

		if !flag {
			return errNoPermissions
		}
	}

	if rc.Delete {
		return l.repo.DB.DeleteReview(rc.ID)
	}

	return l.repo.DB.UpdateReview(rc)
}
