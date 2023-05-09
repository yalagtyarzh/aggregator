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

type dbPSQL struct {
	logger.ILogger
	db *sqlx.DB
}

const (
	foreignKeyViolation = "23503"
	uniqueKeyViolation  = "23505"
)

var (
	ErrForeignKeyViolation = errors.New("foreign key violation error")
	ErrUniqueViolation     = errors.New("unique violation error")
)

func NewDB(db *sqlx.DB, log logger.ILogger) IDB {
	return &dbPSQL{
		log,
		db,
	}
}

func (d *dbPSQL) GetReviewsByProductID(pid int) ([]models.Review, error) {
	r := make([]models.Review, 0)

	stmt := `select r.id as id, r.score, r.content, r.content_html, r.created_at, r.updated_at, u.id as user_id, u.first_name, u.last_name, u.user_name
			 from reviews r
			 join users u on r.user_id = u.id
			 where r.product_id = $1
	`

	if err := d.db.Select(&r, stmt, pid); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return r, nil
}
func (d *dbPSQL) GetReviewByUserAndProduct(productId int, userId uuid.UUID) (*models.Review, error) {
	var r models.Review

	stmt := `select r.id as id, r.score, r.content, r.content_html, r.created_at, r.updated_at, u.id as user_id, u.first_name, u.last_name, u.user_name
			 from reviews r
			 join users u on r.user_id = u.id
			 where u.id = $1 AND r.product_id = $2 and r.is_deleted=false LIMIT 1
	`

	if err := d.db.Get(&r, stmt, userId, productId); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &r, nil
}

func (d *dbPSQL) GetReviewByID(reviewID int) (*models.Review, error) {
	var r models.Review

	stmt := `select r.id as id, r.score, r.content, r.content_html, r.created_at, r.updated_at, u.id as user_id, u.first_name, u.last_name, u.user_name
			 from reviews r
			 join users u on r.user_id = u.id
			 where r.id = $1 LIMIT 1
	`

	if err := d.db.Get(&r, stmt, reviewID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &r, nil
}

func (d *dbPSQL) GetProduct(productId int) (*models.Product, error) {
	var p models.Product

	stmt := `select id, title, description, year, studio, rating,  created_at, updated_at, 
			 coalesce((SELECT round(avg(score)) from reviews where product_id = $1), 0) as score from products
			 where id = $1
			 group by id
			 limit 1`

	if err := d.db.Get(&p, stmt, productId); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	var g []models.Genre
	if err := d.db.Select(&g, `select genre from products_genres where product_id=$1`, productId); err != nil {
		return nil, err
	}

	p.Genres = g

	return &p, nil
}

func (d *dbPSQL) GetProductWithDeleted(productId int, isDeleted bool) (*models.Product, error) {
	var p models.Product

	stmt := `select id, title, description, year, studio, rating, created_at, updated_at, 
			 coalesce((SELECT round(avg(score)) from reviews where product_id = $1), 0) as score from products
			 where id=$1 and is_deleted=$2
			 group by id
			 limit 1`

	if err := d.db.Get(&p, stmt, productId, isDeleted); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	var g []models.Genre
	if err := d.db.Select(&g, `select genre from products_genres where product_id=$1`, productId); err != nil {
		return nil, err
	}

	p.Genres = g

	return &p, nil
}

func (d *dbPSQL) GetProducts(after int, limit int, year int, genre string, isDeleted bool) ([]models.Product, error) {
	p := make([]models.Product, 0)

	stmt := `select p.id, p.title, p.description, p.year, p.studio, p.rating, p.created_at, p.updated_at, 
			 coalesce((SELECT round(avg(score)) from reviews), 0) as score from products p
			 join products_genres pg on p.id=pg.product_id`

	where := make([]string, 0)

	if year > 0 {
		where = append(where, fmt.Sprintf("p.year = %d", year))
	}

	if genre != "" {
		where = append(where, fmt.Sprintf("pg.genre = '%s'", genre))
	}

	where = append(where, fmt.Sprintf("is_deleted = %t", isDeleted))

	stmt = stmt + " WHERE " + strings.Join(where, " AND ")

	stmt += " GROUP BY p.id LIMIT $1 OFFSET $2"
	err := d.db.Select(&p, stmt, limit, after)
	if err != nil {
		return nil, err
	}

	stmt = `select genre from products_genres where product_id=$1`
	for i, v := range p {
		var g []models.Genre
		err = d.db.Select(&g, stmt, v.ID)
		if err != nil {
			return nil, err
		}

		p[i].Genres = g
	}

	return p, nil
}

func (d *dbPSQL) UpdateReview(rc models.ReviewUpdate) error {
	_, err := d.db.Exec("update reviews set score=$1, content=$2, content_html=$3 where id=$4", rc.Score, rc.Content, rc.ContentHTML, rc.ID)
	return err
}

func (d *dbPSQL) DeleteReview(reviewID int) error {
	_, err := d.db.Exec("update reviews set is_deleted=true where id=$1", reviewID)
	return err
}

func (d *dbPSQL) InsertReview(rc models.ReviewCreate, userID uuid.UUID) error {
	tx, err := d.db.Beginx()
	defer func() {
		defer func() {
			if err != nil {
				err = tx.Rollback()
				if err != nil {
					d.ILogger.Error(err)
				}
			}
		}()
	}()

	_, err = tx.Exec("INSERT INTO reviews(score, content, content_html, user_id, product_id) VALUES($1, $2, $3, $4, $5)", rc.Score, rc.Content, rc.ContentHTML, userID, rc.ProductID)
	if driverErr, ok := err.(*pq.Error); ok {
		if driverErr.Code == foreignKeyViolation {
			return ErrForeignKeyViolation
		}
	}

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (d *dbPSQL) DeleteProduct(productID int) error {
	_, err := d.db.Exec("update products set is_deleted=true where id = $1", productID)
	return err
}

func (d *dbPSQL) UpdateProduct(p models.ProductUpdate) error {
	tx, err := d.db.Beginx()
	defer func() {
		defer func() {
			if err != nil {
				err = tx.Rollback()
				if err != nil {
					d.ILogger.Error(err)
				}
			}
		}()
	}()
	_, err = tx.Exec("update products set title = $1, description = $2, year = $3, studio = $4, rating = $5 where id = $6", p.Title, p.Description, p.Year, p.Studio, p.Rating, p.ID)
	if driverErr, ok := err.(*pq.Error); ok {
		if driverErr.Code == foreignKeyViolation {
			return ErrForeignKeyViolation
		}
	}
	if err != nil {
		return err
	}

	_, err = tx.Exec("delete from products_genres where product_id = $1", p.ID)
	if err != nil {
		return err
	}

	if p.Genres == nil || len(p.Genres) == 0 {
		return nil
	}

	for _, v := range p.Genres {
		_, err = tx.Exec("insert into products_genres(product_id, genre) VALUES($1, $2)", p.ID, v.Genre)
		if driverErr, ok := err.(*pq.Error); ok {
			if driverErr.Code == foreignKeyViolation {
				return ErrForeignKeyViolation
			}
		}
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (d *dbPSQL) InsertProduct(p models.ProductCreate) error {
	tx, err := d.db.Beginx()
	defer func() {
		defer func() {
			if err != nil {
				err = tx.Rollback()
				if err != nil {
					d.ILogger.Error(err)
				}
			}
		}()
	}()

	var productID int
	err = tx.QueryRow("insert into products(title, description, year, studio, rating) values($1, $2, $3, $4, $5) returning id", p.Title, p.Description, p.Year, p.Studio, p.Rating).Scan(&productID)
	if driverErr, ok := err.(*pq.Error); ok {
		if driverErr.Code == foreignKeyViolation {
			return ErrForeignKeyViolation
		}
	}
	if err != nil {
		return err
	}

	for _, v := range p.Genres {
		_, err = tx.Exec("insert into products_genres(product_id, genre) VALUES($1, $2)", productID, v.Genre)
		if driverErr, ok := err.(*pq.Error); ok {
			if driverErr.Code == foreignKeyViolation {
				return ErrForeignKeyViolation
			}
		}
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (d *dbPSQL) GetUserByUsername(username string) (*models.User, error) {
	var u models.User

	if err := d.db.Get(&u, "select id, first_name, last_name, user_name, email, password, role, created_at, updated_at from users where user_name=$1", username); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}

func (d *dbPSQL) GetUserByID(id uuid.UUID) (*models.User, error) {
	var u models.User

	if err := d.db.Get(&u, "select id, first_name, last_name, user_name, email, password, role, created_at, updated_at from users where id=$1", id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}

func (d *dbPSQL) InsertUser(u models.CreateUser) (uuid.UUID, error) {
	var userID uuid.UUID

	err := d.db.Get(&userID, "insert into users(first_name, last_name, user_name, email, password, role) values($1, $2, $3, $4, $5, 'Registered') RETURNING id", u.FirstName, u.LastName, u.UserName, u.Email, u.Password)
	if driverErr, ok := err.(*pq.Error); ok {
		if driverErr.Code == foreignKeyViolation {
			return uuid.UUID{}, ErrForeignKeyViolation
		}

		if driverErr.Code == uniqueKeyViolation {
			return uuid.UUID{}, ErrUniqueViolation
		}
	}

	return userID, err
}

func (d *dbPSQL) InsertToken(userId uuid.UUID, refreshToken string) error {
	_, err := d.db.Exec("insert into users_tokens(user_id, refresh_token) VALUES($1, $2)", userId, refreshToken)
	if driverErr, ok := err.(*pq.Error); ok {
		if driverErr.Code == uniqueKeyViolation {
			return ErrUniqueViolation
		}
	}
	return err
}

func (d *dbPSQL) UpdateToken(userId uuid.UUID, refreshToken string) error {
	_, err := d.db.Exec("UPDATE users_tokens SET refresh_token=$1 WHERE user_id=$2", refreshToken, userId)
	if driverErr, ok := err.(*pq.Error); ok {
		if driverErr.Code == uniqueKeyViolation {
			return ErrUniqueViolation
		}
	}
	return err
}

func (d *dbPSQL) GetToken(userId uuid.UUID) (*models.Token, error) {
	var res models.Token
	if err := d.db.Get(&res, "select user_id, refresh_token from users_tokens where user_id=$1 limit 1", userId); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &res, nil
}

func (d *dbPSQL) DeleteToken(token string) error {
	_, err := d.db.Exec("delete from users_tokens where refresh_token=$1", token)
	return err
}

func (d *dbPSQL) FindToken(token string) (*models.Token, error) {
	var res models.Token
	if err := d.db.Get(&res, "select user_id, refresh_token from users_tokens where user_id=$1 limit 1", token); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &res, nil
}

func (d *dbPSQL) SelectGenres() ([]models.Genre, error) {
	g := make([]models.Genre, 0)

	if err := d.db.Select(&g, `SELECT name as genre from genres`); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return g, nil
}

func (d *dbPSQL) UpdateUserRole(userId uuid.UUID, role string) error {
	_, err := d.db.Exec(`UPDATE users SET role=$1 WHERE id=$2`, role, userId)
	if driverErr, ok := err.(*pq.Error); ok {
		if driverErr.Code == foreignKeyViolation {
			return ErrForeignKeyViolation
		}
	}
	return err
}

func (d *dbPSQL) GetUsers() ([]models.User, error) {
	u := make([]models.User, 0)

	stmt := `select id, first_name, last_name, user_name, email, role, created_at, updated_at from users`

	if err := d.db.Select(&u, stmt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return u, nil
}
