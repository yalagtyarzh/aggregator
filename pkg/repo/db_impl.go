package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"github.com/yalagtyarzh/aggregator/pkg/models"
)

type authServicesDBPSQL struct {
	logger.ILogger
	db *sqlx.DB
}

func (a authServicesDBPSQL) GetReviewsByProductID(pid int) ([]models.Review, error) {
	r := make([]models.Review, 0)

	return r, nil
}

func NewAuthServicesDB(db *sqlx.DB, log logger.ILogger) IDB {
	return &authServicesDBPSQL{
		log,
		db,
	}
}
