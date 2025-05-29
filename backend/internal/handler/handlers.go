package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gegxkss/BookTalks-site/backend/internal/data"
	"github.com/gorilla/mux"
)

// POST /api/books - Добавление новой книги
func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Валидация (пример, нужно расширить)
	if book.Name == "" || book.AuthorID == 0 || book.GenreID == 0 { // Updated validation
		http.Error(w, "Name, AuthorID and GenreID are required", http.StatusBadRequest)
		return
	}

	// Сохранение книги в базе данных (предполагается, что у вас есть функция в data пакете)
	newBook, err := data.CreateBook(book) // Assuming you have a CreateBook function
	if err != nil {
		log.Println("Error creating book:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newBook); err != nil {
		log.Println("Error encoding JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// Получение списка книг
func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := data.GetBooks()
	if err != nil {
		log.Println("Error getting books:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(books); err != nil {
		log.Println("Error encoding JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// PUT Обновление информации о книге
func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	book.ID = id

	err = data.UpdateBook(book)
	if err != nil {
		log.Println("Error updating book:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// Удаление книги
func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	err = data.DeleteBook(id)
	if err != nil {
		log.Println("Error deleting book:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content (успешное удаление)
}

// handler/handler
// Constants
const (
	MaxUploadSize = 10 * 1024 * 1024 // 10MB
	UploadPath    = "./uploads"      // Путь для сохранения загруженных файлов
)

// Обработчик для страницы профиля (требуется авторизация)
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("./frontend", "profile.html") // Используем filepath.Join
	http.ServeFile(w, r, filePath)
}

// Обработчик для главной страницы (общедоступно)

// Обработчик для страницы регистрации (общедоступно)
func RegistrationPageHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("./frontend", "registration.html") // Используем filepath.Join
	http.ServeFile(w, r, filePath)
}

func QuoteHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("./frontend", "quote.html") // Используем filepath.Join
	http.ServeFile(w, r, filePath)
}

// Обработчик для страницы входа (общедоступно)
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("./frontend", "signin.html") // Используем filepath.Join
	http.ServeFile(w, r, filePath)
}

// generateUUID generates a unique UUID.
func generateUUID() (string, error) {
	uuid := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, uuid)
	if err != nil {
		return "", fmt.Errorf("failed to generate UUID: %w", err)
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return hex.EncodeToString(uuid), nil
}

// Обработчик для главной страницы (общедоступно)

// generateUniqueFilename generates a unique filename with extension.
func generateUniqueFilename(ext string) (string, error) {
	uuid, err := generateUUID()
	if err != nil {
		return "", fmt.Errorf("failed to generate unique filename: %w", err)
	}
	return uuid + ext, nil
}

// init creates the uploads directory if it doesn't exist.
func init() {
	if _, err := os.Stat(UploadPath); os.IsNotExist(err) {
		if err := os.MkdirAll(UploadPath, 0755); err != nil {
			log.Fatalf("Failed to create upload directory: %v", err)
		}
	}
}

// --- Handler Functions ---

// GetBooksHandler retrieves a list of all books.
func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := []domain.Book{
		{ID: 1, Name: "Книга 1", AuthorID: 1, GenreID: 1},
		{ID: 2, Name: "Книга 2", AuthorID: 2, GenreID: 2},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(books); err != nil {
		log.Printf("Failed to encode books to JSON: %v", err)
		http.Error(w, "Failed to encode books", http.StatusInternalServerError)
	}
}

// GetBookDetailsHandler retrieves details for a specific book.
func GetBookDetailsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookIDStr := vars["book_id"]

	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	bookDetails, err := data.GetBookDetails(bookID)
	if err != nil {
		if err.Error() == "book not found" {
			http.Error(w, "Book not found", http.StatusNotFound)
		} else {
			log.Printf("Failed to get book details: %v", err)
			http.Error(w, "Failed to get book details", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(bookDetails); err != nil {
		log.Printf("Failed to encode book details to JSON: %v", err)
		http.Error(w, "Failed to encode book details", http.StatusInternalServerError)
	}
}

// AddQuoteHandler adds a quote to a book (requires authentication).
func AddQuoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookIDStr := vars["book_id"]

	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("userId")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userIDStr := cookie.Value

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID in cookie", http.StatusBadRequest)
		return
	}

	type AddQuoteRequest struct {
		Text string `json:"text"`
	}

	var req AddQuoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = data.AddQuote(bookID, userID, req.Text)
	if err != nil {
		log.Printf("Error adding quote: %v", err)
		http.Error(w, "Failed to add quote", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"status": "success"}); err != nil {
		log.Printf("Failed to encode success response to JSON: %v", err)
		http.Error(w, "Failed to encode success response", http.StatusInternalServerError)
	}
}

// AddReviewHandler adds a review to a book (requires authentication).
func AddReviewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookIDStr := vars["book_id"]

	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("userId")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userIDStr := cookie.Value

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID in cookie", http.StatusBadRequest)
		return
	}

	type AddReviewRequest struct {
		Text string `json:"text"`
	}

	var req AddReviewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = data.AddReview(bookID, userID, req.Text)
	if err != nil {
		log.Printf("Error adding review: %v", err)
		http.Error(w, "Failed to add review", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"status": "success"}); err != nil {
		log.Printf("Failed to encode success response to JSON: %v", err)
		http.Error(w, "Failed to encode success response", http.StatusInternalServerError)
	}
}

// GetUserHandler retrieves user information by ID (requires authentication).
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["user_id"]

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("userId")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userIDFromCookieStr := cookie.Value

	userIDFromCookie, err := strconv.Atoi(userIDFromCookieStr)
	if err != nil {
		http.Error(w, "Invalid user ID in cookie", http.StatusBadRequest)
		return
	}

	if userID != userIDFromCookie {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := data.GetUser(userID)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("Failed to encode user to JSON: %v", err)
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
	}
}

// RegisterHandler handles user registration.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Parse multipart form
	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
		log.Printf("Failed to parse multipart form: %v", err)
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}

	// 2. Get form values
	nickname := r.FormValue("nickname")
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	sex := r.FormValue("sex")
	birthDateStr := r.FormValue("birth_date")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// 3. Parse birth date
	var birthDate time.Time
	if birthDateStr != "" {
		var err error
		birthDate, err = time.Parse("2006-01-02", birthDateStr)
		if err != nil {
			http.Error(w, "Invalid birth date format", http.StatusBadRequest)
			return
		}
	}

	// 4. Handle profile image upload
	var filename string
	file, header, err := r.FormFile("profile_image")
	if err == nil { // File was uploaded
		defer file.Close()

		// Validate file size
		if header.Size > MaxUploadSize {
			http.Error(w, "File size exceeds maximum limit", http.StatusBadRequest)
			return
		}

		// Generate unique filename
		ext := filepath.Ext(header.Filename)
		filename, err = generateUniqueFilename(ext)
		if err != nil {
			log.Printf("Failed to generate unique filename: %v", err)
			http.Error(w, "Failed to generate unique filename", http.StatusInternalServerError)
			return
		}

		// Create file to save the image
		filePath := filepath.Join(UploadPath, filename)
		dst, err := os.Create(filePath)
		if err != nil {
			log.Printf("Failed to create file: %v", err)
			http.Error(w, "Failed to save file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		// Copy uploaded file content to the created file
		_, err = io.Copy(dst, file)
		if err != nil {
			log.Printf("Failed to save file: %v", err)
			http.Error(w, "Failed to save file", http.StatusInternalServerError)
			return
		}

		log.Printf("File received: %s, Size: %d bytes", header.Filename, header.Size)

	} else if err != http.ErrMissingFile { // Error during upload
		log.Printf("Error retrieving file: %v", err)
		http.Error(w, "Error retrieving file", http.StatusInternalServerError)
		return
	}

	// 5. Create RegisterRequest
	req := domain.RegisterRequest{
		Nickname:             nickname,
		FirstName:            firstName,
		LastName:             lastName,
		Sex:                  sex,
		BirthDate:            birthDate,
		Email:                email,
		Password:             password,
		ProfileImageFileName: filename, // Save filename
	}
	log.Printf("RegisterRequest: %+v", req)

	// 6. Register user
	userID, err := data.RegisterUser(req)
	if err != nil {
		log.Printf("Error registering user: %v", err)
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	// 7. Send JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created
	if err := json.NewEncoder(w).Encode(map[string]string{"status": "success", "user_id": strconv.Itoa(userID)}); err != nil {
		log.Printf("Failed to encode success response to JSON: %v", err)
		http.Error(w, "Failed to encode success response", http.StatusInternalServerError)
	}
}

// LoginHandler handles user login.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req domain.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userID, err := data.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		log.Printf("Error authenticating user: %v", err)
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	data.SetUserCookie(w, userID)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"status": "success", "user_id": strconv.Itoa(userID)}); err != nil {
		log.Printf("Failed to encode success response to JSON: %v", err)
		http.Error(w, "Failed to encode success response", http.StatusInternalServerError)
	}
}
