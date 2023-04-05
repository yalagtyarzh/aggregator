package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/yalagtyarzh/aggregator/pkg/logger"
)

type authServicesDBPSQL struct {
	logger.ILogger
	db *sqlx.DB
}

func NewAuthServicesDB(db *sqlx.DB, log logger.ILogger) IDB {
	return &authServicesDBPSQL{
		log,
		db,
	}
}
