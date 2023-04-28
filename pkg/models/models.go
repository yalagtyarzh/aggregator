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
	ReleaseDate time.Time `json:"releaseDate" db:"release_date"`
	Studio      string    `json:"studio" db:"studio"`
	Rating      string    `json:"rating" db:"rating"`
	Score       int       `json:"score" db:"score"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"UpdatedAt" db:"updated_at"`
}

type ProductCreate struct {
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Year        int       `json:"year" db:"year"`
	Genres      []Genre   `json:"genres"`
	ReleaseDate time.Time `json:"releaseDate" db:"release_date"`
	Studio      string    `json:"studio" db:"studio"`
	Rating      string    `json:"rating" db:"rating"`
}

type ProductUpdate struct {
	ID          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Year        int       `json:"year" db:"year"`
	Genres      []Genre   `json:"genres"`
	ReleaseDate time.Time `json:"releaseDate" db:"release_date"`
	Studio      string    `json:"studio" db:"studio"`
	Rating      string    `json:"rating" db:"rating"`
	Delete      bool      `json:":delete"`
}

type Review struct {
	ID          int       `json:"id" db:"id"`
	Score       int       `json:"score" db:"score"`
	Content     string    `json:"content" db:"content"`
	ContentHTML string    `json:"contentHTML" db:"content_html"`
	User        User      `json:"user"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"UpdatedAt" db:"updated_at"`
}

type ReviewCreate struct {
	ProductID   int    `json:"productId"`
	Score       int    `json:"score"`
	Content     string `json:"content"`
	ContentHTML string `json:"contentHTML"`
}

type ReviewUpdate struct {
	ID          int    `json:"id"`
	Score       int    `json:"score"`
	Content     string `json:"content"`
	ContentHTML string `json:"contentHTML"`
	Delete      bool   `json:"bool"`
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
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

type CreateUserResponse struct {
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

type Permission struct {
	Permission string `json:"permission" db:"permission"`
}

type TokenPayload struct {
	UserID uuid.UUID `json:"userId"`
	Email  string    `json:"email"`
	jwt.StandardClaims
}
