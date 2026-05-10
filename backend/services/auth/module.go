package auth

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	authhttp "sabr/backend/services/auth/internal/adapters/http"
	authpostgres "sabr/backend/services/auth/internal/adapters/postgres"
	"sabr/backend/services/auth/internal/usecase"

	jwtpkg "sabr/backend/pkg/jwt"
	"sabr/backend/pkg/security"
)

type Module struct {
	API        *authhttp.API
	JWTManager *jwtpkg.Manager
}

func NewModule(
	db *pgxpool.Pool,
	jwtSecret string,
	accessTokenTTL time.Duration,
	refreshTokenTTL time.Duration,
	now func() time.Time,
) Module {
	jwtManager := jwtpkg.NewManager(jwtSecret, accessTokenTTL)
	passwordManager := security.NewPasswordManager()

	userRepo := authpostgres.NewUserRepository(db)
	refreshRepo := authpostgres.NewRefreshTokenRepository(db)

	service := usecase.NewService(
		userRepo,
		refreshRepo,
		passwordManager,
		jwtManager,
		refreshTokenTTL,
		now,
	)

	handler := authhttp.NewHandler(service)

	return Module{
		API:        authhttp.NewAPI(handler, jwtManager),
		JWTManager: jwtManager,
	}
}
