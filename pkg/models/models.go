package models

import (
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
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"UpdatedAt" db:"updated_at"`
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
	Score       int    `json:"score" db:"score"`
	Content     string `json:"content" db:"content"`
	ContentHTML string `json:"contentHTML" db:"content_html"`
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

type Score struct {
	Score int `json:"score" db:"score"`
}

type Genre struct {
	Name string `json:"genre" db:"genre"`
}
