package logger

import (
	"log/slog"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"
	"go.uber.org/zap/zapcore"
)

// Инициализация логирования
func NewLogger(lvl string, opts ...zap.Option) (*slog.Logger, error) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.DisableStacktrace = true
	config.Level = level(lvl)
	z, err := config.Build(opts...)
	if err != nil {
		return nil, err
	}
	return slog.New(zapslog.NewHandler(z.Core(),
			&zapslog.HandlerOptions{AddSource: !config.DisableCaller})),
		nil
}

// Настройка уровня логирования
func level(lvl string) zap.AtomicLevel {
	switch strings.ToLower(lvl) {
	case "error":
		return zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case "warn":
		return zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "info":
		return zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "debug":
		return zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "fatal":
		return zap.NewAtomicLevelAt(zapcore.FatalLevel)
	default:
		return zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}
}
