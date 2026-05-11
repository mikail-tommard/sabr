package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"sabr/backend/services/auth/internal/domain"
)

const userRegisteredPath = "/internal/users/events/user-registered"

type userRegisteredRequest struct {
	UserID     string `json:"userId"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	OccurredAt string `json:"occurredAt"`
}

type Publisher struct {
	client  *http.Client
	baseURL string
}

func NewPublisher(client *http.Client, baseURL string) *Publisher {
	return &Publisher{
		client:  client,
		baseURL: strings.TrimRight(baseURL, "/"),
	}
}

func (p *Publisher) PublishUserRegistered(ctx context.Context, event domain.UserRegistered) error {
	body, err := json.Marshal(userRegisteredRequest{
		UserID:     event.UserID,
		Email:      event.Email,
		Name:       event.Name,
		Username:   event.Username,
		OccurredAt: event.OccurredAt.UTC().Format(time.RFC3339),
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		p.baseURL+userRegisteredPath,
		bytes.NewReader(body),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := p.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("publish user registered: users service returned status %d", resp.StatusCode)
	}

	return nil
}

func NewDefaultClient(timeout time.Duration) (*http.Client, error) {
	if timeout <= 0 {
		return nil, errors.New("http publisher timeout must be positive")
	}

	return &http.Client{Timeout: timeout}, nil
}
