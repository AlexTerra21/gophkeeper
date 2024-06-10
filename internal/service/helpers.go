package service

import (
	"github.com/AlexTerra21/gophkeeper/internal/config"
	"github.com/AlexTerra21/gophkeeper/internal/logger"
	"github.com/AlexTerra21/gophkeeper/internal/storage"
)

// Подготовка тестовой среды
func PrepareTestEnv() (*Service, error) {
	config, _ := config.NewConfig()
	config.DBConnectString = "postgresql://gopherkeeper:gopherkeeper@localhost/gopherkeeper_test?sslmode=disable"
	log, _ := logger.NewLogger(config.LogLevel)
	storage, err := storage.NewStorage(config, log)
	if err != nil {
		return nil, err
	}
	err = ClearTestDB(storage)
	if err != nil {
		return nil, err
	}
	service := &Service{
		cfg:     config,
		log:     log,
		storage: storage,
	}
	return service, nil
}

// Очистка тестовых таблиц
func ClearTestDB(storage *storage.Storage) error {
	db := storage.GetDB()
	_, err := db.Exec("TRUNCATE secrets")
	if err != nil {
		return err
	}
	_, err = db.Exec("TRUNCATE users")
	if err != nil {
		return err
	}
	return nil
}
