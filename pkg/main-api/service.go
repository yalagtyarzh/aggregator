package main_api

import (
	"github.com/gorilla/mux"
	"github.com/yalagtyarzh/aggregator/pkg/http/middleware"
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"github.com/yalagtyarzh/aggregator/pkg/repo"
)

type MainAPIService struct {
	log      logger.ILogger
	logic    IMainAPILogic
	handlers *Handlers
}

func NewMainAPIService(repositoryPool *repo.MainAPIRepository, log logger.ILogger) *MainAPIService {
	logic := NewMainAPILogic(repositoryPool, log)
	return &MainAPIService{
		log:      log,
		logic:    logic,
		handlers: NewMainAPIHandlers(logic, log),
	}
}

func (s *MainAPIService) Router(r *mux.Router, mw *middleware.Middleware) {
	v1 := r.PathPrefix("/api/v1").Subrouter()

	v1.Use(mw.AuthMiddleware)
}

func (s *MainAPIService) Start() error {
	s.log.Infof(s.Name() + " started")
	return nil
}

func (s *MainAPIService) Stop() error {
	s.log.Infof(s.Name() + " stopped")
	return nil
}

func (s *MainAPIService) Name() string {
	return "Aggregator main API service"
}
