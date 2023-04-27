package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"github.com/yalagtyarzh/aggregator/pkg/models"
	"strings"
)

type DBPSQL struct {
	logger.ILogger
	db *sqlx.DB
}

const foreignKeyViolation = "23503"

var (
	ErrForeignKeyViolation = errors.New("foreign key violation error")
)

func NewDB(db *sqlx.DB, log logger.ILogger) IDB {
	return &DBPSQL{
		log,
		db,
	}
}

func (d *DBPSQL) GetReviewsByProductID(pid int) ([]models.Review, error) {
	r := make([]models.Review, 0)

	stmt := `select r.id, r.score, r.content, r.content_html, r.created_at, r.updated_at, u.id, u.first_name, u.last_name, u.user_name, u.email, u.password, u.role, u.created_at, u.updated_at
			 from reviews r
			 join products_reviews pr on r.id = pr.review_id
			 join users u on r.user_id = u.id
			 where pr.product_id = $1
	`

	if err := d.db.Select(&r, stmt, pid); err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return r, nil
}

func (d *DBPSQL) GetReviewByID(reviewID int) (*models.Review, error) {
	var r models.Review

	stmt := `select r.id, r.score, r.content, r.content_html, r.created_at, r.updated_at, u.id, u.first_name, u.last_name, u.user_name, u.email, u.password, u.role, u.created_at, u.updated_at
			 from reviews r
			 join users u on r.user_id = u.id
			 where r.id = $1 LIMIT 1
	`

	if err := d.db.Select(&r, stmt, reviewID); err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &r, nil
}

func (d *DBPSQL) GetProductById(productId int) (*models.Product, error) {
	var p models.Product

	stmt := `select p.id, p.title, p.description, p.year, p.release_date, p.studio, p.rating, round(avg(r.score)), p.created_at, p.updated_at from products p
             join products_reviews pr on p.id = pr.product_id
             join reviews on pr.review_id = r.id                                                                                                            
             where p.id = $1 limit 1
             group by p.id`

	if err := d.db.Get(&p, stmt, productId); err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err := d.db.Select(&p.Genres, `select genre from products_genres where product_id=$1`, productId); err != nil {
		return nil, err
	}

	return &p, nil
}

func (d *DBPSQL) GetProducts(after int, limit int, year int, genre string) ([]models.Product, error) {
	p := make([]models.Product, 0)

	stmt := `select p.id, p.title, p.description, p.year, p.release_date, p.studio, p.rating, p.created_at, p.updated_at 
			 from products p
			 join products_genres pg on p.id=pg.product_id
			 join products_reviews pr on p.id = pr.product_id
			 join reviews on pr.review_id = r.id`

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

	stmt += " GROUP BY p.id LIMIT $1 OFFSET $2"
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

func (d *DBPSQL) GetPermissionsByRole(userID uuid.UUID) ([]models.Permission, error) {
	p := make([]models.Permission, 0)

	stmt := `select permission from roles_permissions where role=(select role from users where id = $1 LIMIT 1)`
	err := d.db.Select(&p, stmt, userID)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (d *DBPSQL) UpdateReview(rc models.ReviewUpdate) error {
	_, err := d.db.Exec("update reviews set score=$1, content=$2, content_html=$3 where id = $4", rc.Score, rc.Content, rc.ContentHTML, rc.ID)
	return err
}

func (d *DBPSQL) DeleteReview(reviewID int) error {
	_, err := d.db.Exec("delete from reviews where id = $1", reviewID)
	return err
}

func (d *DBPSQL) InsertReview(rc models.ReviewCreate, userID uuid.UUID) error {
	res, err := d.db.Exec("INSERT INTO reviews(score, content, content_html, user_id) VALUES($1, $2, $3, $4)", rc.Score, rc.Content, rc.ContentHTML, userID)
	if err != nil {
		return err
	}

	reviewID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	_, err = d.db.Exec("INSERT INTO products_reviews(product_id, review_id) VALUES($1, $2)", rc.ProductID, reviewID)
	return err
}

func (d *DBPSQL) DeleteProduct(productID int) error {
	_, err := d.db.Exec("delete from products where id = $1", productID)
	return err
}

func (d *DBPSQL) UpdateProduct(p models.ProductUpdate) error {
	_, err := d.db.Exec("update products set title = $1, description = $2, year = $3, release_date = $4, studio = $5, rating = $6 where id = $7", p.Title, p.Description, p.Year, p.ReleaseDate, p.Studio, p.Rating, p.ID)
	if err != nil {
		return err
	}

	_, err = d.db.Exec("delete from products_genres where product_id = $1", p.ID)

	if p.Genres == nil || len(p.Genres) == 0 {
		return nil
	}

	for _, v := range p.Genres {
		_, err = d.db.Exec("insert into products_genres(product_id, genre) VALUES($1, $2)", p.ID, v.Genre)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *DBPSQL) InsertProduct(p models.ProductCreate) error {
	res, err := d.db.Exec("insert into products(title, description, year, release_date, studio, rating) values($1, $2, $3, $4, $5, $6)", p.Title, p.Description, p.Year, p.ReleaseDate, p.Studio, p.Rating)
	if err != nil {
		return err
	}

	productID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	for _, v := range p.Genres {
		_, err = d.db.Exec("insert into products_genres(product_id, genre) VALUES($1, $2)", productID, v.Genre)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *DBPSQL) GetUserByEmail(email string) (*models.User, error) {
	var u models.User

	if err := d.db.Get(&u, "select id, first_name, last_name, user_name, email, password, role, created_at, updated_at, is_deleted from users where email=$1", email); err != nil {
		if err != sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}

func (d *DBPSQL) InsertUser(u models.CreateUser) error {
	_, err := d.db.Exec("insert into users(first_name, last_name, user_name, email, password, role) values($1, $2, $3, $4, $5, $6)", u.FirstName, u.LastName, u.UserName, u.Email, u.Password, u.Role)
	if driverErr, ok := err.(*pq.Error); ok {
		if driverErr.Code == foreignKeyViolation {
			return ErrForeignKeyViolation
		}
	}

	return err
}
