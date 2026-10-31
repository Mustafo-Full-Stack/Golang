package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"demo-project/internal/handler/response"
)

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if p := recover(); p != nil {
				err := fmt.Errorf("panic: %v\n%s", p, debug.Stack())
				response.Error(w, r, http.StatusInternalServerError, "internal server error", err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
