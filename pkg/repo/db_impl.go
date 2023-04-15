package repo

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"github.com/yalagtyarzh/aggregator/pkg/models"
	"strings"
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

	if err := d.db.Select(&p.Genres, `select genre from products_genres where product_id=$1`, productId); err != nil {
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

func (d *authServicesDBPSQL) GetProducts(after int, limit int, year int, genre string) ([]models.Product, error) {
	p := make([]models.Product, 0)

	stmt := `select p.id, p.title, p.description, p.year, p.release_date, p.studio, p.rating, p.created_at, p.updated_at 
			 from products p
			 join products_genres pg on p.id=pg.product_id`

	where := make([]string, 0)

	if year > 0 {
		where = append(where, fmt.Sprintf("p.year = %d", year))
	}

	if genre != "" {
		where = append(where, fmt.Sprintf("pg.genre = '%s'", genre))
	}

	if len(where) < 0 {
		stmt = stmt + " WHERE " + strings.Join(where, " AND ")
	}

	stmt += " LIMIT $1 OFFSET $2"
	err := d.db.Select(&p, stmt, limit, after)
	if err != nil {
		return nil, err
	}

	stmt = `select genre from products_genres where product_id=$1`
	for _, v := range p {
		err = d.db.Select(&v.Genres, stmt, v.ID)
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}
