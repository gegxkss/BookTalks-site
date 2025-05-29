package cookies

import (
	"fmt"
	"net/http"
	"strconv"
)

// RequireAuth проверяет наличие аутентификации пользователя.
func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := GetUserIDFromCookie(r) // Извлекаем и проверяем ID из cookie.
		if err != nil {
			// Если cookie отсутствует или невалидный, перенаправляем на страницу регистрации
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Пользователь аутентифицирован, просто вызываем следующий обработчик
		next.ServeHTTP(w, r)
	}
}

// GetUserIDFromCookie получает и проверяет ID пользователя из cookie.
func GetUserIDFromCookie(r *http.Request) (int, error) {
	cookie, err := r.Cookie("userId")
	if err != nil {
		return 0, err // Cookie отсутствует.
	}

	userIDStr := cookie.Value
	_, err = strconv.Atoi(userIDStr) // Просто проверяем, что ID - число.  Не используем ID.
	if err != nil {
		return 0, fmt.Errorf("invalid user ID in cookie")
	}
	return 0, nil // Возвращаем 0, так как ID не нужен.
}
