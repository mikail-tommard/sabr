package middleware

import (
	"context"
	"net/http"
	"slices"
	"strings"

	"sabr/backend/pkg/json"
	jwtpkg "sabr/backend/pkg/jwt"
)

type contextKey string

const userContextKey contextKey = "auth.user"

type AuthUser struct {
	UserID string
	Role   string
}

func Auth(jwtManager *jwtpkg.Manager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := bearerToken(r.Header.Get("Authorization"))
			if tokenString == "" {
				json.Error(w, http.StatusUnauthorized, "missing bearer token")
				return
			}

			claims, err := jwtManager.Parse(tokenString)
			if err != nil {
				json.Error(w, http.StatusUnauthorized, "invalid access token")
				return
			}

			ctx := context.WithValue(r.Context(), userContextKey, AuthUser{
				UserID: claims.UserID,
				Role:   claims.Role,
			})

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RequireRoles(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, ok := UserFromContext(r.Context())
			if !ok {
				json.Error(w, http.StatusUnauthorized, "unauthorized")
				return
			}

			if !slices.Contains(roles, user.Role) {
				json.Error(w, http.StatusForbidden, "forbidden")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func UserFromContext(ctx context.Context) (AuthUser, bool) {
	user, ok := ctx.Value(userContextKey).(AuthUser)
	return user, ok
}

func bearerToken(header string) string {
	if header == "" {
		return ""
	}

	const prefix = "Bearer "
	if !strings.HasPrefix(header, prefix) {
		return ""
	}

	return strings.TrimSpace(strings.TrimPrefix(header, prefix))
}
