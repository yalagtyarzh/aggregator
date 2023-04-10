package user_api

import (
	"github.com/gorilla/mux"
	"github.com/yalagtyarzh/aggregator/pkg/http/middleware"
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"github.com/yalagtyarzh/aggregator/pkg/repo"
)

type UserAPIService struct {
	log      logger.ILogger
	logic    IUserAPILogic
	handlers *Handlers
}

func NewUserAPIService(repositoryPool *repo.UserAPIRepository, log logger.ILogger) *UserAPIService {
	logic := NewUserAPILogic(repositoryPool, log)
	return &UserAPIService{
		log:      log,
		logic:    logic,
		handlers: NewUserAPIHandlers(logic, log),
	}
}

func (s *UserAPIService) Router(r *mux.Router, mw *middleware.Middleware) {
	v1 := r.PathPrefix("/api/v1").Subrouter()

	v1.Use(mw.AuthMiddleware)
}

func (s *UserAPIService) Start() error {
	s.log.Infof(s.Name() + " started")
	return nil
}

func (s *UserAPIService) Stop() error {
	s.log.Infof(s.Name() + " stopped")
	return nil
}

func (s *UserAPIService) Name() string {
	return "Aggregator main API service"
}
