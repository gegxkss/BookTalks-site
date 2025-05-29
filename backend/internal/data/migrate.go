package data

import (
	"log"
)

// MigrateDB выполняет миграции базы данных (создает таблицы, если их нет)
func MigrateDB() error {
	// Genre Table
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS genre (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL
		);
	`)
	if err != nil {
		log.Printf("Error creating genre table: %s", err)
		return err
	}

	// Author Table
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS author (
			id SERIAL PRIMARY KEY,
			first_name VARCHAR(255) NOT NULL,
			last_name VARCHAR(255) NOT NULL,
			surname VARCHAR(255)
		);
	`)
	if err != nil {
		log.Printf("Error creating author table: %s", err)
		return err
	}

	// User Table
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			nickname VARCHAR(255) NOT NULL UNIQUE,
			first_name VARCHAR(255) NOT NULL,
			last_name VARCHAR(255) NOT NULL,
			sex VARCHAR(10),
			birth_date DATE,
			email VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT NOW()
		);
	`)
	if err != nil {
		log.Printf("Error creating users table: %s", err)
		return err
	}

	// Book Table
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS book (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			genre_id INTEGER REFERENCES genre(id),
			author_id INTEGER REFERENCES author(id),
			coverimage_filename VARCHAR(255),
			created_at TIMESTAMP DEFAULT NOW()
		);
	`)
	if err != nil {
		log.Printf("Error creating book table: %s", err)
		return err
	}

	// Rating Table
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS rating (
			id SERIAL PRIMARY KEY,
			book_id INTEGER REFERENCES book(id),
			user_id INTEGER REFERENCES users(id),
			amount INTEGER NOT NULL
		);
	`)
	if err != nil {
		log.Printf("Error creating rating table: %s", err)
		return err
	}

	// Quote Table
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS quote (
			id SERIAL PRIMARY KEY,
			book_id INTEGER REFERENCES book(id),
			user_id INTEGER REFERENCES users(id),
			text TEXT NOT NULL
		);
	`)
	if err != nil {
		log.Printf("Error creating quote table: %s", err)
		return err
	}

	// Review Table
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS review (
			id SERIAL PRIMARY KEY,
			user_id INTEGER REFERENCES users(id),
			book_id INTEGER REFERENCES book(id),
			text TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT NOW()
		);
	`)
	if err != nil {
		log.Printf("Error creating review table: %s", err)
		return err
	}

	// UserBook Table
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS user_book (
			id SERIAL PRIMARY KEY,
			user_id INTEGER REFERENCES users(id),
			book_id INTEGER REFERENCES book(id)
		);
	`)
	if err != nil {
		log.Printf("Error creating user_book table: %s", err)
		return err
	}

	log.Println("Database migration completed successfully")
	return nil
}
