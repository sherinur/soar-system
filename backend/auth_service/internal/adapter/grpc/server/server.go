package grpcserver

import (
	"context"
	"fmt"
	"net"

	"github.com/sherinur/soar-system/backend/auth_service/config"
	"go.uber.org/zap"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const serverIPAddress = "0.0.0.0:%d"

type API struct {
	server *grpc.Server
	cfg    config.Config
	addr   string

	log *zap.Logger
}

func New(cfg config.Config, log *zap.Logger) (*API, error) {
	return &API{
		cfg:  cfg,
		addr: fmt.Sprintf(serverIPAddress, cfg.Server.GRPCServer.Port),
		log:  log,
	}, nil
}

func (a *API) Run(ctx context.Context) error {
	return a.run(ctx)
}

func (a *API) run(ctx context.Context) error {
	a.server = grpc.NewServer(a.setOptions(ctx)...)

	// feedbacksvc.RegisterFeedbackServiceServer(a.server, a.log))

	reflection.Register(a.server)

	a.log.Info("Service started",
		zap.String("protocol", "gRPC"),
		zap.String("address", a.addr),
	)

	listener, err := net.Listen("tcp", a.addr)
	if err != nil {
		a.log.Error("Failed to create listener", zap.Error(err))
		return fmt.Errorf("failed to create listener: %w", err)
	}

	err = a.server.Serve(listener)
	if err != nil {
		a.log.Error("gRPC server failed", zap.Error(err))
		return fmt.Errorf("failed to serve grpc: %w", err)
	}

	return nil
}
