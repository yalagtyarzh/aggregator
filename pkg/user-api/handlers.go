package user_api

import (
	"github.com/yalagtyarzh/aggregator/pkg/logger"
)

type Handlers struct {
	logic IUserAPILogic
	log   logger.ILogger
}

func NewUserAPIHandlers(u IUserAPILogic, l logger.ILogger) *Handlers {
	return &Handlers{u, l}
}
