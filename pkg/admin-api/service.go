package admin_api

import (
	"github.com/gorilla/mux"
	"github.com/yalagtyarzh/aggregator/pkg/http/middleware"
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"github.com/yalagtyarzh/aggregator/pkg/repo"
)

type AdminAPIService struct {
	log      logger.ILogger
	logic    IAdminAPILogic
	handlers *Handlers
}

func NewAdminAPIService(repositoryPool *repo.AdminAPIRepository, log logger.ILogger) *AdminAPIService {
	logic := NewAdminAPILogic(repositoryPool, log)
	return &AdminAPIService{
		log:      log,
		logic:    logic,
		handlers: NewAdminAPIHandlers(logic, log),
	}
}

func (s *AdminAPIService) Router(r *mux.Router, mw *middleware.Middleware) {
	v1 := r.PathPrefix("/api/v1/admin").Subrouter()
	v1.HandleFunc("/product/create", s.handlers.ProductCreate).Methods("POST", "OPTIONS")
	v1.HandleFunc("/product/update", s.handlers.ProductUpdate).Methods("POST", "OPTIONS")

	v1.Use(mw.AuthMiddleware)
}

func (s *AdminAPIService) Start() error {
	s.log.Infof(s.Name() + " started")
	return nil
}

func (s *AdminAPIService) Stop() error {
	s.log.Infof(s.Name() + " stopped")
	return nil
}

func (s *AdminAPIService) Name() string {
	return "Aggregator admin API service"
}
