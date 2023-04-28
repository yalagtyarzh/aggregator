package provider

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/yalagtyarzh/aggregator/pkg/config"
)

func NewDBConnection(c config.DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open(c.Scheme, c.ConnStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, err
	}

	return db, nil
}
