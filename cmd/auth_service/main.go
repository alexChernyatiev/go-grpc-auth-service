package main

import (
	"auth_service/internal/app"
	"auth_service/internal/config"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	// config
	cfg := config.MustLoad()

	// logger
	logger := setupLogger(cfg.Env)

	// init app
	application := app.New(
		logger,
		cfg.Grpc.Port,
		cfg.StoragePath,
		cfg.TokenTTL,
	)

	// run app
	application.GRPCServer.MustRun()
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		panic("unknown env: " + env)
	}

	return logger
}
