package app

import (
	"context"

	"github.com/sherinur/soar-system/backend/auth_service/config"
	grpcserver "github.com/sherinur/soar-system/backend/auth_service/internal/adapter/grpc/server"
	"github.com/sherinur/soar-system/backend/auth_service/pkg/postgrescon"
	"go.uber.org/zap"
)

const serviceName = "auth-service"

type App struct {
	cfg *config.Config
	log *zap.Logger

	grpcServer *grpcserver.API
	// httpServer *server.API

	telemetry *Telemetry
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	// logger
	logger, err := NewLogger(cfg)
	if err != nil {
		return nil, err
	}

	// Postgres connection
	logger.Info("connecting to postgresql", zap.String("database", cfg.Postgres.DBName))
	_, err = postgrescon.Connect(&cfg.Postgres)
	if err != nil {
		logger.Fatal("error occured when connecting to postgresql", zap.Error(err))
	}

	// controllers ...
	// grpcServer, err := grpcserver.New(*cfg, logger, feedbackUsecase)
	// if err != nil {
	// 	return nil, err
	// }

	// telemetry
	telemetry, err := InitTelemetry(ctx, cfg.Telemetry, logger)
	if err != nil {
		return nil, err
	}

	app := &App{
		log:       logger,
		telemetry: telemetry,
	}

	return app, nil
}

func (a *App) Run() error {
	a.log.Info("Starting the service")
	return a.grpcServer.Run(context.Background())
}

func (a *App) Stop() error {
	return nil
}
