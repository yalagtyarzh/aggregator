package config

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/google/uuid"
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"os"
	"sync"
	"time"
)

type IConfig interface {
	Validate() error
}

var (
	once sync.Once
)

type BasicConfig struct {
	InstanceID       uuid.UUID
	Environment      string `env:"ENVIRONMENT" envDefault:"local"`
	Namespace        string `env:"NAMESPACE" envDefault:"aggregator"`
	AppName          string `env:"APP_NAME" envDefault:"main-api"`
	PrometheusPrefix string `env:"PROMETHEUS_PREFIX" envDefault:"aggregator"`
}

type DBConfig struct {
	Scheme  string `env:"BACKEND_DB_SCHEME" envDefault:"postgres"`
	ConnStr string `env:"BACKEND_DB_CONN_STR,required"`
}

type ServerOptionsConfig struct {
	Bind         string        `env:"SERVER_BIND" envDefault:":9000"`
	ProbeBind    string        `env:"SERVER_PROBE_BIND" envDefault:":9091"`
	MetricsBind  string        `env:"SERVER_METRICS_BIND" envDefault:":9090"`
	ReadTimeout  time.Duration `env:"SERVER_READ_TIMEOUT" envDefault:"40s"`
	WriteTimeout time.Duration `env:"SERVER_WRITE_TIMEOUT" envDefault:"40s"`
	IdleTimeout  time.Duration `env:"SERVER_IDLE_TIMEOUT" envDefault:"40s"`
}

type MainAPIConfig struct {
	Logger        logger.LogConfig
	Basic         BasicConfig
	DB            DBConfig
	ServerOptions ServerOptionsConfig
}

func GetMainAPIConfig() *MainAPIConfig {
	var c MainAPIConfig
	once.Do(func() {
		if err := env.Parse(&c); err != nil {
			fmt.Printf("read configuration error: %s\n", err.Error())
			os.Exit(1)
		}

		if err := c.Validate(); err != nil {
			fmt.Printf("validating configuration error: %s\n", err.Error())
			os.Exit(1)
		}
	})

	return &c
}

func (c *MainAPIConfig) Validate() error {
	return nil
}
