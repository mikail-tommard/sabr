package auth

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	autheventshttp "sabr/backend/services/auth/internal/adapters/events/http"
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
	usersServiceURL string,
	accessTokenTTL time.Duration,
	refreshTokenTTL time.Duration,
	now func() time.Time,
) (Module, error) {
	jwtManager := jwtpkg.NewManager(jwtSecret, accessTokenTTL)
	passwordManager := security.NewPasswordManager()
	httpClient, err := autheventshttp.NewDefaultClient(5 * time.Second)
	if err != nil {
		return Module{}, err
	}
	eventPublisher := autheventshttp.NewPublisher(httpClient, usersServiceURL)

	identityRepo := authpostgres.NewIdentityRepository(db)
	refreshRepo := authpostgres.NewRefreshTokenRepository(db)

	service := usecase.NewService(
		identityRepo,
		refreshRepo,
		passwordManager,
		jwtManager,
		eventPublisher,
		refreshTokenTTL,
		now,
	)

	handler := authhttp.NewHandler(service)

	return Module{
		API:        authhttp.NewAPI(handler, jwtManager),
		JWTManager: jwtManager,
	}, nil
}
