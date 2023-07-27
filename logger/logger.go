package logger

import (
	"errors"
	"os"

	"github.com/smochii/go-clean-architecture/config"
	"golang.org/x/exp/slog"
)

var Logger *slog.Logger

func Error(err string) {
	Logger.Error(err)
}

func Warn(err string) {
	Logger.Warn(err)
}

func Info(err string) {
	Logger.Info(err)
}

func Debug(err string) {
	Logger.Debug(err)
}

func parse(logLevel string) (slog.Leveler, error) {
	switch logLevel {
	case "debug":
		return slog.LevelDebug, nil
	case "info":
		return slog.LevelInfo, nil
	case "warn":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	default:
		return nil, errors.New("unknown log level")
	}
}

func init() {
	level, err := parse(config.Conf.App.LogLevel)
	if err != nil {
		panic(err)
	}
	Logger = slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	}))
}
