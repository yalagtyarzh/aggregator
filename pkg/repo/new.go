package repo

import (
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"github.com/yalagtyarzh/aggregator/pkg/provider"
)

type UserAPIRepository struct {
	DB  IDB
	JWT *JWTer
}

func NewUserAPIRepoPool(p *provider.UserAPIProvider, l logger.ILogger) *UserAPIRepository {
	db := NewDB(p.DB(), l)
	return &UserAPIRepository{DB: db, JWT: NewJWTer(db, p.JWT())}
}

type AdminAPIRepository struct {
	DB  IDB
	JWT *JWTer
}

func NewAdminAPIRepoPool(p *provider.AdminAPIProvider, l logger.ILogger) *AdminAPIRepository {
	db := NewDB(p.DB(), l)
	return &AdminAPIRepository{DB: db, JWT: NewJWTer(db, p.JWT())}
}
