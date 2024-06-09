package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"log/slog"
	"net"
	"strings"

	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	"github.com/AlexTerra21/gophkeeper/internal/auth"
	"github.com/AlexTerra21/gophkeeper/internal/config"
	"github.com/AlexTerra21/gophkeeper/internal/storage"
)

func logInterceptor(logger *slog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		logger.Info("gRPC request",
			"method", info.FullMethod,
			"request", req,
		)
		return handler(ctx, req)
	}
}

func authInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	if strings.Contains(info.FullMethod, "Register") || strings.Contains(info.FullMethod, "Login") {
		return handler(ctx, req)
	}
	var token string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		values := md.Get("authorization")
		if len(values) > 0 {
			token = values[0]
		}
	}
	userID := auth.GetUserID(token)
	newCtx := ctx
	if userID > 0 {
		incomCtx, _ := metadata.FromIncomingContext(ctx)
		newMD := metadata.Pairs("userID", fmt.Sprintf("%d", userID))
		newCtx = metadata.NewIncomingContext(ctx, metadata.Join(incomCtx, newMD))
	}
	return handler(newCtx, req)
}

func NewGRPCServer(lc fx.Lifecycle, config *config.Config, logger *slog.Logger, storage *storage.Storage) (*grpc.Server, error) {

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		return nil, fmt.Errorf("cannot load TLS credentials: %v", err)
	}

	srv := grpc.NewServer(
		grpc.Creds(tlsCredentials),
		grpc.ChainUnaryInterceptor(logInterceptor(logger), authInterceptor),
	)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Starting gRPC server")

			addr := "localhost:3200"
			ln, err := net.Listen("tcp", addr)
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
			storage.Close()
			srv.GracefulStop()

			return nil
		},
	})

	return srv, nil
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair("cert/server-cert.pem", "cert/server-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}
