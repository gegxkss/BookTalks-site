package main

import (
	"log"
	"net/http"

	"backend/internal/data"
	"backend/internal/handler"
	"github.com/gorilla/mux"
)

func main() {
	data.ConnectDB() // Подключение к базе данных
	defer data.DB.Close()

	err := data.MigrateDB() // Выполнение миграций
	if err != nil {
		log.Fatal("Could not migrate the database: ", err)
	}

	r := mux.NewRouter()
	handler.SetupRoutes(r) // Настройка маршрутов

	log.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r)) // Запуск сервера
}
