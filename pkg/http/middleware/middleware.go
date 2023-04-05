package middleware

import (
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"github.com/yalagtyarzh/aggregator/pkg/prometheus"
	"time"
)

type Middleware struct {
	appStartTime  time.Time
	appLogger     logger.ILogger
	appPrometheus *prometheus.Prometheus
}

func New(appLogger logger.ILogger, appPrometheus *prometheus.Prometheus) *Middleware {
	return &Middleware{time.Now(), appLogger, appPrometheus}
}

func (m *Middleware) GetStartTime() time.Time {
	return m.appStartTime
}
