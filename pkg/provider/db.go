package provider

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/yalagtyarzh/aggregator/pkg/config"
)

func NewAuthServicesDBConnection(c config.DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open(c.Scheme, c.ConnStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, err
	}

	/*
		db.SetMaxOpenConns(cfg.MaxOpenConnections)
		if cfg.MaxOpenConnections <= 0 {
			db.SetMaxOpenConns(10)
		}
		db.SetMaxIdleConns(cfg.MaxIdleConnections)
		db.SetConnMaxLifetime(cfg.MaxConnectionLifetime)
	*/

	return db, nil
}
