package repo

import (
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"github.com/yalagtyarzh/aggregator/pkg/provider"
)

type UserAPIRepository struct {
	DB IDB
}

func NewUserAPIRepoPool(p *provider.UserAPIProvider, l logger.ILogger) *UserAPIRepository {
	return &UserAPIRepository{DB: NewDB(p.DB(), l)}
}
