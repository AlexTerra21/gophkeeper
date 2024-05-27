package main

import (
	"context"
	"fmt"
	"log"

	"go.uber.org/fx"
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

// go build -o cmd/gophkeeper/gophkeeper.exe -ldflags "-X main.buildVersion=v1.20.0 -X 'main.buildDate=$(date +'%Y/%m/%d %H:%M:%S')' -X 'main.buildCommit=$(git log -1 | grep commit)'" cmd/gophkeeper/*.go
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

			//newConfig,
		),
		fx.Invoke(run),
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
func newConfig() {

}
