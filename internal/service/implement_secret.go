package service

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/AlexTerra21/gophkeeper/internal/errs"
	"github.com/AlexTerra21/gophkeeper/internal/models"
	"github.com/AlexTerra21/gophkeeper/internal/storage"
	"github.com/AlexTerra21/gophkeeper/pb"
)

// Реализация контракта SavePassword
func (s *Service) SavePassword(ctx context.Context, req *pb.SavePasswordRequest) (*pb.Empty, error) {

	userID := GetUserIDFromMetadata(ctx)
	if userID < 0 {
		s.log.Error("Error", "Unauthenticated", "Invalid auth token")
		return nil, status.New(codes.Unauthenticated, "Invalid auth token").Err()
	}
	secretPassword := &models.PasswordSecret{
		Login:    req.Login,
		Password: req.Password,
	}

	secretData, err := secretPassword.ToBinary()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot encode secret: %v", err)
	}

	secret := storage.Secret{
		UserID:     userID,
		SecretType: int(models.SecretTypePassword),
		SecretName: req.Name,
		SecretData: secretData,
	}

	err = s.storage.SaveSecret(ctx, secret)
	if err != nil {
		s.log.Error("Error", "error adding new secret password", err)
		if errors.Is(err, errs.ErrConflict) {
			return nil, status.Errorf(codes.AlreadyExists, "SavePassword: Conflict: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "SavePassword: Internal error: %v", err)
	}

	return &pb.Empty{}, nil
}

// Реализация контракта GetPassword
func (s *Service) GetPassword(ctx context.Context, req *pb.GetSecretRequest) (*pb.PasswordResponse, error) {
	userID := GetUserIDFromMetadata(ctx)
	if userID < 0 {
		s.log.Error("Error", "Unauthenticated", "Invalid auth token")
		return nil, status.New(codes.Unauthenticated, "Invalid auth token").Err()
	}

	metadata, err := s.storage.GetSecret(ctx, userID, req.Name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "GetPassword: Internal error: %v", err)
	}

	secret := &models.PasswordSecret{}
	dec := gob.NewDecoder(bytes.NewReader(metadata.SecretData))
	err = dec.Decode(&secret)
	if err != nil {
		s.log.Error("Error", "decode", err.Error())
		return nil, status.Errorf(codes.Internal, "GetPassword: decode error: %v", err)
	}

	return &pb.PasswordResponse{
		Login:    secret.Login,
		Password: secret.Password,
	}, nil
}

// Реализация контракта SaveCard
func (s *Service) SaveCard(ctx context.Context, req *pb.SaveCardRequest) (*pb.Empty, error) {
	userID := GetUserIDFromMetadata(ctx)
	if userID < 0 {
		s.log.Error("Error", "Unauthenticated", "Invalid auth token")
		return nil, status.New(codes.Unauthenticated, "Invalid auth token").Err()
	}
	secretCard := &models.CardSecret{
		Number:     req.Number,
		HolderName: req.HolderName,
		CCV:        req.Ccv,
		Date:       req.Date,
	}

	secretData, err := secretCard.ToBinary()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot encode secret: %v", err)
	}

	secret := storage.Secret{
		UserID:     userID,
		SecretType: int(models.SecretTypeCard),
		SecretName: req.CardName,
		SecretData: secretData,
	}

	err = s.storage.SaveSecret(ctx, secret)
	if err != nil {
		s.log.Error("Error", "error adding new card", err)
		if errors.Is(err, errs.ErrConflict) {
			return nil, status.Errorf(codes.AlreadyExists, "SaveCard: Conflict: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "SaveCard: Internal error: %v", err)
	}

	return &pb.Empty{}, nil
}

// Реализация контракта GetCard
func (s *Service) GetCard(ctx context.Context, req *pb.GetSecretRequest) (*pb.CardResponse, error) {
	userID := GetUserIDFromMetadata(ctx)
	if userID < 0 {
		s.log.Error("Error", "Unauthenticated", "Invalid auth token")
		return nil, status.New(codes.Unauthenticated, "Invalid auth token").Err()
	}

	metadata, err := s.storage.GetSecret(ctx, userID, req.Name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "GetCard: Internal error: %v", err)
	}

	secret := &models.CardSecret{}
	dec := gob.NewDecoder(bytes.NewReader(metadata.SecretData))
	err = dec.Decode(&secret)
	if err != nil {
		s.log.Error("Error", "decode", err.Error())
		return nil, status.Errorf(codes.Internal, "GetCard: decode error: %v", err)
	}

	return &pb.CardResponse{
		Number:     secret.Number,
		HolderName: secret.HolderName,
		Date:       secret.Date,
		Ccv:        secret.CCV,
	}, nil
}

// Реализация контракта SaveText
func (s *Service) SaveText(ctx context.Context, req *pb.SaveTextRequest) (*pb.Empty, error) {
	userID := GetUserIDFromMetadata(ctx)
	if userID < 0 {
		s.log.Error("Error", "Unauthenticated", "Invalid auth token")
		return nil, status.New(codes.Unauthenticated, "Invalid auth token").Err()
	}

	secretText := &models.TextSecret{
		Text: req.Text,
	}

	secretData, err := secretText.ToBinary()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot encode secret: %v", err)
	}

	secret := storage.Secret{
		UserID:     userID,
		SecretType: int(models.SecretTypeText),
		SecretName: req.Name,
		SecretData: secretData,
	}

	err = s.storage.SaveSecret(ctx, secret)
	if err != nil {
		s.log.Error("Error", "error adding new note", err)
		if errors.Is(err, errs.ErrConflict) {
			return nil, status.Errorf(codes.AlreadyExists, "SaveText: Conflict: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "SaveText: Internal error: %v", err)
	}

	return &pb.Empty{}, nil
}

// Реализация контракта GetText
func (s *Service) GetText(ctx context.Context, req *pb.GetSecretRequest) (*pb.TextResponse, error) {
	userID := GetUserIDFromMetadata(ctx)
	if userID < 0 {
		s.log.Error("Error", "Unauthenticated", "Invalid auth token")
		return nil, status.New(codes.Unauthenticated, "Invalid auth token").Err()
	}

	metadata, err := s.storage.GetSecret(ctx, userID, req.Name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "GetText: Internal error: %v", err)
	}

	secret := &models.TextSecret{}
	dec := gob.NewDecoder(bytes.NewReader(metadata.SecretData))
	err = dec.Decode(&secret)
	if err != nil {
		s.log.Error("Error", "decode", err.Error())
		return nil, status.Errorf(codes.Internal, "GetText: decode error: %v", err)
	}

	return &pb.TextResponse{
		Text: secret.Text,
	}, nil
}

// Получить UserID из метаданных
func GetUserIDFromMetadata(ctx context.Context) int64 {
	userID := -1
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		values := md.Get("userID")
		// fmt.Printf("values = %v\n", values)

		if len(values) > 0 {
			userID, _ = strconv.Atoi(values[0])
		}
	}
	return (int64)(userID)
}
