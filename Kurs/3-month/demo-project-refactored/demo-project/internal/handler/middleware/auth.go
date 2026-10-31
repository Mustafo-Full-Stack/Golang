package middleware

import (
	"context"
	"net/http"
	"strings"

	"demo-project/internal/handler/response"
	"demo-project/internal/models"
	"demo-project/internal/service"
)

type currentUserKey struct{}

// Auth пропускает запрос дальше, только если в нём действующий access-токен,
// и кладёт текущего пользователя в контекст.
func Auth(auth service.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, err := auth.Authenticate(r.Context(), BearerToken(r))
			if err != nil {
				response.ServiceError(w, r, err)
				return
			}

			ctx := context.WithValue(r.Context(), currentUserKey{}, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// BearerToken достаёт токен из заголовка "Authorization: Bearer <token>".
// Нет заголовка — пустая строка: такой токен не пройдёт проверку в сервисе.
func BearerToken(r *http.Request) string {
	token, found := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")
	if !found {
		return ""
	}
	return strings.TrimSpace(token)
}

// CurrentUser вызывать только в ручках за middleware.Auth: там пользователь точно есть.
func CurrentUser(ctx context.Context) models.CurrentUser {
	user, _ := ctx.Value(currentUserKey{}).(models.CurrentUser)
	return user
}
