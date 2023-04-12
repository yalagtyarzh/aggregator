package user_api

import "github.com/yalagtyarzh/aggregator/pkg/models"

type IUserAPILogic interface {
	GetReviews(productId int) ([]models.Review, error)
	GetProduct(productId int) (models.Product, error)
	GetProductScore(productId int) (models.Score, error)
}
