package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"sabr/backend/pkg/config"
	"sabr/backend/services/auth"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("load config", "error", err)
		os.Exit(1)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	dbpool, err := pgxpool.New(ctx, cfg.DatabaseDSN)
	if err != nil {
		logger.Error("connect db", "error", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	if err = dbpool.Ping(ctx); err != nil {
		logger.Error("ping db", "error", err)
		os.Exit(1)
	}

	authModule, err := auth.NewModule(
		dbpool,
		cfg.JWTSecret,
		cfg.UsersServiceURL,
		cfg.AccessTokenTTL,
		cfg.RefreshTokenTTL,
		time.Now,
	)
	if err != nil {
		logger.Error("build auth module", "error", err)
		os.Exit(1)
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
	router.Handle("/auth/", http.StripPrefix("/auth", authModule.API.Router()))

	server := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		logger.Info("http server started", "addr", cfg.HTTPAddr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("listen and serve", "error", err)
			stop()
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error("shutdown", "error", err)
		os.Exit(1)
	}
}
