package repo

import (
	"github.com/google/uuid"
	"github.com/yalagtyarzh/aggregator/pkg/models"
)

type IDB interface {
	GetReviewsByProductID(productId int) ([]models.Review, error)
	GetReviewByID(reviewID int) (*models.Review, error)
	GetReviewByUserAndProduct(productId int, userId uuid.UUID) (*models.Review, error)
	GetProductWithDeleted(productId int, isDeleted bool) (*models.Product, error)
	GetProduct(productId int) (*models.Product, error)
	GetProducts(after int, limit int, year int, genre string, isDeleted bool) ([]models.Product, error)
	UpdateReview(rc models.ReviewUpdate) error
	DeleteReview(reviewID int) error
	InsertReview(rc models.ReviewCreate, userID uuid.UUID) error
	DeleteProduct(productID int) error
	UpdateProduct(p models.ProductUpdate) error
	InsertProduct(p models.ProductCreate) error
	GetUserByUsername(username string) (*models.User, error)
	GetUserByID(id uuid.UUID) (*models.User, error)
	InsertUser(u models.CreateUser) (uuid.UUID, error)
	InsertToken(userId uuid.UUID, refreshToken string) error
	GetToken(userId uuid.UUID) (*models.Token, error)
	UpdateToken(userId uuid.UUID, refreshToken string) error
	DeleteToken(token string) error
	FindToken(token string) (*models.Token, error)
	SelectGenres() ([]models.Genre, error)
	UpdateUserRole(userId uuid.UUID, role string) error
	GetUsers() ([]models.User, error)
}
