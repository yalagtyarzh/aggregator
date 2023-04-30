package middleware

import (
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"github.com/yalagtyarzh/aggregator/pkg/prometheus"
	"github.com/yalagtyarzh/aggregator/pkg/repo"
	"time"
)

type Middleware struct {
	appStartTime  time.Time
	appLogger     logger.ILogger
	appPrometheus *prometheus.Prometheus
	appJWT        *repo.JWTer
}

func New(appLogger logger.ILogger, appPrometheus *prometheus.Prometheus, appJWT *repo.JWTer) *Middleware {
	return &Middleware{time.Now(), appLogger, appPrometheus, appJWT}
}

func (m *Middleware) GetStartTime() time.Time {
	return m.appStartTime
}
