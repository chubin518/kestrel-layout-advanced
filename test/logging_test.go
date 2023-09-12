package test

import (
	"context"
	"testing"
	"time"

	"github.com/chubin518/kestrel-layout-advanced/pkg/logging"
)

func TestLogging(t *testing.T) {
	logging.Debug("debug message")
	logging.Info("info message")
	logging.WarnContext(nil, "warn message %d", time.Now().UnixNano())

	log := logging.WithAttrs(map[string]any{
		"version": "0.0.1",
	})
	log.Error("error message with version")

	logging.SetDefault(log)

	logging.Info("info message with version")

}

func TestLoggingWithOptions(t *testing.T) {
	log := logging.New(logging.WithLevel(logging.LevelDebug))
	log.Error("error message with options")

	logging.SetDefault(log)

	logging.Info("info message with options")

	logging.Debug("debug message with options")
}

func TestLoggingWithContext(t *testing.T) {
	ctx := logging.WithContext(context.Background(), map[string]any{
		"traceId": time.Now().UnixNano(),
	})

	logging.Info("info message without context")
	logging.WithAttrs(map[string]any{
		"ctx": 1,
	}).InfoContext(ctx, "info message with context")

	ctx = logging.WithContext(ctx, map[string]any{
		"v": "1.0.0",
	})
	logging.Warn("warn message without context")
	logging.WithAttrs(map[string]any{
		"ctx": 2,
	}).WarnContext(ctx, "warn message with context")
	ctx = logging.WithContext(context.Background(), map[string]any{
		"traceId": time.Now().UnixNano(),
	})
	logging.Error("error message without context")
	logging.WithAttrs(map[string]any{
		"ctx": 2,
	}).ErrorContext(ctx, "error message with context")

	log := logging.Default()

	// log = log.WithAttrs(map[string]any{"d": false})

	log.Info("testing start ...")
	ctx = log.WithContext(context.Background(), map[string]any{
		"traceid": time.Now().UnixNano(),
	})
	log.InfoContext(ctx, "info message")

	ctx = log.WithContext(context.Background(), map[string]any{
		"traceid": time.Now().UnixNano(),
	})

	log.ErrorContext(ctx, "error message")

	log.Info("testing end ...")

}
