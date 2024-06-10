package service

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/AlexTerra21/gophkeeper/internal/auth"
	"github.com/AlexTerra21/gophkeeper/internal/errs"
	"github.com/AlexTerra21/gophkeeper/internal/storage"
	"github.com/AlexTerra21/gophkeeper/pb"
)

func (s *Service) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.Empty, error) {

	userID, err := s.storage.AddUser(ctx, &storage.User{
		Name:     req.Username,
		Password: req.Password,
	})
	if err != nil {
		s.log.Debug("Error", "error adding new user", err)
		if errors.Is(err, errs.ErrConflict) {
			return &pb.Empty{}, status.New(codes.AlreadyExists, "Error add user: Conflict. "+err.Error()).Err()
		}
		return &pb.Empty{}, status.New(codes.Internal, "Error add user: Internal. "+err.Error()).Err()
	}
	token, err := auth.BuildJWTString(userID)
	if err == nil {
		header := metadata.Pairs("Authorization", token)
		_ = grpc.SendHeader(ctx, header)
	}

	return &pb.Empty{}, nil
}

func (s *Service) Login(ctx context.Context, req *pb.LoginRequest) (*pb.Empty, error) {
	s.log.Debug("Login", "req", req)

	userID, err := s.storage.CheckLoginPassword(ctx, &storage.User{
		Name:     req.Username,
		Password: req.Password,
	})
	if err != nil {
		s.log.Debug("Error", "database error", err)
		return &pb.Empty{}, status.New(codes.Internal, "Database error: Internal. "+err.Error()).Err()
	}
	if userID < 0 {
		s.log.Debug("Error", "Unauthenticated", "Invalid login or password")
		return &pb.Empty{}, status.New(codes.Unauthenticated, "Invalid login or password").Err()
	}
	token, err := auth.BuildJWTString(userID)
	if err == nil {
		header := metadata.Pairs("Authorization", token)
		_ = grpc.SendHeader(ctx, header)
	}
	return &pb.Empty{}, nil
}
