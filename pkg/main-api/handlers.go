package main_api

import (
	"github.com/yalagtyarzh/aggregator/pkg/logger"
)

type Handlers struct {
	logic IMainAPILogic
	log   logger.ILogger
}

func NewMainAPIHandlers(u IMainAPILogic, l logger.ILogger) *Handlers {
	return &Handlers{u, l}
}
