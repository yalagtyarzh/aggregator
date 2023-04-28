package repo

import (
	"github.com/google/uuid"
	"github.com/yalagtyarzh/aggregator/pkg/models"
)

type IDB interface {
	GetReviewsByProductID(productId int) ([]models.Review, error)
	GetReviewByID(reviewID int) (*models.Review, error)
	GetProductById(productId int) (*models.Product, error)
	GetProducts(after int, limit int, year int, genre string) ([]models.Product, error)
	GetPermissionsByRole(userID uuid.UUID) ([]models.Permission, error)
	UpdateReview(rc models.ReviewUpdate) error
	DeleteReview(reviewID int) error
	InsertReview(rc models.ReviewCreate, userID uuid.UUID) error
	DeleteProduct(productID int) error
	UpdateProduct(p models.ProductUpdate) error
	InsertProduct(p models.ProductCreate) error
	GetUserByEmail(email string) (*models.User, error)
	InsertUser(u models.CreateUser) (uuid.UUID, error)
	InsertToken(userId uuid.UUID, refreshToken string) error
	GetToken(userId uuid.UUID) (*models.Token, error)
	UpdateToken(userId uuid.UUID, refreshToken string) error
}
