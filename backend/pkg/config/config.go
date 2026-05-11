package config

import (
	"errors"
	"os"
	"time"
)

const (
	defaultHTTPAddr        = ":8080"
	defaultAccessTokenTTL  = 15 * time.Minute
	defaultRefreshTokenTTL = 30 * 24 * time.Hour
)

type Config struct {
	HTTPAddr        string
	DatabaseDSN     string
	JWTSecret       string
	UsersServiceURL string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

func Load() (Config, error) {
	cfg := Config{
		HTTPAddr:        getEnv("APP_ADDR", defaultHTTPAddr),
		DatabaseDSN:     os.Getenv("DB_DSN"),
		JWTSecret:       os.Getenv("JWT_SECRET"),
		UsersServiceURL: os.Getenv("USERS_SERVICE_URL"),
		AccessTokenTTL:  getDurationEnv("ACCESS_TOKEN_TTL", defaultAccessTokenTTL),
		RefreshTokenTTL: getDurationEnv("REFRESH_TOKEN_TTL", defaultRefreshTokenTTL),
	}

	if cfg.DatabaseDSN == "" {
		return Config{}, errors.New("DB_DSN is required")
	}
	if cfg.JWTSecret == "" {
		return Config{}, errors.New("JWT_SECRET is required")
	}
	if cfg.UsersServiceURL == "" {
		return Config{}, errors.New("USERS_SERVICE_URL is required")
	}

	return cfg, nil
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

func getDurationEnv(key string, fallback time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	parsed, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}

	return parsed
}
