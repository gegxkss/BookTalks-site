package data

// UpdateRating обновляет рейтинг книги
func UpdateRating(bookID int, userID int, rating int) error {
	// Сначала проверяем, есть ли уже оценка от этого пользователя для этой книги
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM rating WHERE book_id = $1 AND user_id = $2", bookID, userID).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		// Если оценка уже есть, обновляем её
		query := `
			UPDATE rating
			SET amount = $3
			WHERE book_id = $1 AND user_id = $2
		`
		_, err = DB.Exec(query, bookID, userID, rating)
		return err
	} else {
		// Если оценки нет, добавляем новую
		query := `
			INSERT INTO rating (book_id, user_id, amount)
			VALUES ($1, $2, $3)
		`
		_, err = DB.Exec(query, bookID, userID, rating)
		return err
	}
}
