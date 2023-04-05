package repo

import (
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"github.com/yalagtyarzh/aggregator/pkg/provider"
)

type MainAPIRepository struct {
	AuthServicesDB IDB
}

func NewMainAPIRepoPool(p *provider.MainAPIProvider, l logger.ILogger) *MainAPIRepository {
	return &MainAPIRepository{AuthServicesDB: NewAuthServicesDB(p.AuthServicesDB(), l)}
}
