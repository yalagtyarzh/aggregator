package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
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

type JWTConfig struct {
	AccessSecret  string `env:"ACCESS_SECRET" envDefault:"access-secret"`
	RefreshSecret string `env:"REFRESH_SECRET" envDefault:"refresh-secret"`
	SigningKey    string `env:"SIGNING_KEY" envDefault:"WcPGXClpKD7Bc1C0CCDA1060E2GGlTfamrd8-W0ghBE"`
}

type UserAPIConfig struct {
	Logger        logger.LogConfig
	Basic         BasicConfig
	DB            DBConfig
	ServerOptions ServerOptionsConfig
	JWT           JWTConfig `envPrefix:"JWT_"`
}

type AdminAPIConfig struct {
	Logger        logger.LogConfig
	Basic         BasicConfig
	DB            DBConfig
	ServerOptions ServerOptionsConfig
}

func GetUserAPIConfig() *UserAPIConfig {
	var c UserAPIConfig
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

func (c *UserAPIConfig) Validate() error {
	return nil
}

func GetAdminAPIConfig() *AdminAPIConfig {
	var c AdminAPIConfig
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

func (c *AdminAPIConfig) Validate() error {
	return nil
}
