package logging

import (
	"context"
)

type ContextKey struct{}

type ILogging interface {

	// Debug logs at LevelDebug.
	Debug(msg string, args ...any)

	// DebugContext logs at LevelDebug with the given context.
	DebugContext(ctx context.Context, msg string, args ...any)

	// Info logs at LevelInfo
	Info(msg string, args ...any)

	// InfoContext logs at LevelInfo with the given context.
	InfoContext(ctx context.Context, msg string, args ...any)

	// Warn logs at LevelWarn
	Warn(msg string, args ...any)

	// WarnContext logs at LevelWarn with the given context.
	WarnContext(ctx context.Context, msg string, args ...any)

	// Error logs at LevelError
	Error(msg string, args ...any)

	// ErrorContext logs at LevelError with the given context.
	ErrorContext(ctx context.Context, msg string, args ...any)

	// Fatal logs at LevelFatal.
	Fatal(msg string, args ...any)

	// FatalContext logs at LevelFatal with the given context.
	FatalContext(ctx context.Context, msg string, args ...any)

	// Log emits a log record with the given level and message.
	Log(ctx context.Context, level Level, msg string, args ...any)

	// WithAttrs returns a Logger that includes the given attributes in each output operation.
	WithAttrs(attrs map[string]any) ILogging

	// WithContext Add attrs to the specified context
	WithContext(ctx context.Context, attrs map[string]any) context.Context
}
