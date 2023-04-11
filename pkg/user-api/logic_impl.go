package user_api

import (
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

func (l *UserAPILogic) GetReviews(pid int) ([]models.Review, error) {
	r, err := l.repo.DB.GetReviewsByProductID(pid)
	if err != nil {
		return nil, err
	}

	return r, nil
}
