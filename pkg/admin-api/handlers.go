package admin_api

import (
	"github.com/yalagtyarzh/aggregator/pkg/logger"
)

type Handlers struct {
	logic IAdminAPILogic
	log   logger.ILogger
}

func NewAdminAPIHandlers(u IAdminAPILogic, l logger.ILogger) *Handlers {
	return &Handlers{u, l}
}
