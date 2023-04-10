package admin_api

import (
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"github.com/yalagtyarzh/aggregator/pkg/repo"
)

type AdminAPILogic struct {
	log  logger.ILogger
	repo *repo.UserAPIRepository
}

func NewUserAPILogic(repositoryPool *repo.UserAPIRepository, log logger.ILogger) IAdminAPILogic {
	return &AdminAPILogic{
		log:  log,
		repo: repositoryPool,
	}
}
