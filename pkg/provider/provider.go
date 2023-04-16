package provider

import (
	"github.com/jmoiron/sqlx"
	"github.com/yalagtyarzh/aggregator/pkg/config"
	"github.com/yalagtyarzh/aggregator/pkg/logger"
)

type IProvider interface {
	Close()
}

type UserAPIProvider struct {
	log logger.ILogger
	db  *sqlx.DB
}

func NewUserAPIProvider(appConfig *config.UserAPIConfig, log logger.ILogger) *UserAPIProvider {
	db, err := NewDBConnection(appConfig.DB)
	if err != nil {
		log.Fatalf("error connecting to aggregator db: %s", err.Error())
	}
	log.Infof("connected to aggregator db")

	return &UserAPIProvider{
		log: log,
		db:  db,
	}
}

func (p *UserAPIProvider) Close() {
	if err := p.db.Close(); err != nil {
		p.log.Errorf("error while closing db: %s", err.Error())
	}

	p.log.Infof("Resources closed")
}

func (p *UserAPIProvider) DB() *sqlx.DB {
	return p.db
}
