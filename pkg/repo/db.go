package repo

import "github.com/yalagtyarzh/aggregator/pkg/models"

type IDB interface {
	GetReviewsByProductID(productId int) ([]models.Review, error)
	GetProductById(productId int) (*models.Product, error)
	GetProductScoreById(productId int) (*models.Score, error)
	GetProducts(after int, limit int, year int, genre string) ([]models.Product, error)
}
