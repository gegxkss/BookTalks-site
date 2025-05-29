package data

// AddReview добавляет рецензию к книге
func AddReview(bookID int, userID int, text string) error {
	query := `
		INSERT INTO review (book_id, user_id, text)
		VALUES ($1, $2, $3)
	`
	_, err := DB.Exec(query, bookID, userID, text)
	return err
}
