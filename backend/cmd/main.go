package main

import (
	"fmt"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gegxkss/BookTalks-site/backend/internal/data"
	"github.com/gegxkss/BookTalks-site/backend/internal/handler"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database
	data.ConnectDB()

	// Migrate the database
	err = data.MigrateDB()
	if err != nil {
		log.Fatal("Error migrating database: ", err)
	}

	// Set up the router
	router := mux.NewRouter()
	mime.AddExtensionType(".css", "text/css") // Add MIME type for CSS
	router.Use(enableCORS)                    // Apply CORS middleware

	// Serve static files
	frontendDir, err := filepath.Abs("./frontend")
	if err != nil {
		log.Fatal("Error getting absolute path to frontend directory:", err)
	}
	fs := http.FileServer(http.Dir(frontendDir))

	// Log requests to the file server
	loggedFS := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("FileServer: %s", r.URL.Path)
		fs.ServeHTTP(w, r)
	})

	//  обработчики для страниц профиля, регистрации и логина
	router.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(frontendDir, "profile.html"))
	})
	router.HandleFunc("/registration", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(frontendDir, "registration.html"))
	})
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(frontendDir, "signin.html"))
	})
	router.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(frontendDir, "addBook.html"))
	})
	router.HandleFunc("/library", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(frontendDir, "library.html"))
	})
	router.HandleFunc("/quote", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(frontendDir, "quote.html"))
	})
	router.HandleFunc("/review", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(frontendDir, "resenzii.html"))
	})
	router.HandleFunc("/recommendation", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(frontendDir, "recommend.html"))
	})
	//обработчик для главной страницы
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(frontendDir, "main.html"))
	})

	router.PathPrefix("/frontend/").Handler(http.StripPrefix("/frontend/", loggedFS)) // для css тупого этого

	handler.SetupRoutes(router) // Маршруты API

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
