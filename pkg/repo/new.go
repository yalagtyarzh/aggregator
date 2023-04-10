package repo

import (
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"github.com/yalagtyarzh/aggregator/pkg/provider"
)

type UserAPIRepository struct {
	AuthServicesDB IDB
}

func NewUserAPIRepoPool(p *provider.UserAPIProvider, l logger.ILogger) *UserAPIRepository {
	return &UserAPIRepository{AuthServicesDB: NewAuthServicesDB(p.AuthServicesDB(), l)}
}
