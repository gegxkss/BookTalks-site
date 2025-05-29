package data

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv" // Импорт godotenv
	_ "github.com/lib/pq"      // Импорт драйвера PostgreSQL
)

// Переменная для хранения подключения к базе данных
var DB *sql.DB

// ConnectDB устанавливает соединение с базой данных
func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	var err2 error
	DB, err2 = sql.Open("postgres", dbURL)
	if err2 != nil {
		log.Fatal("Error connecting to database: ", err2)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err)
	}

	log.Println("Connected to the database")
}
