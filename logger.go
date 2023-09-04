package logger

import (
	"github.com/lmittmann/tint"
	"log/slog"
	"os"
	"time"
)

type Logger struct {
	level  slog.Level
	logger *slog.Logger
}

// New конфигурация логгера
// Для конфигурации вывода slog при DEV_MODE=1 используется пакет ... https://github.com/lmittmann/tint
func New(devMode bool) *Logger {
	// customization slog logger
	var logger *slog.Logger

	if devMode {
		logger = slog.New(tint.NewHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.RFC822,
		}))
	} else {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))
	}

	return &Logger{
		logger: logger,
	}
}

func (l *Logger) Info(message string, args ...any) {
	l.logger.Info(message, args...)
}

func (l *Logger) Error(message string, args ...any) {
	l.logger.Error(message, args...)
}

func (l *Logger) Warn(message string, args ...any) {
	l.logger.Warn(message, args...)
}

func (l *Logger) Debug(message string, args ...any) {
	l.logger.Debug(message, args...)
}
