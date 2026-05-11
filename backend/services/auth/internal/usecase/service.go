package usecase

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"time"

	"github.com/google/uuid"

	"sabr/backend/services/auth/internal/domain"
)

type IdentityRepository interface {
	Create(ctx context.Context, identity domain.Identity) (domain.Identity, error)
	GetByEmail(ctx context.Context, email string) (domain.Identity, error)
	GetByID(ctx context.Context, userID string) (domain.Identity, error)
}

type RefreshTokenRepository interface {
	Create(ctx context.Context, token domain.RefreshToken) error
	GetByHash(ctx context.Context, tokenHash string) (domain.RefreshToken, error)
	Revoke(ctx context.Context, tokenID string, revokedAt time.Time) error
}

type PasswordManager interface {
	Hash(password string) (string, error)
	Compare(hash string, password string) error
}

type JWTManager interface {
	Generate(userID string, role string, now time.Time) (string, time.Time, error)
}

type EventPublisher interface {
	PublishUserRegistered(ctx context.Context, event domain.UserRegistered) error
}

type RegisterInput struct {
	Name     string
	Username string
	Email    string
	Password string
}

type LoginInput struct {
	Email    string
	Password string
}

type RefreshInput struct {
	RefreshToken string
}

type AuthResult struct {
	User   IdentityOutput
	Tokens TokenOutput
}

type IdentityOutput struct {
	ID        string
	Email     string
	Role      string
	CreatedAt time.Time
}

type TokenOutput struct {
	AccessToken           string
	AccessTokenExpiresAt  time.Time
	RefreshToken          string
	RefreshTokenExpiresAt time.Time
}

type Service struct {
	identities      IdentityRepository
	refreshTokens   RefreshTokenRepository
	passwords       PasswordManager
	jwt             JWTManager
	events          EventPublisher
	refreshTokenTTL time.Duration
	now             func() time.Time
}

func NewService(
	identities IdentityRepository,
	refreshTokens RefreshTokenRepository,
	passwords PasswordManager,
	jwt JWTManager,
	events EventPublisher,
	refreshTokenTTL time.Duration,
	now func() time.Time,
) *Service {
	return &Service{
		identities:      identities,
		refreshTokens:   refreshTokens,
		passwords:       passwords,
		jwt:             jwt,
		events:          events,
		refreshTokenTTL: refreshTokenTTL,
		now:             now,
	}
}

func (s *Service) Register(ctx context.Context, input RegisterInput) (AuthResult, error) {
	name := strings.TrimSpace(input.Name)
	username := normalizeUsername(input.Username)
	email := normalizeEmail(input.Email)
	password := input.Password

	if _, err := s.identities.GetByEmail(ctx, email); err == nil {
		return AuthResult{}, domain.ErrEmailTaken
	}

	passwordHash, err := s.passwords.Hash(password)
	if err != nil {
		return AuthResult{}, err
	}

	now := s.now()
	identity, err := s.identities.Create(ctx, domain.Identity{
		ID:           uuid.NewString(),
		Email:        email,
		PasswordHash: passwordHash,
		Role:         domain.DefaultRole,
		CreatedAt:    now,
		UpdatedAt:    now,
	})
	if err != nil {
		return AuthResult{}, err
	}

	if err = s.events.PublishUserRegistered(ctx, domain.UserRegistered{
		UserID:     identity.ID,
		Email:      identity.Email,
		Name:       name,
		Username:   username,
		OccurredAt: now,
	}); err != nil {
		return AuthResult{}, err
	}

	return s.issueTokens(ctx, identity)
}

func (s *Service) Login(ctx context.Context, input LoginInput) (AuthResult, error) {
	email := normalizeEmail(input.Email)

	identity, err := s.identities.GetByEmail(ctx, email)
	if err != nil {
		return AuthResult{}, domain.ErrInvalidCredentials
	}

	if err = s.passwords.Compare(identity.PasswordHash, input.Password); err != nil {
		return AuthResult{}, domain.ErrInvalidCredentials
	}

	return s.issueTokens(ctx, identity)
}

func (s *Service) Refresh(ctx context.Context, input RefreshInput) (AuthResult, error) {
	tokenHash := hashToken(input.RefreshToken)

	stored, err := s.refreshTokens.GetByHash(ctx, tokenHash)
	if err != nil {
		return AuthResult{}, domain.ErrInvalidRefreshToken
	}

	now := s.now()
	if stored.RevokedAt != nil || now.After(stored.ExpiresAt) {
		return AuthResult{}, domain.ErrInvalidRefreshToken
	}

	if err = s.refreshTokens.Revoke(ctx, stored.ID, now); err != nil {
		return AuthResult{}, err
	}

	identity, err := s.identities.GetByID(ctx, stored.UserID)
	if err != nil {
		return AuthResult{}, domain.ErrUserNotFound
	}

	return s.issueTokens(ctx, identity)
}

func (s *Service) Me(ctx context.Context, userID string) (IdentityOutput, error) {
	identity, err := s.identities.GetByID(ctx, userID)
	if err != nil {
		return IdentityOutput{}, domain.ErrUserNotFound
	}

	return newIdentityOutput(identity), nil
}

func (s *Service) issueTokens(ctx context.Context, identity domain.Identity) (AuthResult, error) {
	now := s.now()

	accessToken, accessExpiresAt, err := s.jwt.Generate(identity.ID, identity.Role, now)
	if err != nil {
		return AuthResult{}, err
	}

	refreshToken, err := newRefreshToken()
	if err != nil {
		return AuthResult{}, err
	}

	refreshExpiresAt := now.Add(s.refreshTokenTTL)
	refreshRecord := domain.RefreshToken{
		ID:        uuid.NewString(),
		UserID:    identity.ID,
		TokenHash: hashToken(refreshToken),
		ExpiresAt: refreshExpiresAt,
		CreatedAt: now,
	}
	if err = s.refreshTokens.Create(ctx, refreshRecord); err != nil {
		return AuthResult{}, err
	}

	return AuthResult{
		User: newIdentityOutput(identity),
		Tokens: TokenOutput{
			AccessToken:           accessToken,
			AccessTokenExpiresAt:  accessExpiresAt,
			RefreshToken:          refreshToken,
			RefreshTokenExpiresAt: refreshExpiresAt,
		},
	}, nil
}

func newIdentityOutput(identity domain.Identity) IdentityOutput {
	return IdentityOutput{
		ID:        identity.ID,
		Email:     identity.Email,
		Role:      identity.Role,
		CreatedAt: identity.CreatedAt,
	}
}

func normalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

func normalizeUsername(username string) string {
	return strings.ToLower(strings.TrimSpace(username))
}

func newRefreshToken() (string, error) {
	randomBytes := make([]byte, 32)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(randomBytes), nil
}

func hashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return base64.RawURLEncoding.EncodeToString(sum[:])
}
