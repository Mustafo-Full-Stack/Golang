package handler

import (
	"net/http"

	"demo-project/internal/handler/middleware"
	v1 "demo-project/internal/handler/v1"
	"demo-project/internal/service"
)

func NewRouter(services *service.Services, cfg v1.Config) http.Handler {
	api := v1.NewHandler(services, cfg)
	requireAuth := middleware.Auth(services.Auth)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/auth/register", api.Register)
	mux.HandleFunc("POST /api/v1/auth/login", api.Login)
	// Refresh без requireAuth: в заголовке у него refresh-токен, а не access.
	mux.HandleFunc("POST /api/v1/auth/refresh", api.Refresh)
	mux.Handle("POST /api/v1/auth/logout", requireAuth(http.HandlerFunc(api.Logout)))

	mux.Handle("GET /api/v1/users/me", requireAuth(http.HandlerFunc(api.GetMe)))

	mux.Handle("GET /api/v1/sessions", requireAuth(http.HandlerFunc(api.ListSessions)))
	mux.Handle("DELETE /api/v1/sessions", requireAuth(http.HandlerFunc(api.RevokeOtherSessions)))
	mux.Handle("DELETE /api/v1/sessions/{id}", requireAuth(http.HandlerFunc(api.RevokeSession)))

	mux.Handle("POST /api/v1/files", requireAuth(http.HandlerFunc(api.UploadFile)))
	mux.Handle("GET /api/v1/files/{id}", requireAuth(http.HandlerFunc(api.GetFile)))
	mux.Handle("DELETE /api/v1/files/{id}", requireAuth(http.HandlerFunc(api.DeleteFile)))

	mux.Handle("POST /api/v1/posts", requireAuth(http.HandlerFunc(api.CreatePost)))
	mux.Handle("GET /api/v1/posts", requireAuth(http.HandlerFunc(api.ListPosts)))
	mux.Handle("GET /api/v1/posts/{id}", requireAuth(http.HandlerFunc(api.GetPost)))
	mux.Handle("PUT /api/v1/posts/{id}", requireAuth(http.HandlerFunc(api.UpdatePost)))
	mux.Handle("DELETE /api/v1/posts/{id}", requireAuth(http.HandlerFunc(api.DeletePost)))

	// Logging снаружи: он видит и ответ 500, который отдаст Recover после паники.
	return middleware.Logging(middleware.Recover(mux))
}
