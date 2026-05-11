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
	mux.HandleFunc("POST /register", a.Handler.Register)
	mux.HandleFunc("POST /login", a.Handler.Login)
	mux.HandleFunc("POST /refresh", a.Handler.Refresh)
	mux.Handle("GET /me", middleware.Auth(a.JWTManager)(http.HandlerFunc(a.Handler.Me)))

	return mux
}
