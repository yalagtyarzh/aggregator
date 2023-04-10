package user_api

import (
	"github.com/yalagtyarzh/aggregator/pkg/logger"
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
