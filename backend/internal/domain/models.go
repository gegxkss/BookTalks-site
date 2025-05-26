package domain

import "time"

type Genre struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Book struct {
	ID                 int       `db:"id"`
	Name               string    `db:"name"`
	GenreID            int       `db:"genre_id"`
	AuthorID           int       `db:"author_id"`
	CoverImageFilename string    `db:"coverimage_filename"`
	CreatedAt          time.Time `db:"created_at"`
}

type Author struct {
	ID        int    `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Surname   string `db:"surname"`
}

type User struct {
	ID        int       `db:"id"`
	Nickname  string    `db:"nickname"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Sex       string    `db:"sex"`
	BirthDate time.Time `db:"birth_date"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}

type Rating struct {
	ID     int `db:"id"`
	BookID int `db:"book_id"`
	UserID int `db:"user_id"`
	Amount int `db:"amount"`
}

type Quote struct {
	ID     int    `db:"id"`
	BookID int    `db:"book_id"`
	UserID int    `db:"user_id"`
	Text   string `db:"text"`
}

type Review struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	BookID    int       `db:"book_id"`
	Text      string    `db:"text"`
	CreatedAt time.Time `db:"created_at"`
}

type UserBook struct {
	ID     int `db:"id"`
	UserID int `db:"user_id"`
	BookID int `db:"book_id"`
}
