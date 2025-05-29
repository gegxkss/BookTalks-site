package data

// AddQuote добавляет цитату к книге
func AddQuote(bookID int, userID int, text string) error {
	query := `
		INSERT INTO quote (book_id, user_id, text)
		VALUES ($1, $2, $3)
	`
	_, err := DB.Exec(query, bookID, userID, text)
	return err
}
