package main_api

import (
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"github.com/yalagtyarzh/aggregator/pkg/repo"
)

type MainAPILogic struct {
	log  logger.ILogger
	repo *repo.MainAPIRepository
}

func NewMainAPILogic(repositoryPool *repo.MainAPIRepository, log logger.ILogger) IMainAPILogic {
	return &MainAPILogic{
		log:  log,
		repo: repositoryPool,
	}
}
