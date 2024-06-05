package server

import (
	"context"
	"log/slog"
	"net"

	"go.uber.org/fx"
	"google.golang.org/grpc"

	"github.com/AlexTerra21/gophkeeper/internal/config"
	"github.com/AlexTerra21/gophkeeper/pb"
)

// GRPCServer
type GRPCServer struct {
	pb.UnimplementedGophkeeperServer
	server *grpc.Server
	config *config.Config
}

// Конструктор GRPCServer
// func NewGRPCServer(config *config.Config) (*GRPCServer, error) {
// 	s := grpc.NewServer(grpc.ChainUnaryInterceptor(logInterceptor, authInterceptor))
// 	return &GRPCServer{
// 		server: s,
// 		config: config,
// 	}, nil
// }

func NewGRPCServer(lc fx.Lifecycle, config *config.Config, logger *slog.Logger) *grpc.Server {
	srv := grpc.NewServer()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Starting gRPC server")

			ln, err := net.Listen("tcp", ":9000")
			if err != nil {
				return err
			}

			go func() {
				if err := srv.Serve(ln); err != nil {
					logger.Error("Failed to Serve gRPC", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Gracefully stopping gRPC server")

			srv.GracefulStop()

			return nil
		},
	})

	return srv
}
