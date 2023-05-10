package user_api

import (
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"github.com/yalagtyarzh/aggregator/pkg/models"
)

type IUserAPILogic interface {
	GetReviews(productId int) ([]models.Review, error)
	GetProduct(productId int) (models.Product, error)
	GetProducts(after, limit, year int, genre string) ([]models.Product, error)
	CreateReview(rc models.ReviewCreate, userID uuid.UUID) error
	UpdateReview(rc models.ReviewUpdate, id uuid.UUID) error
	CreateUser(req models.CreateUser) (models.UserResponse, error)
	Login(username, password string) (models.UserResponse, error)
	Logout(token string) error
	Refresh(token string) (models.UserResponse, error)
	ListGenres() ([]models.Genre, error)
	GraphqlList(query string) (*graphql.Result, error)
}

var productType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Product",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"year": &graphql.Field{
			Type: graphql.Int,
		},
		"studio": &graphql.Field{
			Type: graphql.String,
		},
		"rating": &graphql.Field{
			Type: graphql.String,
		},
		"createdAt": &graphql.Field{
			Type: graphql.DateTime,
		},
		"updatedAt": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})
