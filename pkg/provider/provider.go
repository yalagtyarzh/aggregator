package provider

import (
	"github.com/jmoiron/sqlx"
	"github.com/yalagtyarzh/aggregator/pkg/config"
	"github.com/yalagtyarzh/aggregator/pkg/logger"
)

type IProvider interface {
	Close()
}

type MainAPIProvider struct {
	log logger.ILogger
	db  *sqlx.DB
}

func NewMainAPIProvider(appConfig *config.MainAPIConfig, log logger.ILogger) *MainAPIProvider {
	db, err := NewAuthServicesDBConnection(appConfig.DB)
	if err != nil {
		log.Fatalf("error connecting to authorized services db: %s", err.Error())
	}
	log.Infof("connected to authorized services db")

	return &MainAPIProvider{
		log: log,
		db:  db,
	}
}

func (p *MainAPIProvider) Close() {
	if err := p.db.Close(); err != nil {
		p.log.Errorf("error while closing db: %s", err.Error())
	}

	p.log.Infof("Resources closed")
}

func (p *MainAPIProvider) AuthServicesDB() *sqlx.DB {
	return p.db
}
