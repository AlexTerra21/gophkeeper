package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	"github.com/AlexTerra21/gophkeeper/internal/config"
	"github.com/AlexTerra21/gophkeeper/internal/logger"
	"github.com/AlexTerra21/gophkeeper/internal/server"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"google.golang.org/grpc"
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

// go build -o cmd/gophkeeper/gophkeeper.exe -ldflags "-X main.buildVersion=v1.20.0 -X 'main.buildDate=$(date +'%Y/%m/%d %H:%M:%S')' -X 'main.buildCommit=$(git log -1 | grep commit)'" cmd/gophkeeper/*.go

// ./cmd/gophkeeper/gophkeeper.exe -c ./config/config.json
func main() {
	ctx := context.Background()

	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
	fmt.Printf("Build commit: %s\n", buildCommit)

	// Подготовка Dependency Injection
	app := fx.New(App(ctx))

	// Запуск приложения
	if err := app.Start(ctx); err != nil {
		log.Println("Start gRPC server error: " + err.Error())
	}

	// Отлавливаем сигналы операционной системы
	<-app.Done()

	// Завершаем работу
	if err := app.Stop(ctx); err != nil {
		log.Println("Stop gRPC server error: " + err.Error())
	}

}

func App(ctx context.Context) fx.Option {
	return fx.Options(
		fx.Provide(
			// Контекст делаем общедоступным
			func() context.Context { return ctx },

			newConfig,
			newLogger,

			// Annotate gRPC server instance as grpc.ServiceRegistrar
			fx.Annotate(
				server.NewGRPCServer,
				fx.As(new(grpc.ServiceRegistrar)),
			),

			// // Annotate service as generated interface
			// fx.Annotate(
			// 	pb.NewGophkeeper,
			// 	fx.As(new(pb.GophkeeperServer)),
			// ),
		),
		fx.Invoke(
			// Start annotated gRPC server
			func(grpc.ServiceRegistrar) {},

			// pb.RegisterGophkeeperServer,
		),
		fx.WithLogger(func(logger *slog.Logger) fxevent.Logger {
			return &fxevent.SlogLogger{Logger: logger}
		}),
	)
}

func run(

	// Объект жизненного цикла fx
	lifecycle fx.Lifecycle,
) {
	lifecycle.Append(fx.Hook{
		// Событие при старте
		OnStart: func(ctx context.Context) error {

			return nil
		},
		// Событие при остановке
		OnStop: func(ctx context.Context) error {
			// Логика на завершение
			return nil
		},
	})
}

// Конфигурируем приложение
func newConfig() (*config.Config, error) {
	return config.NewConfig()
}

func newLogger(cfg *config.Config) (*slog.Logger, error) {
	lg, err := logger.NewLogger(cfg.LogLevel)
	if err != nil {
		return nil, err
	}
	lg.Info("configuration", "config", cfg.Json())
	return logger.NewLogger(cfg.LogLevel)
}
