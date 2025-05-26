package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"backend/internal/data"

	"github.com/gorilla/mux"
)

// GetBooks handler
func GetBooks(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	rows, err := data.DB.Query(ctx, "SELECT id, name, genre_id, author_id, coverimage_filename, created_at FROM book")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []map[string]interface{} // Используем map[string]interface{} для гибкости

	for rows.Next() {
		var id, genreID, authorID int
		var name, coverImageFilename string
		var createdAt string

		err = rows.Scan(&id, &name, &genreID, &authorID, &coverImageFilename, &createdAt)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		book := map[string]interface{}{
			"id":                  id,
			"name":                name,
			"genre_id":            genreID,
			"author_id":           authorID,
			"coverimage_filename": coverImageFilename,
			"created_at":          createdAt,
		}
		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GetBook handler
func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookIDStr := params["id"] // Get the book ID as a string

	// Convert bookID to integer
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var id, genreID, authorID int
	var name, coverImageFilename string
	var createdAt string

	query := "SELECT id, name, genre_id, author_id, coverimage_filename, created_at FROM book WHERE id = $1"

	ctx := context.Background() // Get context

	// Pass context to the QueryRow
	err = data.DB.QueryRow(ctx, query, bookID).Scan(&id, &name, &genreID, &authorID, &coverImageFilename, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
			return
		}
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	book := map[string]interface{}{
		"id":                  id,
		"name":                name,
		"genre_id":            genreID,
		"author_id":           authorID,
		"coverimage_filename": coverImageFilename,
		"created_at":          createdAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// CreateBook handler
func CreateBook(w http.ResponseWriter, r *http.Request) {
	// Extract data from the request body (assuming JSON)
	var newBook map[string]interface{} // Use a map to hold the incoming JSON

	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Extract values from the map
	name, ok := newBook["name"].(string)
	if !ok {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	genreIDFloat, ok := newBook["genre_id"].(float64)
	if !ok {
		http.Error(w, "Genre ID is required and must be an integer", http.StatusBadRequest)
		return
	}
	genreID := int(genreIDFloat) // Convert float64 to int

	authorIDFloat, ok := newBook["author_id"].(float64)
	if !ok {
		http.Error(w, "Author ID is required and must be an integer", http.StatusBadRequest)
		return
	}
	authorID := int(authorIDFloat) // Convert float64 to int

	coverImageFilename, _ := newBook["coverimage_filename"].(string) // Optional field

	// SQL INSERT statement
	query := `
        INSERT INTO book (name, genre_id, author_id, coverimage_filename)
        VALUES ($1, $2, $3, $4)
        RETURNING id, name, genre_id, author_id, coverimage_filename, created_at
    `

	var id int
	var createdAt string

	ctx := context.Background() // Get context

	// Pass context to the QueryRow
	err = data.DB.QueryRow(ctx, query, name, genreID, authorID, coverImageFilename).Scan(&id, &name, &genreID, &authorID, &coverImageFilename, &createdAt)
	if err != nil {
		log.Println("DB Insert Error:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Create a response with the new book
	createdBook := map[string]interface{}{
		"id":                  id,
		"name":                name,
		"genre_id":            genreID,
		"author_id":           authorID,
		"coverimage_filename": coverImageFilename,
		"created_at":          createdAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // Set the status to 201
	json.NewEncoder(w).Encode(createdBook)
}

// UpdateBook handler
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookIDStr := params["id"]

	// Convert bookID to integer
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var updatedBook map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	name, ok := updatedBook["name"].(string)
	if !ok {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	genreIDFloat, ok := updatedBook["genre_id"].(float64)
	if !ok {
		http.Error(w, "Genre ID is required and must be an integer", http.StatusBadRequest)
		return
	}
	genreID := int(genreIDFloat)

	authorIDFloat, ok := updatedBook["author_id"].(float64)
	if !ok {
		http.Error(w, "Author ID is required and must be an integer", http.StatusBadRequest)
		return
	}
	authorID := int(authorIDFloat)

	coverImageFilename, _ := updatedBook["coverimage_filename"].(string)

	query := `
        UPDATE book
        SET name = $1, genre_id = $2, author_id = $3, coverimage_filename = $4
        WHERE id = $5
        RETURNING id, name, genre_id, author_id, coverimage_filename, created_at
    `

	var id int
	var createdAt string

	ctx := context.Background() // Add context

	// Pass context to the QueryRow and use the converted bookID
	err = data.DB.QueryRow(ctx, query, name, genreID, authorID, coverImageFilename, bookID).Scan(&id, &name, &genreID, &authorID, &coverImageFilename, &createdAt)
	if err != nil {
		log.Println("DB Update Error:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	updatedBookResponse := map[string]interface{}{
		"id":                  id,
		"name":                name,
		"genre_id":            genreID,
		"author_id":           authorID,
		"coverimage_filename": coverImageFilename,
		"created_at":          createdAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBookResponse)
}

// DeleteBook handler
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookIDStr := params["id"]

	// Convert bookID to integer
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM book WHERE id = $1"

	ctx := context.Background() // Add context

	_, err = data.DB.Exec(ctx, query, bookID) // Use the converted bookID
	if err != nil {
		log.Println("DB Delete Error:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content
}
