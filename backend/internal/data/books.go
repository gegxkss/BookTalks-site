package data

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gegxkss/BookTalks-site/backend/internal/domain"
)

// GetBookDetails получает детальную информацию о книге по ID
func GetBookDetails(bookID int) (domain.BookDetails, error) {
	var bookDetails domain.BookDetails
	var book domain.Book
	var genre domain.Genre
	var author domain.Author

	// Получение информации о книге, авторе и жанре
	query := `
        SELECT
            b.id, b.name, b.genre_id, b.author_id, b.coverimage_filename, b.created_at,
            g.id AS genre_id, g.name AS genre_name,
            a.id AS author_id, a.first_name AS author_first_name, a.last_name AS author_last_name, a.surname AS author_surname
        FROM
            book b
        JOIN
            genre g ON b.genre_id = g.id
        JOIN
            author a ON b.author_id = a.id
        WHERE
            b.id = $1
    `

	err := DB.QueryRow(query, bookID).Scan(
		&book.ID,
		&book.Name,
		&book.GenreID,
		&book.AuthorID,
		&book.CoverImageFilename,
		&book.CreatedAt,
		&genre.ID,
		&genre.Name,
		&author.ID,
		&author.FirstName,
		&author.LastName,
		&author.Surname,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.BookDetails{}, fmt.Errorf("book not found")
		}
		log.Printf("Error getting book details: %v", err)
		return domain.BookDetails{}, err
	}

	bookDetails.ID = book.ID
	bookDetails.Name = book.Name
	bookDetails.Genre = genre
	bookDetails.Author = author
	bookDetails.CoverImageFilename = book.CoverImageFilename
	bookDetails.CreatedAt = book.CreatedAt

	// Получение цитат
	quotes, err := getQuotesForBook(bookID)
	if err != nil {
		log.Printf("Error getting quotes for book: %v", err)
	}
	bookDetails.Quotes = quotes

	// Получение рецензий
	reviews, err := getReviewsForBook(bookID)
	if err != nil {
		log.Printf("Error getting reviews for book: %v", err)
	}
	bookDetails.Reviews = reviews

	// Получение среднего рейтинга
	rating, err := getAverageRatingForBook(bookID)
	if err != nil {
		log.Printf("Error getting average rating for book: %v", err)
	}
	bookDetails.Rating = rating

	return bookDetails, nil
}

// getQuotesForBook получает цитаты для книги
func getQuotesForBook(bookID int) ([]domain.Quote, error) {
	query := `
        SELECT
            id, book_id, user_id, text
        FROM
            quote
        WHERE
            book_id = $1
    `
	rows, err := DB.Query(query, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotes []domain.Quote
	for rows.Next() {
		var quote domain.Quote
		err := rows.Scan(&quote.ID, &quote.BookID, &quote.UserID, &quote.Text)
		if err != nil {
			return nil, err
		}
		quotes = append(quotes, quote)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return quotes, nil
}

// getReviewsForBook получает рецензии для книги
func getReviewsForBook(bookID int) ([]domain.Review, error) {
	query := `
        SELECT
            id, user_id, book_id, text, created_at
        FROM
            review
        WHERE
            book_id = $1
    `
	rows, err := DB.Query(query, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []domain.Review
	for rows.Next() {
		var review domain.Review
		err := rows.Scan(&review.ID, &review.UserID, &review.BookID, &review.Text, &review.CreatedAt)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return reviews, nil
}

// getAverageRatingForBook получает средний рейтинг для книги
func getAverageRatingForBook(bookID int) (float64, error) {
	query := `
        SELECT AVG(amount)
        FROM rating
        WHERE book_id = $1
    `

	var avgRating sql.NullFloat64
	err := DB.QueryRow(query, bookID).Scan(&avgRating)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	if avgRating.Valid {
		return avgRating.Float64, nil
	}

	return 0, nil // Если нет оценок, возвращаем 0
}
