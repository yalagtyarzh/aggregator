package repo

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"github.com/yalagtyarzh/aggregator/pkg/models"
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

func (d *authServicesDBPSQL) GetReviewsByProductID(pid int) ([]models.Review, error) {
	r := make([]models.Review, 0)

	stmt := `select r.id, r.score, r.content, r.content_html, r.created_at, r.updated_at 
			 from reviews r
			 join products_reviews pr on r.id = pr.review_id
			 where pr.product_id = $1
	`

	if err := d.db.Select(&r, stmt, pid); err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return r, nil
}

func (d *authServicesDBPSQL) GetProductById(productId int) (*models.Product, error) {
	var p models.Product

	stmt := `select id, title, description, year, release_date, studio, rating, created_at, updated_at from products where id = $1 limit 1`

	if err := d.db.Get(&p, stmt, productId); err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &p, nil
}

func (d *authServicesDBPSQL) GetProductScoreById(productId int) (*models.Score, error) {
	var s models.Score

	stmt := `select round(AVG(r.score)) as score from reviews r join products_review pr on r.id = pr.review_id where pr.product_id = $1`

	if err := d.db.Get(&s, stmt, productId); err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &s, nil
}
