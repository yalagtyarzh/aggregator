package common

import (
	"context"
	"github.com/yalagtyarzh/aggregator/pkg/config"
	"github.com/yalagtyarzh/aggregator/pkg/http/middleware"
	"github.com/yalagtyarzh/aggregator/pkg/http/router"
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"github.com/yalagtyarzh/aggregator/pkg/prometheus"
	"github.com/yalagtyarzh/aggregator/pkg/provider"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Application struct {
	appProvider provider.IProvider
	services    []IService
	server      *http.Server
	appLogger   logger.ILogger
	metrics     *http.Server
	probes      *http.Server
}

func NewApp(basic config.BasicConfig, serverOptions config.ServerOptionsConfig, appProvider provider.IProvider,
	appLogger logger.ILogger, services ...IService) *Application {

	appPrometheus := prometheus.New(basic.Namespace, basic.PrometheusPrefix, basic.AppName)
	appMiddleware := middleware.New(appLogger, appPrometheus)
	metrics := router.Metrics()
	probes := router.Probe()
	handler := router.Router(appMiddleware, getApis(services)...)

	return &Application{
		appProvider: appProvider,
		server: &http.Server{
			Addr:         serverOptions.Bind,
			ReadTimeout:  serverOptions.ReadTimeout,
			IdleTimeout:  serverOptions.IdleTimeout,
			WriteTimeout: serverOptions.WriteTimeout,
			Handler:      handler,
		},
		metrics:   &http.Server{Addr: serverOptions.MetricsBind, Handler: metrics},
		probes:    &http.Server{Addr: serverOptions.ProbeBind, Handler: probes},
		services:  services,
		appLogger: appLogger,
	}
}

func getApis(services []IService) (apis []router.IAPI) {
	for i := range services {
		if v, ok := services[i].(router.IAPI); ok {
			apis = append(apis, v)
		}
	}

	return apis
}

func (app *Application) Start() {
	go func() {
		if err := app.metrics.ListenAndServe(); err != nil {
			app.appLogger.Warningf("listen metrics finished: %w", err)
		}
	}()
	app.appLogger.Infof("metric server started on %s", app.metrics.Addr)

	go func() {
		if err := app.probes.ListenAndServe(); err != nil {
			app.appLogger.Warningf("listen probe finished: %w", err)
		}
	}()
	app.appLogger.Infof("probe server started on %s", app.probes.Addr)

	listenErr := make(chan error, 1)
	go func() {
		listenErr <- app.server.ListenAndServe()
	}()
	app.appLogger.Infof("http server started on %s", app.server.Addr)

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	app.startServices()

	select {
	case err := <-listenErr:
		if err != nil {
			app.appLogger.Fatalf("%s", err.Error())
		}
	case s := <-osSignals:
		app.appLogger.Infof("SIGNAL:%s", s)
		app.server.SetKeepAlivesEnabled(false)
		app.stopServices()
		timeout := time.Second * 5
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer func() {
			cancel()
		}()
		if err := app.server.Shutdown(ctx); err != nil {
			app.appLogger.Fatalf("%s", err.Error())
		}
	}
	app.appProvider.Close()
	app.appLogger.Infof("App stopped")
}

func (app *Application) startServices() {
	app.appLogger.Infof("Starting pkg")
	for i := range app.services {
		if err := app.services[i].Start(); err != nil {
			app.appLogger.Fatalf("Couldn't start service %s: %s", app.services[i].Name(), err.Error())
		}
	}
}

func (app *Application) stopServices() {
	for i := range app.services {
		if err := app.services[i].Stop(); err != nil {
			app.appLogger.Errorf("error while stopping service %s: %s", app.services[i].Name(), err.Error())
		}
	}
	app.appLogger.Infof("Stopping pkg...")
}
