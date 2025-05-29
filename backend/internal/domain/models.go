package domain

import (
	"database/sql"

	"time"
)

type Genre struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type Book struct {
	ID                 int       `db:"id" json:"id"`
	Name               string    `db:"name" json:"name"`
	GenreID            int       `db:"genre_id" json:"genre_id"`
	AuthorID           int       `db:"author_id" json:"author_id"`
	CoverImageFilename string    `db:"coverimage_filename" json:"coverimage_filename"`
	CreatedAt          time.Time `db:"created_at" json:"created_at"`
}

type Author struct {
	ID        int    `db:"id" json:"id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
	Surname   string `db:"surname" json:"surname"` // Оставил Surname, как есть
}

type User struct {
	ID                   int
	Nickname             string
	FirstName            string
	LastName             string
	Sex                  string
	BirthDate            time.Time
	Email                string
	Password             string
	CreatedAt            time.Time
	ProfileImageFileName sql.NullString // Изменено на sql.NullString
}

type Rating struct {
	ID     int `db:"id" json:"id"`
	BookID int `db:"book_id" json:"book_id"`
	UserID int `db:"user_id" json:"user_id"`
	Amount int `db:"amount" json:"amount"`
}

type Quote struct {
	ID     int    `db:"id" json:"id"`
	BookID int    `db:"book_id" json:"book_id"`
	UserID int    `db:"user_id" json:"user_id"`
	Text   string `db:"text" json:"text"`
}

type Review struct {
	ID        int       `db:"id" json:"id"`
	UserID    int       `db:"user_id" json:"user_id"`
	BookID    int       `db:"book_id" json:"book_id"`
	Text      string    `db:"text" json:"text"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type UserBook struct {
	ID     int `db:"id" json:"id"`
	UserID int `db:"user_id" json:"user_id"`
	BookID int `db:"book_id" json:"book_id"`
}

type BookDetails struct {
	ID                 int       `json:"id"`
	Name               string    `json:"name"`
	Genre              Genre     `json:"genre"`
	Author             Author    `json:"author"`
	CoverImageFilename string    `json:"coverimage_filename"`
	CreatedAt          time.Time `json:"created_at"`
	Quotes             []Quote   `json:"quotes"`
	Reviews            []Review  `json:"reviews"`
	Rating             float64   `json:"rating"`
}

// RegisterRequest представляет данные для регистрации нового пользователя
type RegisterRequest struct {
	Nickname             string    `json:"nickname"`
	FirstName            string    `json:"first_name"`
	LastName             string    `json:"last_name"`
	Sex                  string    `json:"sex"`
	BirthDate            time.Time `json:"birth_date"`
	Email                string    `json:"email"`
	Password             string    `json:"password"`
	ProfileImageFileName string    `json:"profile_image_filename"`
}

// LoginRequest представляет данные для входа пользователя
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
