package models

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Year        int       `json:"year" db:"year"`
	Genres      []Genre   `json:"genres"`
	Studio      string    `json:"studio" db:"studio"`
	Rating      string    `json:"rating" db:"rating"`
	Score       int       `json:"score" db:"score"`
	ImageLink   string    `json:"imageLink" db:"img_link"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"UpdatedAt" db:"updated_at"`
}

type ProductCreate struct {
	Title       string  `json:"title" db:"title" validate:"required"`
	Description string  `json:"description" db:"description" validate:"required"`
	Year        int     `json:"year" db:"year" validate:"required"`
	Genres      []Genre `json:"genres" validate:"required,min=1"`
	Studio      string  `json:"studio" db:"studio" validate:"required"`
	Rating      string  `json:"rating" db:"rating" validate:"required"`
	ImageLink   string  `json:"imageLink" db:"img_link" validate:"required,http_url"`
}

type ProductUpdate struct {
	ID          int     `json:"id" db:"id" validate:"required"`
	Title       string  `json:"title" db:"title" validate:"required"`
	Description string  `json:"description" db:"description" validate:"required"`
	Year        int     `json:"year" db:"year" validate:"required"`
	Genres      []Genre `json:"genres" validate:"required,min=1"`
	Studio      string  `json:"studio" db:"studio" validate:"required"`
	Rating      string  `json:"rating" db:"rating" validate:"required"`
	ImageLink   string  `json:"imageLink" db:"img_link" validate:"required,http_url"`
	Delete      bool    `json:"delete"`
}

type Review struct {
	ID          int       `json:"id" db:"id"`
	Score       int       `json:"score" db:"score"`
	Content     string    `json:"content" db:"content"`
	ContentHTML string    `json:"contentHTML" db:"content_html"`
	UserID      uuid.UUID `json:"userId" db:"user_id"`
	FirstName   string    `json:"firstName" db:"first_name"`
	LastName    string    `json:"lastName" db:"last_name"`
	UserName    string    `json:"userName" db:"user_name"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"UpdatedAt" db:"updated_at"`
}

type ReviewCreate struct {
	ProductID   int    `json:"productId" validate:"required"`
	Score       int    `json:"score" validate:"required,min=0,max=100"`
	Content     string `json:"content" validate:"required,max=1000"`
	ContentHTML string `json:"contentHTML" validate:"required"`
}

type ReviewUpdate struct {
	ID          int    `json:"id" validate:"required"`
	Score       int    `json:"score" validate:"required,min=0,max=100"`
	Content     string `json:"content" validate:"required,max=1000"`
	ContentHTML string `json:"contentHTML" validate:"required"`
	Delete      bool   `json:"delete"`
}

type User struct {
	ID        uuid.UUID `json:"userId" db:"id"`
	FirstName string    `json:"firstName" db:"first_name"`
	LastName  string    `json:"lastName" db:"last_name"`
	UserName  string    `json:"userName" db:"user_name"`
	Email     string    `json:"-" db:"email"`
	Password  string    `json:"-" db:"password"`
	Role      string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"UpdatedAt" db:"updated_at"`
}

type CreateUser struct {
	FirstName string `json:"firstName" validate:"required,min=2,max=32"`
	LastName  string `json:"lastName" validate:"required,min=2,max=32"`
	UserName  string `json:"userName" validate:"required,min=3,max=32"`
	Email     string `json:"email" validate:"required,email,max=150"`
	Password  string `json:"password" validate:"required,min=3,max=32"`
}

type LoginRequest struct {
	Username string `json:"userName" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=3,max=32"`
}

type UserResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	UserID       string `json:"userId"`
	Email        string `json:"email"`
}

type Token struct {
	UserID       uuid.UUID `json:"userId" db:"user_id"`
	RefreshToken string    `json:"refreshToken" db:"refresh_token"`
}

type Score struct {
	Score int `json:"score" db:"score"`
}

type Genre struct {
	Genre string `json:"genre" db:"genre"`
}

type TokenPayload struct {
	UserID uuid.UUID `json:"userId"`
	Email  string    `json:"email"`
	Role   string    `json:"role"`
	jwt.StandardClaims
}

type StdResp struct {
	Message string `json:"message"`
}

type PromoteDismissRequest struct {
	UserID uuid.UUID `json:"userId"`
}
