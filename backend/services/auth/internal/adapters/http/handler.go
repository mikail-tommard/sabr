package http

import (
	"errors"
	"net/http"
	"strings"
	"time"

	jsonpkg "sabr/backend/pkg/json"
	"sabr/backend/pkg/middleware"
	"sabr/backend/services/auth/internal/domain"
	"sabr/backend/services/auth/internal/usecase"
)

type registerRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type refreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type authResponse struct {
	User   userResponse   `json:"user"`
	Tokens tokensResponse `json:"tokens"`
}

type userResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"createdAt"`
}

type tokensResponse struct {
	AccessToken           string `json:"accessToken"`
	AccessTokenExpiresAt  string `json:"accessTokenExpiresAt"`
	RefreshToken          string `json:"refreshToken"`
	RefreshTokenExpiresAt string `json:"refreshTokenExpiresAt"`
}

type Handler struct {
	service *usecase.Service
}

func NewHandler(service *usecase.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var request registerRequest
	if err := jsonpkg.Read(r, &request); err != nil {
		jsonpkg.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	input := usecase.RegisterInput{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	if err := validateRegister(input); err != nil {
		jsonpkg.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.service.Register(r.Context(), input)
	if err != nil {
		handleError(w, err)
		return
	}

	jsonpkg.Write(w, http.StatusCreated, newAuthResponse(result))
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var request loginRequest
	if err := jsonpkg.Read(r, &request); err != nil {
		jsonpkg.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	input := usecase.LoginInput{
		Email:    request.Email,
		Password: request.Password,
	}

	if err := validateLogin(input); err != nil {
		jsonpkg.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.service.Login(r.Context(), input)
	if err != nil {
		handleError(w, err)
		return
	}

	jsonpkg.Write(w, http.StatusOK, newAuthResponse(result))
}

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	var request refreshRequest
	if err := jsonpkg.Read(r, &request); err != nil {
		jsonpkg.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	input := usecase.RefreshInput{
		RefreshToken: request.RefreshToken,
	}

	if strings.TrimSpace(input.RefreshToken) == "" {
		jsonpkg.Error(w, http.StatusBadRequest, "refreshToken is required")
		return
	}

	result, err := h.service.Refresh(r.Context(), input)
	if err != nil {
		handleError(w, err)
		return
	}

	jsonpkg.Write(w, http.StatusOK, newAuthResponse(result))
}

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	authUser, ok := middleware.UserFromContext(r.Context())
	if !ok {
		jsonpkg.Error(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	user, err := h.service.Me(r.Context(), authUser.UserID)
	if err != nil {
		handleError(w, err)
		return
	}

	jsonpkg.Write(w, http.StatusOK, map[string]userResponse{"user": newUserResponse(user)})
}

func validateRegister(input usecase.RegisterInput) error {
	switch {
	case len(strings.TrimSpace(input.Name)) < 2:
		return errors.New("name must contain at least 2 characters")
	case len(strings.TrimSpace(input.Username)) < 3:
		return errors.New("username must contain at least 3 characters")
	case !strings.Contains(strings.TrimSpace(input.Email), "@"):
		return errors.New("email is invalid")
	case len(input.Password) < 8:
		return errors.New("password must contain at least 8 characters")
	default:
		return nil
	}
}

func validateLogin(input usecase.LoginInput) error {
	switch {
	case !strings.Contains(strings.TrimSpace(input.Email), "@"):
		return errors.New("email is invalid")
	case len(input.Password) == 0:
		return errors.New("password is required")
	default:
		return nil
	}
}

func handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrEmailTaken), errors.Is(err, domain.ErrUsernameTaken):
		jsonpkg.Error(w, http.StatusConflict, err.Error())
	case errors.Is(err, domain.ErrInvalidCredentials), errors.Is(err, domain.ErrInvalidRefreshToken):
		jsonpkg.Error(w, http.StatusUnauthorized, err.Error())
	case errors.Is(err, domain.ErrUserNotFound):
		jsonpkg.Error(w, http.StatusNotFound, err.Error())
	default:
		jsonpkg.Error(w, http.StatusInternalServerError, "internal server error")
	}
}

func newAuthResponse(result usecase.AuthResult) authResponse {
	return authResponse{
		User:   newUserResponse(result.User),
		Tokens: newTokensResponse(result.Tokens),
	}
}

func newUserResponse(user usecase.UserOutput) userResponse {
	return userResponse{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.UTC().Format(time.RFC3339),
	}
}

func newTokensResponse(tokens usecase.TokenOutput) tokensResponse {
	return tokensResponse{
		AccessToken:           tokens.AccessToken,
		AccessTokenExpiresAt:  tokens.AccessTokenExpiresAt.UTC().Format(time.RFC3339),
		RefreshToken:          tokens.RefreshToken,
		RefreshTokenExpiresAt: tokens.RefreshTokenExpiresAt.UTC().Format(time.RFC3339),
	}
}
