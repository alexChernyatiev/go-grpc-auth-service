package app

import (
	grpcapp "auth_service/internal/app/grpc"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	logger *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	// TODO init storage

	// TODO init auth

	grpcApp := grpcapp.New(logger, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
