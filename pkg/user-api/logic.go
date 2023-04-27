package user_api

import (
	"github.com/google/uuid"
	"github.com/yalagtyarzh/aggregator/pkg/models"
)

type IUserAPILogic interface {
	GetReviews(productId int) ([]models.Review, error)
	GetProduct(productId int) (models.Product, error)
	GetProducts(after, limit, year int, genre string) ([]models.Product, error)
	CreateReview(rc models.ReviewCreate, userID uuid.UUID) error
	UpdateReview(rc models.ReviewUpdate, id uuid.UUID) error
	CreateUser(req models.CreateUser) error
}
