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

type UserRepository interface {
	Create(ctx context.Context, user domain.User) (domain.User, error)
	GetByEmail(ctx context.Context, email string) (domain.User, error)
	GetByUsername(ctx context.Context, username string) (domain.User, error)
	GetByID(ctx context.Context, userID string) (domain.User, error)
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
	User   UserOutput
	Tokens TokenOutput
}

type UserOutput struct {
	ID        string
	Name      string
	Username  string
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
	users           UserRepository
	refreshTokens   RefreshTokenRepository
	passwords       PasswordManager
	jwt             JWTManager
	refreshTokenTTL time.Duration
	now             func() time.Time
}

func NewService(
	users UserRepository,
	refreshTokens RefreshTokenRepository,
	passwords PasswordManager,
	jwt JWTManager,
	refreshTokenTTL time.Duration,
	now func() time.Time,
) *Service {
	return &Service{
		users:           users,
		refreshTokens:   refreshTokens,
		passwords:       passwords,
		jwt:             jwt,
		refreshTokenTTL: refreshTokenTTL,
		now:             now,
	}
}

func (s *Service) Register(ctx context.Context, input RegisterInput) (AuthResult, error) {
	name := strings.TrimSpace(input.Name)
	username := normalizeUsername(input.Username)
	email := normalizeEmail(input.Email)
	password := input.Password

	if _, err := s.users.GetByEmail(ctx, email); err == nil {
		return AuthResult{}, domain.ErrEmailTaken
	}

	if _, err := s.users.GetByUsername(ctx, username); err == nil {
		return AuthResult{}, domain.ErrUsernameTaken
	}

	passwordHash, err := s.passwords.Hash(password)
	if err != nil {
		return AuthResult{}, err
	}

	now := s.now()
	user, err := s.users.Create(ctx, domain.User{
		ID:           uuid.NewString(),
		Name:         name,
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
		Role:         domain.DefaultRole,
		CreatedAt:    now,
		UpdatedAt:    now,
	})
	if err != nil {
		return AuthResult{}, err
	}

	return s.issueTokens(ctx, user)
}

func (s *Service) Login(ctx context.Context, input LoginInput) (AuthResult, error) {
	email := normalizeEmail(input.Email)

	user, err := s.users.GetByEmail(ctx, email)
	if err != nil {
		return AuthResult{}, domain.ErrInvalidCredentials
	}

	if err = s.passwords.Compare(user.PasswordHash, input.Password); err != nil {
		return AuthResult{}, domain.ErrInvalidCredentials
	}

	return s.issueTokens(ctx, user)
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

	user, err := s.users.GetByID(ctx, stored.UserID)
	if err != nil {
		return AuthResult{}, domain.ErrUserNotFound
	}

	return s.issueTokens(ctx, user)
}

func (s *Service) Me(ctx context.Context, userID string) (UserOutput, error) {
	user, err := s.users.GetByID(ctx, userID)
	if err != nil {
		return UserOutput{}, domain.ErrUserNotFound
	}

	return newUserOutput(user), nil
}

func (s *Service) issueTokens(ctx context.Context, user domain.User) (AuthResult, error) {
	now := s.now()

	accessToken, accessExpiresAt, err := s.jwt.Generate(user.ID, user.Role, now)
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
		UserID:    user.ID,
		TokenHash: hashToken(refreshToken),
		ExpiresAt: refreshExpiresAt,
		CreatedAt: now,
	}
	if err = s.refreshTokens.Create(ctx, refreshRecord); err != nil {
		return AuthResult{}, err
	}

	return AuthResult{
		User: newUserOutput(user),
		Tokens: TokenOutput{
			AccessToken:           accessToken,
			AccessTokenExpiresAt:  accessExpiresAt,
			RefreshToken:          refreshToken,
			RefreshTokenExpiresAt: refreshExpiresAt,
		},
	}, nil
}

func newUserOutput(user domain.User) UserOutput {
	return UserOutput{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
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
