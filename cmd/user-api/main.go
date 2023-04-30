package main

import (
	"github.com/yalagtyarzh/aggregator/pkg/common"
	"github.com/yalagtyarzh/aggregator/pkg/config"
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"github.com/yalagtyarzh/aggregator/pkg/provider"
	"github.com/yalagtyarzh/aggregator/pkg/repo"
	api "github.com/yalagtyarzh/aggregator/pkg/user-api"
)

func main() {
	cfg := config.GetUserAPIConfig()
	log := logger.NewLogger(cfg.Basic.AppName, cfg.Logger)

	appProvider := provider.NewUserAPIProvider(cfg, log)
	appRepositories := repo.NewUserAPIRepoPool(appProvider, log)

	var appServices []common.IService
	mainAPIService := api.NewUserAPIService(appRepositories, log)
	appServices = append(appServices, mainAPIService)

	app := common.NewApp(cfg.Basic, cfg.ServerOptions, appProvider, log, appRepositories.JWT, appServices...)
	app.Start()
}
