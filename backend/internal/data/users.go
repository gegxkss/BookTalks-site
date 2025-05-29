package data

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gegxkss/BookTalks-site/backend/internal/domain"
)

// GetUser получает пользователя по ID
func GetUser(userID int) (domain.User, error) {
	query := `
		SELECT
			id, nickname, first_name, last_name, sex, birth_date, email, password, created_at, profile_image_file_name
		FROM
			users
		WHERE
			id = $1
	`
	row := DB.QueryRow(query, userID)
	var user domain.User
	err := row.Scan(&user.ID, &user.Nickname, &user.FirstName, &user.LastName, &user.Sex, &user.BirthDate, &user.Email, &user.Password, &user.CreatedAt, &user.ProfileImageFileName)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

// RegisterUser регистрирует нового пользователя
func RegisterUser(req domain.RegisterRequest) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("failed to hash password: %w", err)
	}

	query := `
        INSERT INTO users (nickname, first_name, last_name, sex, birth_date, email, password, profile_image_file_name)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id
    `

	var userID int
	log.Printf("Saving user with ProfileImageFileName: %s", req.ProfileImageFileName)
	var profileImageFileName sql.NullString
	if req.ProfileImageFileName != "" {
		profileImageFileName = sql.NullString{String: req.ProfileImageFileName, Valid: true}
	} else {
		profileImageFileName = sql.NullString{Valid: false}
	}
	err = DB.QueryRow(
		query,
		req.Nickname,
		req.FirstName,
		req.LastName,
		req.Sex,
		req.BirthDate,
		req.Email,
		hashedPassword,
		profileImageFileName,
	).Scan(&userID)

	if err != nil {
		return 0, fmt.Errorf("failed to insert user: %w", err)
	}

	return userID, nil
}

// AuthenticateUser проверяет учетные данные пользователя
func AuthenticateUser(email, password string) (int, error) {
	var userID int
	var hashedPassword string

	query := `SELECT id, password FROM users WHERE email = $1`
	err := DB.QueryRow(query, email).Scan(&userID, &hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("invalid email or password")
		}
		return 0, fmt.Errorf("failed to get user: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return 0, fmt.Errorf("invalid email or password")
	}

	return userID, nil
}

// SetUserCookie sets a user ID cookie
func SetUserCookie(w http.ResponseWriter, userID int) {
	cookie := &http.Cookie{
		Name:     "userId",
		Value:    strconv.Itoa(userID),
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour), // Пример: 24 часа
		HttpOnly: true,
		Secure:   true, //  Устанавливайте Secure: true только если используете HTTPS
	}
	http.SetCookie(w, cookie)
}
