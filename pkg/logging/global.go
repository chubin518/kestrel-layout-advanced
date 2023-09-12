package logging

import (
	"context"
	"sync/atomic"
)

var defaultLogger atomic.Value

func init() {
	log := New()
	SetDefault(log)
}

// New
func New(options ...Option) ILogging {
	w, l := NewWriter(options...)
	return NewZapLogging(w, l)
}

// Default
func Default() ILogging {
	return defaultLogger.Load().(ILogging)
}

// SetDefault
func SetDefault(logging ILogging) {
	defaultLogger.Store(logging)
}

// Debug calls ILogging.Debug on the default logger.
func Debug(msg string, args ...any) {
	Default().Log(context.Background(), LevelDebug, msg, args...)
}

// DebugContext calls ILogging.DebugContext on the default logger.
func DebugContext(ctx context.Context, msg string, args ...any) {
	Default().Log(ctx, LevelDebug, msg, args...)
}

// Info calls ILogging.Info on the default logger.
func Info(msg string, args ...any) {
	Default().Log(context.Background(), LevelInfo, msg, args...)
}

// InfoContext calls ILogging.InfoContext on the default logger.
func InfoContext(ctx context.Context, msg string, args ...any) {
	Default().Log(ctx, LevelInfo, msg, args...)
}

// Warn calls ILogging.Warn on the default logger.
func Warn(msg string, args ...any) {
	Default().Log(context.Background(), LevelWarn, msg, args...)
}

// WarnContext calls ILogging.WarnContext on the default logger.
func WarnContext(ctx context.Context, msg string, args ...any) {
	Default().Log(ctx, LevelWarn, msg, args...)
}

// Error calls ILogging.Error on the default logger.
func Error(msg string, args ...any) {
	Default().Log(context.Background(), LevelError, msg, args...)
}

// ErrorContext calls ILogging.ErrorContext on the default logger.
func ErrorContext(ctx context.Context, msg string, args ...any) {
	Default().Log(ctx, LevelError, msg, args...)
}

// Fatal calls ILogging.Fatal on the default logger.
func Fatal(msg string, args ...any) {
	Default().Log(context.Background(), LevelFatal, msg, args...)
}

// FatalContext calls ILogging.FatalContext on the default logger.
func FatalContext(ctx context.Context, msg string, args ...any) {
	Default().Log(ctx, LevelFatal, msg, args...)
}

// Log calls ILogging.Log on the default logger.
func Log(ctx context.Context, level Level, msg string, args ...any) {
	Default().Log(ctx, level, msg, args...)
}

// WithAttrs calls Logger.With on the default logger.
func WithAttrs(attrs map[string]any) ILogging {
	return Default().WithAttrs(attrs)
}

// WithContext calls Logger.WithContext on the default logger.
func WithContext(ctx context.Context, attrs map[string]any) context.Context {
	return Default().WithContext(ctx, attrs)
}
