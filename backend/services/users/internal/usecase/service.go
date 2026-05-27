package usecase

import (
	"context"
	"sabr/backend/services/users/internal/domain"
	"strings"
	"time"
)

type UserRepository interface {
	Create(ctx context.Context, user domain.User) error
	GetByID(ctx context.Context, id string) (domain.User, error)
	GetByEmail(ctx context.Context, email string) (domain.User, error)
	GetByUsername(ctx context.Context, username string) (domain.User, error)
	Update(ctx context.Context, user domain.User) error
}

type UserService struct {
	repo UserRepository
	now  func() time.Time
}

type CreateProfileFromRegistrationInput struct {
	ID       string
	Email    string
	Name     string
	Username string
}

type UpdateProfileInput struct {
	ID        string
	Name      string
	Username  string
	AvatarURL string
	Bio       string
	CampusID  *string
	Company   *string
}

func NewService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
		now:  time.Now,
	}
}

func (s *UserService) CreateProfileFromRegistration(ctx context.Context, input CreateProfileFromRegistrationInput) error {
	if _, err := s.repo.GetByID(ctx, input.ID); err == nil {
		return domain.ErrUserAlreadyExists
	}

	username := strings.ToLower(strings.TrimSpace(input.Username))
	if _, err := s.repo.GetByUsername(ctx, username); err == nil {
		return domain.ErrUsernameTaken
	}

	now := s.now()

	user, err := domain.NewUser(domain.NewUserParams{
		ID:        input.ID,
		Name:      input.Name,
		Email:     input.Email,
		Username:  input.Username,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return err
	}

	return s.repo.Create(ctx, *user)
}

func (s *UserService) GetByID(ctx context.Context, id string) (domain.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	return s.repo.GetByEmail(ctx, email)
}

func (s *UserService) GetByUsername(ctx context.Context, username string) (domain.User, error) {
	return s.repo.GetByUsername(ctx, username)
}

func (s *UserService) UpdateProfile(ctx context.Context, input UpdateProfileInput) error {
	user, err := s.repo.GetByID(ctx, input.ID)
	if err != nil {
		return err
	}

	nextUsername := strings.ToLower(strings.TrimSpace(input.Username))
	if user.Username != nextUsername {
		if _, err := s.repo.GetByUsername(ctx, nextUsername); err == nil {
			return domain.ErrUsernameTaken
		}
	}

	err = user.UpdateProfile(domain.UpdateProfileParams{
		Name:      input.Name,
		Username:  input.Username,
		AvatarURL: input.AvatarURL,
		Bio:       input.Bio,
		CampusID:  input.CampusID,
		Company:   input.Company,
		UpdatedAt: s.now(),
	})
	if err != nil {
		return err
	}

	return s.repo.Update(ctx, user)
}
