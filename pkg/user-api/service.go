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
	reviews := v1.PathPrefix("/reviews").Subrouter()
	reviews.HandleFunc("/create", s.handlers.ReviewsCreate).Methods("POST", "OPTIONS")
	reviews.HandleFunc("/update", s.handlers.ReviewsUpdate).Methods("POST", "OPTIONS")
	reviews.Use(mw.AuthMiddleware)

	v1.HandleFunc("/reviews/get", s.handlers.ReviewsGet).Methods("GET", "OPTIONS")
	v1.HandleFunc("/products", s.handlers.ProductsGetMany).Methods("GET", "OPTIONS")
	v1.HandleFunc("/products/get", s.handlers.ProductsGet).Methods("GET", "OPTIONS")
	v1.HandleFunc("/genres", s.handlers.GenresGet).Methods("GET", "OPTIONS")

	v1.HandleFunc("/graphql/list", s.handlers.GraphqlList).Methods("POST", "OPTIONS")

	v1.HandleFunc("/registration", s.handlers.Registration).Methods("POST", "OPTIONS")
	v1.HandleFunc("/login", s.handlers.Login).Methods("POST", "OPTIONS")
	v1.HandleFunc("/logout", s.handlers.Logout).Methods("POST", "OPTIONS")
	v1.HandleFunc("/refresh", s.handlers.Refresh).Methods("GET", "OPTIONS")
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
