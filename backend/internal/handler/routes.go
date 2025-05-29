package handler

import (
	"github.com/gorilla/mux"
)

// SetupRoutes настраивает маршруты для API
func SetupRoutes(r *mux.Router) {
	// Определение маршрутов
	r.HandleFunc("/register", RegisterHandler).Methods("POST")
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/books", GetBooks).Methods("GET")
	r.HandleFunc("/books/{book_id}", GetBookDetailsHandler).Methods("GET")
	r.HandleFunc("/books/{book_id}/quotes/add", AddQuoteHandler).Methods("POST")
	r.HandleFunc("/books/{book_id}/reviews/add", AddReviewHandler).Methods("POST")
	r.HandleFunc("/users/{user_id}", GetUserHandler).Methods("GET")
}
