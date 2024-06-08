package service

import (
	"log/slog"

	"github.com/AlexTerra21/gophkeeper/internal/config"
	"github.com/AlexTerra21/gophkeeper/internal/storage"
	"github.com/AlexTerra21/gophkeeper/pb"
)

type Service struct {
	pb.UnimplementedGophkeeperServer

	cfg     *config.Config
	log     *slog.Logger
	storage *storage.Storage
}

func New(cfg *config.Config,
	log *slog.Logger,
	storage *storage.Storage) (*Service, error) {
	return &Service{
		cfg:     cfg,
		log:     log,
		storage: storage,
	}, nil
}
