package http

import (
	"net/http"

	"sabr/backend/pkg/jwt"
	"sabr/backend/pkg/middleware"
)

type API struct {
	Handler    *Handler
	JWTManager *jwt.Manager
}

func NewAPI(handler *Handler, jwtManager *jwt.Manager) *API {
	return &API{
		Handler:    handler,
		JWTManager: jwtManager,
	}
}

func (a *API) Router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /auth/register", a.Handler.Register)
	mux.HandleFunc("POST /auth/login", a.Handler.Login)
	mux.HandleFunc("POST /auth/refresh", a.Handler.Refresh)
	mux.Handle("GET /auth/me", middleware.Auth(a.JWTManager)(http.HandlerFunc(a.Handler.Me)))

	return mux
}
