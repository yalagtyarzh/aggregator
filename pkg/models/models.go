package models

import "time"

type Product struct {
	ID          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Year        int       `json:"year" db:"year"`
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
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"UpdatedAt" db:"updated_at"`
}

type Score struct {
	Score int `json:"score" db:"score"`
}
